package godateparser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/coredds/godateparser/translations"
)

// Extended relative date patterns for v1.0.0
// Supports: period boundaries, quarters, complex expressions

// Period boundary patterns
var periodBoundaryPatterns = []*relativePattern{
	// Beginning/start of period
	{
		regex: regexp.MustCompile(`(?i)^(beginning|start|first day) of (month|year|week)$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			period := strings.ToLower(matches[2])
			return getStartOfPeriod(ctx.settings.RelativeBase, period), nil
		},
	},
	// End/last day of period
	{
		regex: regexp.MustCompile(`(?i)^(end|last day) of (month|year|week)$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			period := strings.ToLower(matches[2])
			return getEndOfPeriod(ctx.settings.RelativeBase, period), nil
		},
	},
	// Beginning/start of last/next period
	{
		regex: regexp.MustCompile(`(?i)^(beginning|start|first day) of (last|next) (month|year|week)$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			direction := strings.ToLower(matches[2])
			period := strings.ToLower(matches[3])
			base := ctx.settings.RelativeBase

			// Move to next/last period first
			if direction == "next" {
				base = addPeriod(base, period, 1)
			} else {
				base = addPeriod(base, period, -1)
			}

			return getStartOfPeriod(base, period), nil
		},
	},
	// End of last/next period
	{
		regex: regexp.MustCompile(`(?i)^(end|last day) of (last|next) (month|year|week)$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			direction := strings.ToLower(matches[2])
			period := strings.ToLower(matches[3])
			base := ctx.settings.RelativeBase

			// Move to next/last period first
			if direction == "next" {
				base = addPeriod(base, period, 1)
			} else {
				base = addPeriod(base, period, -1)
			}

			return getEndOfPeriod(base, period), nil
		},
	},
}

// This/next/last disambiguation patterns
var thisNextPatterns = []*relativePattern{
	// "this Monday", "this Friday"
	{
		regex: regexp.MustCompile(`(?i)^this (monday|tuesday|wednesday|thursday|friday|saturday|sunday)$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			weekday := parseWeekday(matches[1])
			base := ctx.settings.RelativeBase

			// "this Monday" means:
			// - If today is Monday or before, return this week's Monday
			// - If today is after Monday, return next week's Monday
			current := base.Weekday()
			daysAhead := int(weekday - current)

			if daysAhead < 0 {
				daysAhead += 7
			}

			return base.AddDate(0, 0, daysAhead), nil
		},
	},
	// "this month", "this year", "this week"
	{
		regex: regexp.MustCompile(`(?i)^this (month|year|week)$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			// "this month/year/week" returns the start of current period
			period := strings.ToLower(matches[1])
			return getStartOfPeriod(ctx.settings.RelativeBase, period), nil
		},
	},
}

