package godateparser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/coredds/godateparser/translations"
)

// Ordinal date patterns - dates with ordinal numbers
// Examples: "1st", "2nd", "3rd", "21st", "June 3rd 2024", "3rd of June"

type ordinalDatePattern struct {
	regex  *regexp.Regexp
	parser func(*parserContext, []string) (time.Time, error)
}

var ordinalDatePatterns = []*ordinalDatePattern{
	// Ordinal only: "1st", "23rd" (day of current/next month)
	{
		regex: regexp.MustCompile(`(?i)^(\d{1,2})(st|nd|rd|th)$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			day, _ := strconv.Atoi(matches[1])

			base := ctx.settings.RelativeBase
			currentYear := base.Year()
			currentMonth := base.Month()
			currentDay := base.Day()

			// Determine month/year based on PreferDatesFrom setting
			month := currentMonth
			year := currentYear

			if ctx.settings.PreferDatesFrom == "past" {
				// If day is after current day, use last month
				if day > currentDay {
					month--
					if month < 1 {
						month = 12
						year--
					}
				}
			} else {
				// Default "future" behavior
				// If day is before current day, use next month
				if day < currentDay {
					month++
					if month > 12 {
						month = 1
						year++
					}
				}
			}

			// Validate day for the target month
			if err := validateDateComponents(year, int(month), day); err != nil {
				return time.Time{}, err
			}

			loc := ctx.settings.PreferredTimezone
			return time.Date(year, month, day, 0, 0, 0, 0, loc), nil
		},
	},
	// Month and ordinal: "June 3rd", "junio 3"
	{
		regex: nil, // Handled dynamically in tryParseOrdinalDate
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			monthName := strings.ToLower(matches[1])
			month := monthNameToNumberWithLangs(monthName, ctx.languages)
			day, _ := strconv.Atoi(matches[2])

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

			// Validate day
			if err := validateDateComponents(year, int(month), day); err != nil {
				return time.Time{}, err
			}

			loc := ctx.settings.PreferredTimezone
			return time.Date(year, month, day, 0, 0, 0, 0, loc), nil
		},
	},
	// Ordinal and month (no "of"): "3rd June", "21st March"
	{
		regex: regexp.MustCompile(`(?i)^(\d{1,2})(st|nd|rd|th)\s+(january|february|march|april|may|june|july|august|september|october|november|december|jan|feb|mar|apr|may|jun|jul|aug|sep|sept|oct|nov|dec)$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			day, _ := strconv.Atoi(matches[1])
			monthName := strings.ToLower(matches[3])
			month := monthNameToNumberWithLangs(monthName, ctx.languages)

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

			// Validate day
			if err := validateDateComponents(year, int(month), day); err != nil {
				return time.Time{}, err
			}

			loc := ctx.settings.PreferredTimezone
			return time.Date(year, month, day, 0, 0, 0, 0, loc), nil
		},
	},
	// Ordinal and month (with "of"): "3rd of June"
	{
		regex: regexp.MustCompile(`(?i)^(\d{1,2})(st|nd|rd|th)\s+of\s+(january|february|march|april|may|june|july|august|september|october|november|december|jan|feb|mar|apr|may|jun|jul|aug|sep|sept|oct|nov|dec)$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			day, _ := strconv.Atoi(matches[1])
			monthName := strings.ToLower(matches[3])
			month := monthNameToNumberWithLangs(monthName, ctx.languages)

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

			// Validate day
			if err := validateDateComponents(year, int(month), day); err != nil {
				return time.Time{}, err
			}

			loc := ctx.settings.PreferredTimezone
			return time.Date(year, month, day, 0, 0, 0, 0, loc), nil
		},
	},
	// Full date with ordinal: "June 3rd 2024"
	{
		regex: regexp.MustCompile(`(?i)^(january|february|march|april|may|june|july|august|september|october|november|december|jan|feb|mar|apr|may|jun|jul|aug|sep|sept|oct|nov|dec)\s+(\d{1,2})(st|nd|rd|th)\s+(\d{2,4})$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			monthName := strings.ToLower(matches[1])
			month := monthNameToNumberWithLangs(monthName, ctx.languages)
			day, _ := strconv.Atoi(matches[2])
			year, _ := strconv.Atoi(matches[4])

			// Handle 2-digit years
			if year < 100 {
				year = parseTwoDigitYear(year)
			}

			// Validate day
			if err := validateDateComponents(year, int(month), day); err != nil {
				return time.Time{}, err
			}

			loc := ctx.settings.PreferredTimezone
			return time.Date(year, month, day, 0, 0, 0, 0, loc), nil
		},
	},
	// Ordinal, month, year (no "of"): "3rd June 2024", "1st January 2025"
	{
		regex: regexp.MustCompile(`(?i)^(\d{1,2})(st|nd|rd|th)\s+(january|february|march|april|may|june|july|august|september|october|november|december|jan|feb|mar|apr|may|jun|jul|aug|sep|sept|oct|nov|dec)\s+(\d{2,4})$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			day, _ := strconv.Atoi(matches[1])
			monthName := strings.ToLower(matches[3])
			month := monthNameToNumberWithLangs(monthName, ctx.languages)
			year, _ := strconv.Atoi(matches[4])

			// Handle 2-digit years
			if year < 100 {
				year = parseTwoDigitYear(year)
			}

			// Validate day
			if err := validateDateComponents(year, int(month), day); err != nil {
				return time.Time{}, err
			}

			loc := ctx.settings.PreferredTimezone
			return time.Date(year, month, day, 0, 0, 0, 0, loc), nil
		},
	},
	// Ordinal, month, year (with "of"): "3rd of June 2024"
	{
		regex: regexp.MustCompile(`(?i)^(\d{1,2})(st|nd|rd|th)\s+of\s+(january|february|march|april|may|june|july|august|september|october|november|december|jan|feb|mar|apr|may|jun|jul|aug|sep|sept|oct|nov|dec)\s+(\d{2,4})$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			day, _ := strconv.Atoi(matches[1])
			monthName := strings.ToLower(matches[3])
			month := monthNameToNumberWithLangs(monthName, ctx.languages)
			year, _ := strconv.Atoi(matches[4])

			// Handle 2-digit years
			if year < 100 {
				year = parseTwoDigitYear(year)
			}

			// Validate day
			if err := validateDateComponents(year, int(month), day); err != nil {
				return time.Time{}, err
			}

			loc := ctx.settings.PreferredTimezone
			return time.Date(year, month, day, 0, 0, 0, 0, loc), nil
		},
	},
}

