package godateparser

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/coredds/godateparser/translations"
)

// Common absolute date patterns
var absolutePatterns = []*absolutePattern{
	// Japanese/Chinese date format: 2024年12月31日
	{
		regex:  regexp.MustCompile(`^(\d{4})年(\d{1,2})月(\d{1,2})日$`),
		format: "YMD",
		parser: parseCJKDate,
	},
	// ISO 8601: 2024-12-31, 2024-12-31T10:30:00
	{
		regex:  regexp.MustCompile(`(?i)^(\d{4})-(\d{1,2})-(\d{1,2})(?:[T\s](\d{1,2}):(\d{1,2})(?::(\d{1,2}))?)?`),
		format: "YMD",
		parser: parseISO8601,
	},
	// ISO 8601 with 2-digit year: 24-12-31
	{
		regex:  regexp.MustCompile(`(?i)^(\d{2})-(\d{1,2})-(\d{1,2})$`),
		format: "YMD",
		parser: parseISO8601TwoDigitYear,
	},
	// Month name formats: "31 Dec 2024", "15 January 2024", "31 Dec 24"
	{
		regex:  regexp.MustCompile(`(?i)^(\d{1,2})\s+(Jan(?:uary)?|Feb(?:ruary)?|Mar(?:ch)?|Apr(?:il)?|May|Jun(?:e)?|Jul(?:y)?|Aug(?:ust)?|Sep(?:tember)?|Oct(?:ober)?|Nov(?:ember)?|Dec(?:ember)?)[,\s]*(\d{2,4})`),
		format: "DMY",
		parser: parseMonthName,
	},
	// Month name formats: "December 31, 2024", "Dec 31 2024", "Jan 15 2024", "Dec 31 24"
	{
		regex:  regexp.MustCompile(`(?i)^(Jan(?:uary)?|Feb(?:ruary)?|Mar(?:ch)?|Apr(?:il)?|May|Jun(?:e)?|Jul(?:y)?|Aug(?:ust)?|Sep(?:tember)?|Oct(?:ober)?|Nov(?:ember)?|Dec(?:ember)?)\s+(\d{1,2})[,\s]*(\d{2,4})`),
		format: "MDY",
		parser: parseMonthName,
	},
	// Numeric formats: 12/31/2024, 12-31-2024, 12/31/24, 12-31-24
	// Note: MDY and DMY both use the same regex - disambiguation happens in parseNumericDate
	{
		regex:  regexp.MustCompile(`^(\d{1,2})[/-](\d{1,2})[/-](\d{2,4})$`),
		format: "MDY",
		parser: parseNumericDate,
	},
}

type absolutePattern struct {
	regex  *regexp.Regexp
	format string
	parser func(*parserContext, []string) (time.Time, error)
}

// parseAbsolute attempts to parse absolute date formats.
func parseAbsolute(ctx *parserContext) (time.Time, error) {
	input := strings.TrimSpace(ctx.input)

	// Try to extract timezone first
	dateStr, tzInfo, _ := ExtractTimezone(input)

	// Try multi-language month name formats first
	if result, err := tryParseMultiLangMonthName(ctx, dateStr); err == nil {
		// Apply timezone if found
		if tzInfo != nil {
			result = ApplyTimezone(result, tzInfo)
		}
		return result, nil
	}

	// Try each pattern on the date part
	for _, pattern := range absolutePatterns {
		matches := pattern.regex.FindStringSubmatch(dateStr)
		if matches != nil {
			result, err := pattern.parser(ctx, matches)
			if err == nil {
				// Apply timezone if found
				if tzInfo != nil {
					result = ApplyTimezone(result, tzInfo)
				}
				return result, nil
			}
			// If we got a validation error, preserve it
			// (pattern matched but date was invalid)
			var invalidDateErr *ErrInvalidDate
			if errors.As(err, &invalidDateErr) {
				return time.Time{}, err
			}
			var ambiguousErr *ErrAmbiguousDate
			if errors.As(err, &ambiguousErr) {
				return time.Time{}, err
			}
		}
	}

	return time.Time{}, fmt.Errorf("no absolute date pattern matched")
}

// parseCJKDate handles Japanese/Chinese date format (YYYY年MM月DD日).
func parseCJKDate(ctx *parserContext, matches []string) (time.Time, error) {
	year, _ := strconv.Atoi(matches[1])
	month, _ := strconv.Atoi(matches[2])
	day, _ := strconv.Atoi(matches[3])

	// Validate date components
	if err := validateDateComponents(year, month, day); err != nil {
		return time.Time{}, err
	}

	loc := ctx.settings.PreferredTimezone
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)

	return date, nil
}

