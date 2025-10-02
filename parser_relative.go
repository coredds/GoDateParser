package godateparser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/coredds/GoDateParser/translations"
)

// Relative date patterns
var relativePatterns = []*relativePattern{
	// "2 days ago", "3 weeks ago", "a fortnight ago"
	{
		regex: regexp.MustCompile(`(?i)^(a|an|\d+)\s+(second|minute|hour|day|week|fortnight|month|quarter|year|decade)s?\s+ago$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			amount := 1
			if matches[1] != "a" && matches[1] != "an" {
				amount, _ = strconv.Atoi(matches[1])
			}
			unit := strings.ToLower(matches[2])
			return addDuration(ctx.settings.RelativeBase, -amount, unit), nil
		},
	},
	// "in 2 days", "in 3 weeks", "in a fortnight"
	{
		regex: regexp.MustCompile(`(?i)^in\s+(a|an|\d+)\s+(second|minute|hour|day|week|fortnight|month|quarter|year|decade)s?$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			amount := 1
			if matches[1] != "a" && matches[1] != "an" {
				amount, _ = strconv.Atoi(matches[1])
			}
			unit := strings.ToLower(matches[2])
			return addDuration(ctx.settings.RelativeBase, amount, unit), nil
		},
	},
	// "yesterday", "today", "tomorrow"
	{
		regex: regexp.MustCompile(`(?i)^(yesterday|today|tomorrow)$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			base := ctx.settings.RelativeBase
			switch strings.ToLower(matches[1]) {
			case "yesterday":
				return base.AddDate(0, 0, -1), nil
			case "today":
				return base, nil
			case "tomorrow":
				return base.AddDate(0, 0, 1), nil
			}
			return time.Time{}, fmt.Errorf("unknown relative date")
		},
	},
	// "last week", "last month", "last year", "last fortnight", "last decade"
	{
		regex: regexp.MustCompile(`(?i)^last\s+(week|fortnight|month|quarter|year|decade)$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			unit := strings.ToLower(matches[1])
			return addDuration(ctx.settings.RelativeBase, -1, unit), nil
		},
	},
	// "next week", "next month", "next year", "next fortnight", "next decade"
	{
		regex: regexp.MustCompile(`(?i)^next\s+(week|fortnight|month|quarter|year|decade)$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			unit := strings.ToLower(matches[1])
			return addDuration(ctx.settings.RelativeBase, 1, unit), nil
		},
	},
	// "next Monday", "last Friday"
	{
		regex: regexp.MustCompile(`(?i)^(next|last)\s+(monday|tuesday|wednesday|thursday|friday|saturday|sunday)$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			direction := strings.ToLower(matches[1])
			weekday := parseWeekday(matches[2])
			return findWeekday(ctx.settings.RelativeBase, weekday, direction == "next"), nil
		},
	},
	// Standalone weekday (e.g., "Monday" without next/last)
	{
		regex: regexp.MustCompile(`(?i)^(monday|tuesday|wednesday|thursday|friday|saturday|sunday)$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			weekday := parseWeekday(matches[1])
			// Use PreferDatesFrom setting to disambiguate
			preferFuture := ctx.settings.PreferDatesFrom != "past"
			return findWeekday(ctx.settings.RelativeBase, weekday, preferFuture), nil
		},
	},
	// "now"
	{
		regex: regexp.MustCompile(`(?i)^now$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			return time.Now(), nil
		},
	},
}

type relativePattern struct {
	regex  *regexp.Regexp
	parser func(*parserContext, []string) (time.Time, error)
}

// parseRelative attempts to parse relative date expressions.
func parseRelative(ctx *parserContext) (time.Time, error) {
	input := strings.TrimSpace(ctx.input)

	// Try multi-language relative patterns first
	if result, err := tryParseMultiLangRelative(ctx, input); err == nil {
		return result, nil
	}

	// Try extended relative patterns (v1.0 features) - these are more specific
	// For example, "next quarter" should use quarter-aware logic, not simple +3 months
	result, err := tryParseExtendedRelative(ctx)
	if err == nil {
		return result, nil
	}

	// Try basic relative patterns as fallback
	for _, pattern := range relativePatterns {
		matches := pattern.regex.FindStringSubmatch(input)
		if matches != nil {
			return pattern.parser(ctx, matches)
		}
	}

	return time.Time{}, fmt.Errorf("no relative date pattern matched")
}

