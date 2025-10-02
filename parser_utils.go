package godateparser

import (
	"strings"
	"time"

	"github.com/coredds/GoDateParser/translations"
)

// monthNameToNumber converts a month name to time.Month using translation system.
// It checks all enabled languages for the month name.
func monthNameToNumber(name string) time.Month {
	// Fallback to English-only map for backward compatibility
	monthMap := map[string]time.Month{
		"january": time.January, "jan": time.January,
		"february": time.February, "feb": time.February,
		"march": time.March, "mar": time.March,
		"april": time.April, "apr": time.April,
		"may":  time.May,
		"june": time.June, "jun": time.June,
		"july": time.July, "jul": time.July,
		"august": time.August, "aug": time.August,
		"september": time.September, "sep": time.September, "sept": time.September,
		"october": time.October, "oct": time.October,
		"november": time.November, "nov": time.November,
		"december": time.December, "dec": time.December,
	}
	return monthMap[strings.ToLower(name)]
}

// monthNameToNumberWithLangs converts a month name using specific languages.
func monthNameToNumberWithLangs(name string, langs []*translations.Language) time.Month {
	if month, ok := translations.ParseMonth(name, langs...); ok {
		return month
	}
	return 0
}

// weekdayNameToWeekday converts a weekday name using specific languages.
func weekdayNameToWeekday(name string, langs []*translations.Language) time.Weekday {
	if weekday, ok := translations.ParseWeekday(name, langs...); ok {
		return weekday
	}
	return 0
}