// parseISO8601 handles ISO 8601 format dates.
func parseISO8601(ctx *parserContext, matches []string) (time.Time, error) {
	year, _ := strconv.Atoi(matches[1])
	month, _ := strconv.Atoi(matches[2])
	day, _ := strconv.Atoi(matches[3])

	hour, minute, second := 0, 0, 0
	if len(matches) > 4 && matches[4] != "" {
		hour, _ = strconv.Atoi(matches[4])
	}
	if len(matches) > 5 && matches[5] != "" {
		minute, _ = strconv.Atoi(matches[5])
	}
	if len(matches) > 6 && matches[6] != "" {
		second, _ = strconv.Atoi(matches[6])
	}

	// Validate date and time components
	if err := validateDateTime(year, month, day, hour, minute, second); err != nil {
		return time.Time{}, err
	}

	loc := ctx.settings.PreferredTimezone
	date := time.Date(year, time.Month(month), day, hour, minute, second, 0, loc)

	return date, nil
}

// parseISO8601TwoDigitYear handles ISO 8601 format with 2-digit years.
func parseISO8601TwoDigitYear(ctx *parserContext, matches []string) (time.Time, error) {
	yy, _ := strconv.Atoi(matches[1])
	year := parseTwoDigitYear(yy)
	month, _ := strconv.Atoi(matches[2])
	day, _ := strconv.Atoi(matches[3])

	// Validate date components
	if err := validateDateComponents(year, month, day); err != nil {
		return time.Time{}, err
	}

	loc := ctx.settings.PreferredTimezone
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)

	return date, nil
}

// parseNumericDate handles numeric date formats like 12/31/2024 or 31/12/2024.
func parseNumericDate(ctx *parserContext, matches []string) (time.Time, error) {
	num1, _ := strconv.Atoi(matches[1])
	num2, _ := strconv.Atoi(matches[2])
	year, _ := strconv.Atoi(matches[3])

	// Handle 2-digit years
	if year < 100 {
		year = parseTwoDigitYear(year)
	}

	var month, day int

	// Try to detect date order from input if possible
	detectedOrder := ""
	if ctx.autoDetectDateOrder {
		detectedOrder = detectDateOrderFromInput(num1, num2, year)
	}

	// Check for ambiguity in strict mode
	if ctx.settings.Strict && ctx.autoDetectDateOrder && detectedOrder == "" {
		// Could not auto-detect and it's ambiguous
		if isAmbiguousDate(num1, num2, year) {
			return time.Time{}, &ErrAmbiguousDate{
				Input:  ctx.input,
				Reason: "numeric date is ambiguous (could be MDY or DMY format)",
			}
		}
	}

	// Use detected order if we found one, otherwise use configured order
	dateOrder := ctx.settings.DateOrder
	if detectedOrder != "" {
		dateOrder = detectedOrder
	}

	switch dateOrder {
	case "MDY":
		month, day = num1, num2
	case "DMY":
		day, month = num1, num2
	case "YMD":
		// This shouldn't happen for this pattern, but handle it
		month, day = num1, num2
	default:
		month, day = num1, num2
	}

	// Validate date components
	if err := validateDateComponents(year, month, day); err != nil {
		return time.Time{}, err
	}

	loc := ctx.settings.PreferredTimezone
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)

	return date, nil
}

// parseMonthName handles dates with month names.
func parseMonthName(ctx *parserContext, matches []string) (time.Time, error) {
	var day, year int
	var monthStr string

	// Determine which match is which based on pattern
	if _, err := strconv.Atoi(matches[1]); err == nil {
		// First match is a number (day), pattern: "31 Dec 2024"
		day, _ = strconv.Atoi(matches[1])
		monthStr = matches[2]
		year, _ = strconv.Atoi(matches[3])
	} else {
		// First match is month name, pattern: "Dec 31 2024"
		monthStr = matches[1]
		day, _ = strconv.Atoi(matches[2])
		year, _ = strconv.Atoi(matches[3])
	}

	// Handle 2-digit years
	if year < 100 {
		year = parseTwoDigitYear(year)
	}

	month := parseMonthString(monthStr)
	if month == 0 {
		return time.Time{}, fmt.Errorf("invalid month name: %s", monthStr)
	}

	// Validate date components
	if err := validateDateComponents(year, month, day); err != nil {
		return time.Time{}, err
	}

	loc := ctx.settings.PreferredTimezone
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)

	return date, nil
}