// addDuration adds a duration to a base time based on unit and amount.
func addDuration(base time.Time, amount int, unit string) time.Time {
	switch unit {
	case "second":
		return base.Add(time.Duration(amount) * time.Second)
	case "minute":
		return base.Add(time.Duration(amount) * time.Minute)
	case "hour":
		return base.Add(time.Duration(amount) * time.Hour)
	case "day":
		return base.AddDate(0, 0, amount)
	case "week":
		return base.AddDate(0, 0, amount*7)
	case "fortnight":
		return base.AddDate(0, 0, amount*14)
	case "month":
		return base.AddDate(0, amount, 0)
	case "quarter":
		return base.AddDate(0, amount*3, 0)
	case "year":
		return base.AddDate(amount, 0, 0)
	case "decade":
		return base.AddDate(amount*10, 0, 0)
	default:
		return base
	}
}

// parseWeekday converts weekday name to time.Weekday.
func parseWeekday(weekday string) time.Weekday {
	weekday = strings.ToLower(weekday)
	weekdayMap := map[string]time.Weekday{
		"sunday":    time.Sunday,
		"monday":    time.Monday,
		"tuesday":   time.Tuesday,
		"wednesday": time.Wednesday,
		"thursday":  time.Thursday,
		"friday":    time.Friday,
		"saturday":  time.Saturday,
	}
	return weekdayMap[weekday]
}

// findWeekday finds the next or previous occurrence of a weekday.
func findWeekday(base time.Time, target time.Weekday, forward bool) time.Time {
	current := base.Weekday()
	var daysToAdd int

	if forward {
		// "next Monday" means the next occurrence of Monday (could be this week if we haven't passed it)
		daysToAdd = int(target - current)
		if daysToAdd <= 0 {
			daysToAdd += 7
		}
	} else {
		// "last Monday" means the previous occurrence of Monday
		daysToAdd = int(target - current)
		if daysToAdd >= 0 {
			daysToAdd -= 7
		}
	}

	return base.AddDate(0, 0, daysToAdd)
}

// tryParseMultiLangRelative attempts to parse relative dates in multiple languages.
func tryParseMultiLangRelative(ctx *parserContext, input string) (time.Time, error) {
	input = strings.ToLower(strings.TrimSpace(input))
	base := ctx.settings.RelativeBase

	// Try each language's relative terms
	for _, lang := range ctx.languages {
		if lang.RelativeTerms == nil {
			continue
		}

		// Simple terms: yesterday, today, tomorrow
		if input == strings.ToLower(lang.RelativeTerms.Yesterday) {
			return base.AddDate(0, 0, -1), nil
		}
		if input == strings.ToLower(lang.RelativeTerms.Today) {
			return base, nil
		}
		if input == strings.ToLower(lang.RelativeTerms.Tomorrow) {
			return base.AddDate(0, 0, 1), nil
		}

		// Try "hace X días" (X days ago) pattern - PREFIX
		for _, agoTerm := range lang.RelativeTerms.Ago {
			if result, err := tryParseAgoPattern(ctx, input, lang, agoTerm); err == nil {
				return result, nil
			}
		}

		// Try "X días atrás" (X days ago) pattern - SUFFIX
		for _, agoTerm := range lang.RelativeTerms.Ago {
			if result, err := tryParseAgoSuffixPattern(ctx, input, lang, agoTerm); err == nil {
				return result, nil
			}
		}

		// Try "en X días" (in X days) pattern
		for _, inTerm := range lang.RelativeTerms.In {
			if result, err := tryParseInPattern(ctx, input, lang, inTerm); err == nil {
				return result, nil
			}
		}

		// Try "próxima semana" (next week) pattern
		for _, nextTerm := range lang.RelativeTerms.Next {
			if result, err := tryParseNextPattern(ctx, input, lang, nextTerm); err == nil {
				return result, nil
			}
		}

		// Try "última semana" (last week) pattern
		for _, lastTerm := range lang.RelativeTerms.Last {
			if result, err := tryParseLastPattern(ctx, input, lang, lastTerm); err == nil {
				return result, nil
			}
		}

		// Try standalone weekdays
		if result, err := tryParseWeekday(ctx, input, lang); err == nil {
			return result, nil
		}
	}

	return time.Time{}, fmt.Errorf("no multi-language relative pattern matched")
}