// Complex relative expression patterns
var complexRelativePatterns = []*relativePattern{
	// "a week from Tuesday", "2 days from Monday"
	{
		regex: regexp.MustCompile(`(?i)^(a|an|\d+) (day|week|month|year)s? from (monday|tuesday|wednesday|thursday|friday|saturday|sunday)$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			// Parse amount
			amount := 1
			if matches[1] != "a" && matches[1] != "an" {
				amount, _ = strconv.Atoi(matches[1])
			}

			unit := strings.ToLower(matches[2])
			weekday := parseWeekday(matches[3])

			// Find next occurrence of weekday
			base := ctx.settings.RelativeBase
			targetDate := findWeekday(base, weekday, true)

			// Add the offset
			return addDuration(targetDate, amount, unit), nil
		},
	},
	// "3 days after tomorrow", "2 weeks before yesterday"
	{
		regex: regexp.MustCompile(`(?i)^(\d+) (day|week|month|year)s? (after|before) (yesterday|today|tomorrow)$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			amount, _ := strconv.Atoi(matches[1])
			unit := strings.ToLower(matches[2])
			direction := strings.ToLower(matches[3])
			baseDay := strings.ToLower(matches[4])

			if direction == "before" {
				amount = -amount
			}

			// Get base day
			base := ctx.settings.RelativeBase
			switch baseDay {
			case "yesterday":
				base = base.AddDate(0, 0, -1)
			case "tomorrow":
				base = base.AddDate(0, 0, 1)
			}

			return addDuration(base, amount, unit), nil
		},
	},
	// "2 weeks before last Monday"
	{
		regex: regexp.MustCompile(`(?i)^(\d+) (day|week|month|year)s? (after|before) (next|last) (monday|tuesday|wednesday|thursday|friday|saturday|sunday)$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			amount, _ := strconv.Atoi(matches[1])
			unit := strings.ToLower(matches[2])
			direction := strings.ToLower(matches[3])
			nextLast := strings.ToLower(matches[4])
			weekday := parseWeekday(matches[5])

			if direction == "before" {
				amount = -amount
			}

			// Find the weekday
			base := ctx.settings.RelativeBase
			targetDate := findWeekday(base, weekday, nextLast == "next")

			return addDuration(targetDate, amount, unit), nil
		},
	},
}

// Quarter patterns
var quarterPatterns = []*relativePattern{
	// "Q1", "Q2", "Q3", "Q4"
	{
		regex: regexp.MustCompile(`(?i)^Q([1-4])$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			quarter, _ := strconv.Atoi(matches[1])
			year := ctx.settings.RelativeBase.Year()
			return getQuarterStart(year, quarter), nil
		},
	},
	// "Q1 2024", "Q4 2025"
	{
		regex: regexp.MustCompile(`(?i)^Q([1-4])\s+(\d{4})$`),
		parser: func(_ *parserContext, matches []string) (time.Time, error) {
			quarter, _ := strconv.Atoi(matches[1])
			year, _ := strconv.Atoi(matches[2])
			return getQuarterStart(year, quarter), nil
		},
	},
	// "last quarter", "next quarter"
	{
		regex: regexp.MustCompile(`(?i)^(last|next|this) quarter$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			direction := strings.ToLower(matches[1])
			base := ctx.settings.RelativeBase

			currentQuarter := getQuarter(base)
			year := base.Year()

			switch direction {
			case "this":
				return getQuarterStart(year, currentQuarter), nil
			case "next":
				if currentQuarter == 4 {
					return getQuarterStart(year+1, 1), nil
				}
				return getQuarterStart(year, currentQuarter+1), nil
			case "last":
				if currentQuarter == 1 {
					return getQuarterStart(year-1, 4), nil
				}
				return getQuarterStart(year, currentQuarter-1), nil
			}

			return time.Time{}, fmt.Errorf("invalid quarter direction")
		},
	},
}

// Helper functions

// getStartOfPeriod returns the start of the given period
func getStartOfPeriod(t time.Time, period string) time.Time {
	switch period {
	case "week":
		// Monday of current week
		days := int(t.Weekday())
		if days == 0 { // Sunday
			days = 7
		}
		return time.Date(t.Year(), t.Month(), t.Day()-(days-1), 0, 0, 0, 0, t.Location())
	case "month":
		return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	case "year":
		return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
	}
	return t
}

// getEndOfPeriod returns the end of the given period
func getEndOfPeriod(t time.Time, period string) time.Time {
	switch period {
	case "week":
		// Sunday of current week
		days := int(t.Weekday())
		if days == 0 { // Sunday
			return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
		}
		return time.Date(t.Year(), t.Month(), t.Day()+(7-days), 23, 59, 59, 999999999, t.Location())
	case "month":
		// Last day of month
		firstOfNextMonth := time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location())
		lastOfMonth := firstOfNextMonth.AddDate(0, 0, -1)
		return time.Date(lastOfMonth.Year(), lastOfMonth.Month(), lastOfMonth.Day(), 23, 59, 59, 999999999, t.Location())
	case "year":
		return time.Date(t.Year(), 12, 31, 23, 59, 59, 999999999, t.Location())
	}
	return t
}

// addPeriod adds/subtracts a period from a date
func addPeriod(t time.Time, period string, amount int) time.Time {
	switch period {
	case "week":
		return t.AddDate(0, 0, amount*7)
	case "month":
		return t.AddDate(0, amount, 0)
	case "year":
		return t.AddDate(amount, 0, 0)
	}
	return t
}

// getQuarter returns the quarter (1-4) for a given date
func getQuarter(t time.Time) int {
	month := int(t.Month())
	return ((month - 1) / 3) + 1
}

// getQuarterStart returns the start of a quarter
func getQuarterStart(year, quarter int) time.Time {
	month := ((quarter - 1) * 3) + 1
	return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
}

// tryParseExtendedRelative attempts to parse extended relative patterns
func tryParseExtendedRelative(ctx *parserContext) (time.Time, error) {
	input := strings.ToLower(strings.TrimSpace(ctx.input))

	// Try multi-language extended patterns first
	if result, err := tryParseMultiLangExtended(ctx, input); err == nil {
		return result, nil
	}

	// Try English-only patterns as fallback
	// Try period boundaries
	for _, pattern := range periodBoundaryPatterns {
		matches := pattern.regex.FindStringSubmatch(input)
		if matches != nil {
			return pattern.parser(ctx, matches)
		}
	}

	// Try this/next patterns
	for _, pattern := range thisNextPatterns {
		matches := pattern.regex.FindStringSubmatch(input)
		if matches != nil {
			return pattern.parser(ctx, matches)
		}
	}

	// Try complex relative patterns
	for _, pattern := range complexRelativePatterns {
		matches := pattern.regex.FindStringSubmatch(input)
		if matches != nil {
			return pattern.parser(ctx, matches)
		}
	}

	// Try quarter patterns
	for _, pattern := range quarterPatterns {
		matches := pattern.regex.FindStringSubmatch(input)
		if matches != nil {
			return pattern.parser(ctx, matches)
		}
	}

	return time.Time{}, fmt.Errorf("no extended relative pattern matched")
}

// tryParseMultiLangExtended attempts to parse extended patterns in multiple languages
func tryParseMultiLangExtended(ctx *parserContext, input string) (time.Time, error) {
	for _, lang := range ctx.languages {
		if lang.RelativeTerms == nil {
			continue
		}

		// Try period boundaries: "comienzo de mes", "fin de año"
		if result, err := tryParsePeriodBoundary(ctx, input, lang); err == nil {
			return result, nil
		}

		// Try this/next/last with weekdays: "este lunes", "próxima semana"
		if result, err := tryParseThisNextLast(ctx, input, lang); err == nil {
			return result, nil
		}
	}

	return time.Time{}, fmt.Errorf("no multi-lang extended pattern matched")
}

// tryParsePeriodBoundary parses "comienzo de mes", "fin de año", etc.
func tryParsePeriodBoundary(ctx *parserContext, input string, lang *translations.Language) (time.Time, error) {
	terms := lang.RelativeTerms
	base := ctx.settings.RelativeBase

	// Build patterns for beginning/start/end
	beginTerms := terms.Beginning
	beginTerms = append(beginTerms, terms.Start...)
	beginTerms = append(beginTerms, terms.First...)
	endTerms := terms.End

	// Try "beginning/start/end of month/year/week"
	periods := map[string]string{
		// Spanish
		"mes":     "month",
		"meses":   "month",
		"año":     "year",
		"ano":     "year",
		"anos":    "year",
		"años":    "year",
		"semana":  "week",
		"semanas": "week",
		// Portuguese
		"mês": "month",
		// French
		"mois":    "month",
		"semaine": "week",
		"année":   "year",
		"annee":   "year",
		// German
		"monat":   "month",
		"monate":  "month",
		"monaten": "month",
		"jahr":    "year",
		"jahre":   "year",
		"jahren":  "year",
		"woche":   "week",
		"wochen":  "week",
		// Italian
		"mese":      "month",
		"mesi":      "month",
		"anno":      "year",
		"anni":      "year",
		"settimana": "week",
		"settimane": "week",
		// Dutch
		"maand":   "month",
		"maanden": "month",
		"jaar":    "year", // Dutch jaar (same spelling as German)
		"weken":   "week",
		// Russian
		"месяц":   "month",
		"месяца":  "month",
		"месяцев": "month",
		"год":     "year",
		"года":    "year",
		"лет":     "year",
		"неделя":  "week",
		"недели":  "week",
		"недель":  "week",
		"неделю":  "week", // Accusative case
		// English
		"month": "month",
		"year":  "year",
		"week":  "week",
	}

	for periodEs, periodEn := range periods {
		// Try "comienzo de mes" (Spanish/Portuguese/French: de, Italian: di, Dutch: van, Russian: no preposition)
		for _, beginTerm := range beginTerms {
			pattern := fmt.Sprintf(`^%s\s+(de\s+|di\s+|van\s+)?%s$`, regexp.QuoteMeta(beginTerm), regexp.QuoteMeta(periodEs))
			if matched, _ := regexp.MatchString(pattern, input); matched {
				return getStartOfPeriod(base, periodEn), nil
			}
		}

		// Try "fin de mes"
		for _, endTerm := range endTerms {
			pattern := fmt.Sprintf(`^%s\s+(de\s+|di\s+|van\s+)?%s$`, regexp.QuoteMeta(endTerm), regexp.QuoteMeta(periodEs))
			if matched, _ := regexp.MatchString(pattern, input); matched {
				return getEndOfPeriod(base, periodEn), nil
			}
		}

		// Try "próximo mes", "último año"
		for _, nextTerm := range terms.Next {
			pattern := fmt.Sprintf(`^%s\s+%s$`, regexp.QuoteMeta(nextTerm), regexp.QuoteMeta(periodEs))
			if matched, _ := regexp.MatchString(pattern, input); matched {
				return addPeriod(base, periodEn, 1), nil
			}

			// "comienzo de próximo mes" (with various prepositions)
			for _, beginTerm := range beginTerms {
				pattern := fmt.Sprintf(`^%s\s+(de\s+|di\s+|van\s+)?%s\s+%s$`, regexp.QuoteMeta(beginTerm), regexp.QuoteMeta(nextTerm), regexp.QuoteMeta(periodEs))
				if matched, _ := regexp.MatchString(pattern, input); matched {
					nextPeriod := addPeriod(base, periodEn, 1)
					return getStartOfPeriod(nextPeriod, periodEn), nil
				}
			}

			// "fin de próximo mes"
			for _, endTerm := range endTerms {
				pattern := fmt.Sprintf(`^%s\s+(de\s+|di\s+|van\s+)?%s\s+%s$`, regexp.QuoteMeta(endTerm), regexp.QuoteMeta(nextTerm), regexp.QuoteMeta(periodEs))
				if matched, _ := regexp.MatchString(pattern, input); matched {
					nextPeriod := addPeriod(base, periodEn, 1)
					return getEndOfPeriod(nextPeriod, periodEn), nil
				}
			}
		}

		for _, lastTerm := range terms.Last {
			pattern := fmt.Sprintf(`^%s\s+%s$`, regexp.QuoteMeta(lastTerm), regexp.QuoteMeta(periodEs))
			if matched, _ := regexp.MatchString(pattern, input); matched {
				return addPeriod(base, periodEn, -1), nil
			}

			// "comienzo de último mes" (with various prepositions)
			for _, beginTerm := range beginTerms {
				pattern := fmt.Sprintf(`^%s\s+(de\s+|di\s+|van\s+)?%s\s+%s$`, regexp.QuoteMeta(beginTerm), regexp.QuoteMeta(lastTerm), regexp.QuoteMeta(periodEs))
				if matched, _ := regexp.MatchString(pattern, input); matched {
					lastPeriod := addPeriod(base, periodEn, -1)
					return getStartOfPeriod(lastPeriod, periodEn), nil
				}
			}

			// "fin de último mes"
			for _, endTerm := range endTerms {
				pattern := fmt.Sprintf(`^%s\s+(de\s+|di\s+|van\s+)?%s\s+%s$`, regexp.QuoteMeta(endTerm), regexp.QuoteMeta(lastTerm), regexp.QuoteMeta(periodEs))
				if matched, _ := regexp.MatchString(pattern, input); matched {
					lastPeriod := addPeriod(base, periodEn, -1)
					return getEndOfPeriod(lastPeriod, periodEn), nil
				}
			}
		}
	}

	return time.Time{}, fmt.Errorf("no period boundary matched")
}

// tryParseThisNextLast parses "este lunes", "próxima semana", etc.
func tryParseThisNextLast(ctx *parserContext, input string, lang *translations.Language) (time.Time, error) {
	terms := lang.RelativeTerms
	base := ctx.settings.RelativeBase

	// Try "this Monday" / "este lunes"
	for _, thisTerm := range terms.This {
		for weekdayName, weekday := range lang.Weekdays {
			pattern := fmt.Sprintf(`^%s\s+%s$`, regexp.QuoteMeta(thisTerm), regexp.QuoteMeta(weekdayName))
			if matched, _ := regexp.MatchString(pattern, input); matched {
				current := base.Weekday()
				daysAhead := int(weekday - current)
				if daysAhead < 0 {
					daysAhead += 7
				}
				return base.AddDate(0, 0, daysAhead), nil
			}
		}

		// "this month/year/week" / "este mes/año/semana"
		periods := map[string]string{
			// Spanish
			"mes":     "month",
			"meses":   "month",
			"año":     "year",
			"ano":     "year",
			"anos":    "year",
			"años":    "year",
			"semana":  "week",
			"semanas": "week",
			// Portuguese
			"mês": "month",
			// French
			"semaine": "week",
			"mois":    "month",
			"année":   "year",
			"annee":   "year",
			// German
			"monat":   "month",
			"monate":  "month",
			"monaten": "month",
			"jahr":    "year",
			"jahre":   "year",
			"jahren":  "year",
			"woche":   "week",
			"wochen":  "week",
			// Italian
			"mese":      "month",
			"mesi":      "month",
			"anno":      "year",
			"anni":      "year",
			"settimana": "week",
			"settimane": "week",
			// Dutch
			"maand":   "month",
			"maanden": "month",
			"weken":   "week",
			// Russian
			"месяц":   "month",
			"месяца":  "month",
			"месяцев": "month",
			"год":     "year",
			"года":    "year",
			"лет":     "year",
			"неделя":  "week",
			"недели":  "week",
			"недель":  "week",
			// English
			"month": "month",
			"year":  "year",
			"week":  "week",
		}
		for periodEs, periodEn := range periods {
			pattern := fmt.Sprintf(`^%s\s+%s$`, regexp.QuoteMeta(thisTerm), regexp.QuoteMeta(periodEs))
			if matched, _ := regexp.MatchString(pattern, input); matched {
				return getStartOfPeriod(base, periodEn), nil
			}
		}
	}

	return time.Time{}, fmt.Errorf("no this/next/last pattern matched")
}