// parseMonthString converts month names to month numbers.
func parseMonthString(month string) int {
	month = strings.ToLower(month)
	monthMap := map[string]int{
		"jan": 1, "january": 1,
		"feb": 2, "february": 2,
		"mar": 3, "march": 3,
		"apr": 4, "april": 4,
		"may": 5,
		"jun": 6, "june": 6,
		"jul": 7, "july": 7,
		"aug": 8, "august": 8,
		"sep": 9, "september": 9,
		"oct": 10, "october": 10,
		"nov": 11, "november": 11,
		"dec": 12, "december": 12,
	}
	return monthMap[month]
}

// tryParseMultiLangMonthName attempts to parse month names in multiple languages.
// Supports formats like: "31 diciembre 2024", "15 de marzo de 2024", "marzo 15 2024"
func tryParseMultiLangMonthName(ctx *parserContext, input string) (time.Time, error) {
	input = strings.TrimSpace(input)

	// Build regex pattern with all month names from enabled languages
	monthPattern := buildMonthPattern(ctx.languages)
	if monthPattern == "" {
		return time.Time{}, fmt.Errorf("no languages available")
	}

	// Try different patterns
	patterns := []struct {
		regex *regexp.Regexp
		parse func([]string) (int, time.Month, int, error)
	}{
		// "31 diciembre 2024", "15 marzo 2024"
		{
			regex: regexp.MustCompile(fmt.Sprintf(`(?i)^(\d{1,2})\s+(%s)[,\s]+(\d{2,4})$`, monthPattern)),
			parse: func(matches []string) (int, time.Month, int, error) {
				day, _ := strconv.Atoi(matches[1])
				month := monthNameToNumberWithLangs(matches[2], ctx.languages)
				year, _ := strconv.Atoi(matches[3])
				return day, month, year, nil
			},
		},
		// "15 de marzo de 2024" (Spanish with two "de")
		{
			regex: regexp.MustCompile(fmt.Sprintf(`(?i)^(\d{1,2})\s+de\s+(%s)\s+de\s+(\d{2,4})$`, monthPattern)),
			parse: func(matches []string) (int, time.Month, int, error) {
				day, _ := strconv.Atoi(matches[1])
				month := monthNameToNumberWithLangs(matches[2], ctx.languages)
				year, _ := strconv.Atoi(matches[3])
				return day, month, year, nil
			},
		},
		// "3 de junio 2024" (Spanish with one "de")
		{
			regex: regexp.MustCompile(fmt.Sprintf(`(?i)^(\d{1,2})\s+de\s+(%s)\s+(\d{2,4})$`, monthPattern)),
			parse: func(matches []string) (int, time.Month, int, error) {
				day, _ := strconv.Atoi(matches[1])
				month := monthNameToNumberWithLangs(matches[2], ctx.languages)
				year, _ := strconv.Atoi(matches[3])
				return day, month, year, nil
			},
		},
		// "marzo 15 2024", "diciembre 31 2024"
		{
			regex: regexp.MustCompile(fmt.Sprintf(`(?i)^(%s)\s+(\d{1,2})[,\s]+(\d{2,4})$`, monthPattern)),
			parse: func(matches []string) (int, time.Month, int, error) {
				month := monthNameToNumberWithLangs(matches[1], ctx.languages)
				day, _ := strconv.Atoi(matches[2])
				year, _ := strconv.Atoi(matches[3])
				return day, month, year, nil
			},
		},
	}

	for _, pattern := range patterns {
		if matches := pattern.regex.FindStringSubmatch(input); matches != nil {
			day, month, year, err := pattern.parse(matches)
			if err != nil {
				continue
			}

			// Handle 2-digit years
			if year < 100 {
				year = parseTwoDigitYear(year)
			}

			if month == 0 {
				continue
			}

			// Validate date components
			if err := validateDateComponents(year, int(month), day); err != nil {
				return time.Time{}, err
			}

			loc := ctx.settings.PreferredTimezone
			return time.Date(year, month, day, 0, 0, 0, 0, loc), nil
		}
	}

	return time.Time{}, fmt.Errorf("no multi-language month pattern matched")
}

// buildMonthPattern creates a regex pattern with all month names from enabled languages.
func buildMonthPattern(languages []*translations.Language) string {
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