// tryParseAgoPattern parses patterns like "hace 2 días" (2 days ago)
func tryParseAgoPattern(ctx *parserContext, input string, lang *translations.Language, agoTerm string) (time.Time, error) {
	// Build time unit pattern
	units := buildTimeUnitPattern(lang)
	if units == "" {
		return time.Time{}, fmt.Errorf("no time units")
	}

	pattern := fmt.Sprintf(`^%s\s+(\d+)\s+(%s)$`, regexp.QuoteMeta(strings.ToLower(agoTerm)), units)
	re := regexp.MustCompile(pattern)

	if matches := re.FindStringSubmatch(input); matches != nil {
		amount, _ := strconv.Atoi(matches[1])
		unit := normalizeTimeUnit(matches[2], lang)
		return addDuration(ctx.settings.RelativeBase, -amount, unit), nil
	}

	return time.Time{}, fmt.Errorf("no match")
}

// tryParseAgoSuffixPattern parses suffix patterns like "2 días atrás" (2 days ago)
func tryParseAgoSuffixPattern(ctx *parserContext, input string, lang *translations.Language, agoTerm string) (time.Time, error) {
	// Build time unit pattern
	units := buildTimeUnitPattern(lang)
	if units == "" {
		return time.Time{}, fmt.Errorf("no time units")
	}

	// Pattern: "2 días atrás" - number FIRST, unit SECOND, ago term LAST
	pattern := fmt.Sprintf(`^(\d+)\s+(%s)\s+%s$`, units, regexp.QuoteMeta(strings.ToLower(agoTerm)))
	re := regexp.MustCompile(pattern)

	if matches := re.FindStringSubmatch(input); matches != nil {
		amount, _ := strconv.Atoi(matches[1])
		unit := normalizeTimeUnit(matches[2], lang)
		return addDuration(ctx.settings.RelativeBase, -amount, unit), nil
	}

	return time.Time{}, fmt.Errorf("no match")
}

// tryParseInPattern parses patterns like "en 3 semanas" (in 3 weeks)
func tryParseInPattern(ctx *parserContext, input string, lang *translations.Language, inTerm string) (time.Time, error) {
	// Build time unit pattern
	units := buildTimeUnitPattern(lang)
	if units == "" {
		return time.Time{}, fmt.Errorf("no time units")
	}

	pattern := fmt.Sprintf(`^%s\s+(\d+)\s+(%s)$`, regexp.QuoteMeta(strings.ToLower(inTerm)), units)
	re := regexp.MustCompile(pattern)

	if matches := re.FindStringSubmatch(input); matches != nil {
		amount, _ := strconv.Atoi(matches[1])
		unit := normalizeTimeUnit(matches[2], lang)
		return addDuration(ctx.settings.RelativeBase, amount, unit), nil
	}

	return time.Time{}, fmt.Errorf("no match")
}

// tryParseNextPattern parses patterns like "próxima semana" (next week)
func tryParseNextPattern(ctx *parserContext, input string, lang *translations.Language, nextTerm string) (time.Time, error) {
	// Try "next [unit]" patterns
	units := buildTimeUnitPattern(lang)
	if units == "" {
		return time.Time{}, fmt.Errorf("no time units")
	}

	pattern := fmt.Sprintf(`^%s\s+(%s)$`, regexp.QuoteMeta(strings.ToLower(nextTerm)), units)
	re := regexp.MustCompile(pattern)

	if matches := re.FindStringSubmatch(input); matches != nil {
		unit := normalizeTimeUnit(matches[1], lang)
		return addDuration(ctx.settings.RelativeBase, 1, unit), nil
	}

	// Try "next [weekday]" patterns
	weekdayPattern := buildWeekdayPattern(lang)
	if weekdayPattern != "" {
		pattern := fmt.Sprintf(`^%s\s+(%s)$`, regexp.QuoteMeta(strings.ToLower(nextTerm)), weekdayPattern)
		re := regexp.MustCompile(pattern)

		if matches := re.FindStringSubmatch(input); matches != nil {
			if weekday, ok := lang.Weekdays[matches[1]]; ok {
				return findWeekday(ctx.settings.RelativeBase, weekday, true), nil
			}
		}
	}

	return time.Time{}, fmt.Errorf("no match")
}

