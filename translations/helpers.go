package translations

import (
	"strings"
	"time"
)

// ParseMonth attempts to parse a month name in any supported language.
func ParseMonth(input string, languages ...*Language) (time.Month, bool) {
	input = strings.ToLower(strings.TrimSpace(input))

	for _, lang := range languages {
		if month, ok := lang.Months[input]; ok {
			return month, true
		}
	}
	return 0, false
}

// ParseWeekday attempts to parse a weekday name in any supported language.
func ParseWeekday(input string, languages ...*Language) (time.Weekday, bool) {
	input = strings.ToLower(strings.TrimSpace(input))

	for _, lang := range languages {
		if weekday, ok := lang.Weekdays[input]; ok {
			return weekday, true
		}
	}
	return 0, false
}

// MatchesRelativeTerm checks if input matches any relative term in the given category.
func MatchesRelativeTerm(input string, terms []string) bool {
	input = strings.ToLower(strings.TrimSpace(input))

	for _, term := range terms {
		if term != "" && strings.ToLower(term) == input {
			return true
		}
	}
	return false
}

// ContainsRelativeTerm checks if input contains any relative term in the given category.
func ContainsRelativeTerm(input string, terms []string) bool {
	input = strings.ToLower(input)

	for _, term := range terms {
		if term != "" && strings.Contains(input, strings.ToLower(term)) {
			return true
		}
	}
	return false
}

// BuildTimeUnitPattern creates a regex pattern for time units in multiple languages.
func BuildTimeUnitPattern(languages []*Language) string {
	units := make(map[string]bool)

	for _, lang := range languages {
		if lang.RelativeTerms == nil {
			continue
		}

		addUnits := func(terms []string) {
			for _, term := range terms {
				if term != "" {
					units[term] = true
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
	}

	// Build pattern
	pattern := ""
	for unit := range units {
		if pattern != "" {
			pattern += "|"
		}
		pattern += unit
	}

	return pattern
}

// NormalizeTimeUnit normalizes a time unit string to English for internal processing.
func NormalizeTimeUnit(input string, languages []*Language) string {
	input = strings.ToLower(strings.TrimSpace(input))

	for _, lang := range languages {
		if lang.RelativeTerms == nil {
			continue
		}

		checkUnit := func(terms []string, normalized string) bool {
			for _, term := range terms {
				if strings.ToLower(term) == input {
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
	}

	return input
}

// GetWeekdayPattern builds a regex pattern for weekdays in multiple languages.
func GetWeekdayPattern(languages []*Language) string {
	weekdays := make(map[string]bool)

	for _, lang := range languages {
		for weekday := range lang.Weekdays {
			weekdays[weekday] = true
		}
	}

	pattern := ""
	for weekday := range weekdays {
		if pattern != "" {
			pattern += "|"
		}
		pattern += weekday
	}

	return pattern
}

// GetMonthPattern builds a regex pattern for months in multiple languages.
func GetMonthPattern(languages []*Language) string {
	months := make(map[string]bool)

	for _, lang := range languages {
		for month := range lang.Months {
			months[month] = true
		}
	}

	pattern := ""
	for month := range months {
		if pattern != "" {
			pattern += "|"
		}
		pattern += month
	}

	return pattern
}
