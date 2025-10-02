package godateparser

import (
	"time"
)

// TwoDigitYearCutoff is the year cutoff for interpreting 2-digit years.
// Years below this are interpreted as 20xx, years >= are interpreted as 19xx.
// Default is 70 (so 69 = 2069, 70 = 1970, matching Python's dateparser).
const TwoDigitYearCutoff = 70

// parseTwoDigitYear interprets a 2-digit year as a full 4-digit year.
// Years 00-69 are interpreted as 2000-2069
// Years 70-99 are interpreted as 1970-1999
func parseTwoDigitYear(yy int) int {
	if yy < 0 || yy > 99 {
		return yy // Already a full year
	}

	if yy < TwoDigitYearCutoff {
		return 2000 + yy
	}
	return 1900 + yy
}

// validateDateComponents checks if the date components form a valid date.
// Returns an error if the date is invalid (e.g., Feb 31, month 13, etc.)
func validateDateComponents(year, month, day int) error {
	// Check month range
	if month < 1 || month > 12 {
		return &ErrInvalidDate{
			Year:   year,
			Month:  month,
			Day:    day,
			Reason: "month must be between 1 and 12",
		}
	}

	// Check day range
	if day < 1 || day > 31 {
		return &ErrInvalidDate{
			Year:   year,
			Month:  month,
			Day:    day,
			Reason: "day must be between 1 and 31",
		}
	}

	// Check if date is valid (e.g., Feb 31 is invalid)
	// Create the date and verify it matches what we specified
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	if date.Year() != year || int(date.Month()) != month || date.Day() != day {
		return &ErrInvalidDate{
			Year:   year,
			Month:  month,
			Day:    day,
			Reason: "invalid day for the given month/year",
		}
	}

	return nil
}

// validateDateTime validates date and time components together.
func validateDateTime(year, month, day, hour, minute, second int) error {
	// First validate date
	if err := validateDateComponents(year, month, day); err != nil {
		return err
	}

	// Validate time components
	if hour < 0 || hour > 23 {
		return &ErrInvalidDate{
			Year:   year,
			Month:  month,
			Day:    day,
			Reason: "hour must be between 0 and 23",
		}
	}

	if minute < 0 || minute > 59 {
		return &ErrInvalidDate{
			Year:   year,
			Month:  month,
			Day:    day,
			Reason: "minute must be between 0 and 59",
		}
	}

	if second < 0 || second > 59 {
		return &ErrInvalidDate{
			Year:   year,
			Month:  month,
			Day:    day,
			Reason: "second must be between 0 and 59",
		}
	}

	return nil
}

// isAmbiguousDate checks if a numeric date is ambiguous (e.g., 01/02/2024).
// Returns true if the date could be interpreted as both MDY and DMY.
func isAmbiguousDate(num1, num2, year int) bool {
	// Check if num1 could be month and num2 could be day (MDY)
	mdyPossible := num1 >= 1 && num1 <= 12 && num2 >= 1 && num2 <= 31

	// Check if num2 could be month and num1 could be day (DMY)
	dmyPossible := num2 >= 1 && num2 <= 12 && num1 >= 1 && num1 <= 31

	// Only ambiguous if both interpretations are possible and would result in different dates
	return mdyPossible && dmyPossible && num1 != num2
}

// detectDateOrderFromInput attempts to detect date order from the input.
// Returns "MDY", "DMY", or "" if ambiguous.
func detectDateOrderFromInput(num1, num2, year int) string {
	// If num1 > 12, it must be day (DMY format)
	if num1 > 12 {
		return "DMY"
	}

	// If num2 > 12, it must be day (MDY format)
	if num2 > 12 {
		return "MDY"
	}

	// Check if both interpretations would be valid
	mdyValid := validateDateComponents(year, num1, num2) == nil
	dmyValid := validateDateComponents(year, num2, num1) == nil

	if mdyValid && dmyValid {
		return "" // Ambiguous
	}

	if mdyValid {
		return "MDY"
	}

	if dmyValid {
		return "DMY"
	}

	return "" // Neither valid
}