// tryParseLastPattern parses patterns like "última semana" (last week)
func tryParseLastPattern(ctx *parserContext, input string, lang *translations.Language, lastTerm string) (time.Time, error) {
	// Try "last [unit]" patterns
	units := buildTimeUnitPattern(lang)
	if units == "" {
		return time.Time{}, fmt.Errorf("no time units")
	}

	pattern := fmt.Sprintf(`^%s\s+(%s)$`, regexp.QuoteMeta(strings.ToLower(lastTerm)), units)
	re := regexp.MustCompile(pattern)

	if matches := re.FindStringSubmatch(input); matches != nil {
		unit := normalizeTimeUnit(matches[1], lang)
		return addDuration(ctx.settings.RelativeBase, -1, unit), nil
	}

	// Try "last [weekday]" patterns
	weekdayPattern := buildWeekdayPattern(lang)
	if weekdayPattern != "" {
		pattern := fmt.Sprintf(`^%s\s+(%s)$`, regexp.QuoteMeta(strings.ToLower(lastTerm)), weekdayPattern)
		re := regexp.MustCompile(pattern)

		if matches := re.FindStringSubmatch(input); matches != nil {
			if weekday, ok := lang.Weekdays[matches[1]]; ok {
				return findWeekday(ctx.settings.RelativeBase, weekday, false), nil
			}
		}
	}

	return time.Time{}, fmt.Errorf("no match")
}

// tryParseWeekday parses standalone weekday names
func tryParseWeekday(ctx *parserContext, input string, lang *translations.Language) (time.Time, error) {
	if weekday, ok := lang.Weekdays[input]; ok {
		preferFuture := ctx.settings.PreferDatesFrom != "past"
		return findWeekday(ctx.settings.RelativeBase, weekday, preferFuture), nil
	}
	return time.Time{}, fmt.Errorf("not a weekday")
}

// buildTimeUnitPattern creates a regex pattern for time units in a language
func buildTimeUnitPattern(lang *translations.Language) string {
	if lang.RelativeTerms == nil {
		return ""
	}

	units := make(map[string]bool)
	addUnits := func(terms []string) {
		for _, term := range terms {
			if term != "" {
				units[regexp.QuoteMeta(strings.ToLower(term))] = true
			}
		}
	}

	addUnits(lang.RelativeTerms.Second)
	addUnits(lang.RelativeTerms.Minute)
	addUnits(lang.RelativeTerms.Hour)
	addUnits(lang.RelativeTerms.Day)
	addUnits(lang.RelativeTerms.Week)
	addUnits(lang.RelativeTerms.Fortnight)
	addUnits(lang.RelativeTerms.Month)
	addUnits(lang.RelativeTerms.Quarter)
	addUnits(lang.RelativeTerms.Year)
	addUnits(lang.RelativeTerms.Decade)

	if len(units) == 0 {
		return ""
	}

	result := ""
	for unit := range units {
		if result != "" {
			result += "|"
		}
		result += unit
	}

	return result
}

// buildWeekdayPattern creates a regex pattern for weekdays in a language
func buildWeekdayPattern(lang *translations.Language) string {
	weekdays := make(map[string]bool)
	for weekday := range lang.Weekdays {
		weekdays[regexp.QuoteMeta(strings.ToLower(weekday))] = true
	}

	if len(weekdays) == 0 {
		return ""
	}

	result := ""
	for weekday := range weekdays {
		if result != "" {
			result += "|"
		}
		result += weekday
	}

	return result
}

// normalizeTimeUnit normalizes a time unit from any language to English
func normalizeTimeUnit(unit string, lang *translations.Language) string {
	unit = strings.ToLower(unit)

	if lang.RelativeTerms == nil {
		return unit
	}

	checkUnit := func(terms []string, normalized string) bool {
		for _, term := range terms {
			if strings.ToLower(term) == unit {
				return true
			}
		}
		return false
	}

	if checkUnit(lang.RelativeTerms.Second, "second") {
		return "second"
	}
	if checkUnit(lang.RelativeTerms.Minute, "minute") {
		return "minute"
	}
	if checkUnit(lang.RelativeTerms.Hour, "hour") {
		return "hour"
	}
	if checkUnit(lang.RelativeTerms.Day, "day") {
		return "day"
	}
	if checkUnit(lang.RelativeTerms.Week, "week") {
		return "week"
	}
	if checkUnit(lang.RelativeTerms.Fortnight, "fortnight") {
		return "fortnight"
	}
	if checkUnit(lang.RelativeTerms.Month, "month") {
		return "month"
	}
	if checkUnit(lang.RelativeTerms.Quarter, "quarter") {
		return "quarter"
	}
	if checkUnit(lang.RelativeTerms.Year, "year") {
		return "year"
	}
	if checkUnit(lang.RelativeTerms.Decade, "decade") {
		return "decade"
	}

	return unit
}