// tryParseOrdinalDate attempts to parse ordinal date patterns
func tryParseOrdinalDate(ctx *parserContext) (time.Time, error) {
	input := strings.TrimSpace(ctx.input)

	// Build month pattern from enabled languages
	monthPattern := buildMonthPatternForOrdinal(ctx.languages)

	// Try dynamic month-ordinal patterns
	if monthPattern != "" {
		// "June 3rd" or "junio 3"
		re := regexp.MustCompile(fmt.Sprintf(`(?i)^(%s)\s+(\d{1,2})(st|nd|rd|th)?$`, monthPattern))
		if matches := re.FindStringSubmatch(input); matches != nil {
			return ordinalDatePatterns[1].parser(ctx, matches)
		}

		// "3rd of June" or "3 de junio"
		re = regexp.MustCompile(fmt.Sprintf(`(?i)^(\d{1,2})(st|nd|rd|th)\s+(?:of|de)\s+(%s)$`, monthPattern))
		if matches := re.FindStringSubmatch(input); matches != nil {
			return ordinalDatePatterns[2].parser(ctx, matches)
		}

		// "3rd June" or "3 junio"
		re = regexp.MustCompile(fmt.Sprintf(`(?i)^(\d{1,2})(st|nd|rd|th)?\s+(%s)$`, monthPattern))
		if matches := re.FindStringSubmatch(input); matches != nil {
			return ordinalDatePatterns[3].parser(ctx, matches)
		}

		// "June 3rd 2024" or "junio 3 2024"
		re = regexp.MustCompile(fmt.Sprintf(`(?i)^(%s)\s+(\d{1,2})(st|nd|rd|th)?\s+(\d{2,4})$`, monthPattern))
		if matches := re.FindStringSubmatch(input); matches != nil {
			return ordinalDatePatterns[4].parser(ctx, matches)
		}

		// "3rd of June 2024" or "3 de junio 2024"
		re = regexp.MustCompile(fmt.Sprintf(`(?i)^(\d{1,2})(st|nd|rd|th)\s+(?:of|de)\s+(%s)\s+(\d{2,4})$`, monthPattern))
		if matches := re.FindStringSubmatch(input); matches != nil {
			return ordinalDatePatterns[5].parser(ctx, matches)
		}

		// "3rd June 2024" or "3 junio 2024"
		re = regexp.MustCompile(fmt.Sprintf(`(?i)^(\d{1,2})(st|nd|rd|th)?\s+(%s)\s+(\d{2,4})$`, monthPattern))
		if matches := re.FindStringSubmatch(input); matches != nil {
			return ordinalDatePatterns[6].parser(ctx, matches)
		}
	}

	// Try static patterns (ordinal-only)
	for _, pattern := range ordinalDatePatterns {
		if pattern.regex != nil {
			matches := pattern.regex.FindStringSubmatch(input)
			if matches != nil {
				return pattern.parser(ctx, matches)
			}
		}
	}

	return time.Time{}, fmt.Errorf("no ordinal date pattern matched")
}

// buildMonthPatternForOrdinal creates a regex pattern with all month names
func buildMonthPatternForOrdinal(languages []*translations.Language) string {
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
