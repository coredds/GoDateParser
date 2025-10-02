package godateparser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/coredds/GoDateParser/translations"
)

// Incomplete date patterns - dates missing year, day, or both
// Examples: "May", "2024", "June 15"

type incompleteDatePattern struct {
	regex  *regexp.Regexp
	parser func(*parserContext, []string) (time.Time, error)
}

var incompleteDatePatterns = []*incompleteDatePattern{
	// Just year: "2024"
	{
		regex: regexp.MustCompile(`^(19\d{2}|20\d{2})$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			year, _ := strconv.Atoi(matches[1])
			loc := ctx.settings.PreferredTimezone
			// January 1 of that year
			return time.Date(year, 1, 1, 0, 0, 0, 0, loc), nil
		},
	},
	// Just month name: "May", "December", "mayo", "diciembre"
	// Note: Dynamic pattern built at parse time for multi-language support
	{
		regex: nil, // Handled dynamically in tryParseIncompleteDate
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			monthName := strings.ToLower(matches[1])
			month := monthNameToNumberWithLangs(monthName, ctx.languages)

			base := ctx.settings.RelativeBase
			currentYear := base.Year()
			currentMonth := base.Month()

			// Determine year based on PreferDatesFrom setting
			year := currentYear
			if ctx.settings.PreferDatesFrom == "past" {
				// If month is after current month, use last year
				if month > currentMonth {
					year--
				}
			} else {
				// Default "future" behavior
				// If month is before current month, use next year
				if month < currentMonth {
					year++
				}
			}

			loc := ctx.settings.PreferredTimezone
			return time.Date(year, month, 1, 0, 0, 0, 0, loc), nil
		},
	},
	// Month and day without year: "June 15", "junio 15"
	{
		regex: nil, // Handled dynamically in tryParseIncompleteDate
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			monthName := strings.ToLower(matches[1])
			month := monthNameToNumberWithLangs(monthName, ctx.languages)
			day, _ := strconv.Atoi(matches[2])

			// Validate day
			if err := validateDateComponents(0, int(month), day); err != nil {
				return time.Time{}, err
			}

			base := ctx.settings.RelativeBase
			currentYear := base.Year()
			currentMonth := base.Month()
			currentDay := base.Day()

			// Determine year based on PreferDatesFrom setting
			year := currentYear
			if ctx.settings.PreferDatesFrom == "past" {
				// If month/day is after current, use last year
				if month > currentMonth || (month == currentMonth && day > currentDay) {
					year--
				}
			} else {
				// Default "future" behavior
				// If month/day is before current, use next year
				if month < currentMonth || (month == currentMonth && day < currentDay) {
					year++
				}
			}

			loc := ctx.settings.PreferredTimezone
			return time.Date(year, month, day, 0, 0, 0, 0, loc), nil
		},
	},
	// Day and month without year: "15 June", "15 junio"
	{
		regex: nil, // Handled dynamically in tryParseIncompleteDate
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			day, _ := strconv.Atoi(matches[1])
			monthName := strings.ToLower(matches[2])
			month := monthNameToNumberWithLangs(monthName, ctx.languages)

			// Validate day
			if err := validateDateComponents(0, int(month), day); err != nil {
				return time.Time{}, err
			}

			base := ctx.settings.RelativeBase
			currentYear := base.Year()
			currentMonth := base.Month()
			currentDay := base.Day()

			// Determine year based on PreferDatesFrom setting
			year := currentYear
			if ctx.settings.PreferDatesFrom == "past" {
				// If month/day is after current, use last year
				if month > currentMonth || (month == currentMonth && day > currentDay) {
					year--
				}
			} else {
				// Default "future" behavior
				// If month/day is before current, use next year
				if month < currentMonth || (month == currentMonth && day < currentDay) {
					year++
				}
			}

			loc := ctx.settings.PreferredTimezone
			return time.Date(year, month, day, 0, 0, 0, 0, loc), nil
		},
	},
}

// tryParseIncompleteDate attempts to parse incomplete date patterns
func tryParseIncompleteDate(ctx *parserContext) (time.Time, error) {
	input := strings.TrimSpace(ctx.input)

	// Build month pattern from enabled languages
	monthPattern := buildMonthPatternForIncomplete(ctx.languages)

	// Try dynamic month-only pattern
	if monthPattern != "" {
		re := regexp.MustCompile(fmt.Sprintf(`(?i)^(%s)$`, monthPattern))
		if matches := re.FindStringSubmatch(input); matches != nil {
			return incompleteDatePatterns[1].parser(ctx, matches)
		}

		// Try "month day" pattern
		re = regexp.MustCompile(fmt.Sprintf(`(?i)^(%s)\s+(\d{1,2})$`, monthPattern))
		if matches := re.FindStringSubmatch(input); matches != nil {
			return incompleteDatePatterns[2].parser(ctx, matches)
		}

		// Try "day month" pattern
		re = regexp.MustCompile(fmt.Sprintf(`(?i)^(\d{1,2})\s+(%s)$`, monthPattern))
		if matches := re.FindStringSubmatch(input); matches != nil {
			return incompleteDatePatterns[3].parser(ctx, matches)
		}

		// Try "day de month" pattern (Spanish: "3 de junio")
		re = regexp.MustCompile(fmt.Sprintf(`(?i)^(\d{1,2})\s+de\s+(%s)$`, monthPattern))
		if matches := re.FindStringSubmatch(input); matches != nil {
			return incompleteDatePatterns[3].parser(ctx, matches)
		}
	}

	// Try static patterns (year-only)
	for _, pattern := range incompleteDatePatterns {
		if pattern.regex != nil {
			matches := pattern.regex.FindStringSubmatch(input)
			if matches != nil {
				return pattern.parser(ctx, matches)
			}
		}
	}

	return time.Time{}, fmt.Errorf("no incomplete date pattern matched")
}

// buildMonthPatternForIncomplete creates a regex pattern with all month names
func buildMonthPatternForIncomplete(languages []*translations.Language) string {
	monthsMap := make(map[string]bool)

	for _, lang := range languages {
		for month := range lang.Months {
			monthsMap[regexp.QuoteMeta(month)] = true
		}
	}

	months := make([]string, 0, len(monthsMap))
	for month := range monthsMap {
		months = append(months, month)
	}

	if len(months) == 0 {
		return ""
	}

	return strings.Join(months, "|")
}
