package godateparser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Week number patterns - ISO week dates
// Examples: "W42", "Week 15 2024", "2024-W15", "2024W15"

type weekPattern struct {
	regex  *regexp.Regexp
	parser func(*parserContext, []string) (time.Time, error)
}

var weekPatterns = []*weekPattern{
	// ISO 8601 week format: "2024-W15", "2024W15"
	{
		regex: regexp.MustCompile(`^(\d{4})-?W(\d{1,2})$`),
		parser: func(_ *parserContext, matches []string) (time.Time, error) {
			year, _ := strconv.Atoi(matches[1])
			week, _ := strconv.Atoi(matches[2])

			if week < 1 || week > 53 {
				return time.Time{}, &ErrInvalidDate{
					Year:   year,
					Month:  0,
					Day:    0,
					Reason: fmt.Sprintf("week number %d out of range (1-53)", week),
				}
			}

			return getDateFromISOWeek(year, week, 1), nil // Monday of that week
		},
	},
	// Week with year: "Week 15 2024", "Week 42 2023"
	{
		regex: regexp.MustCompile(`(?i)^week\s+(\d{1,2})\s+(\d{4})$`),
		parser: func(_ *parserContext, matches []string) (time.Time, error) {
			week, _ := strconv.Atoi(matches[1])
			year, _ := strconv.Atoi(matches[2])

			if week < 1 || week > 53 {
				return time.Time{}, &ErrInvalidDate{
					Year:   year,
					Month:  0,
					Day:    0,
					Reason: fmt.Sprintf("week number %d out of range (1-53)", week),
				}
			}

			return getDateFromISOWeek(year, week, 1), nil
		},
	},
	// Week with year (alternate): "2024 Week 15"
	{
		regex: regexp.MustCompile(`(?i)^(\d{4})\s+week\s+(\d{1,2})$`),
		parser: func(_ *parserContext, matches []string) (time.Time, error) {
			year, _ := strconv.Atoi(matches[1])
			week, _ := strconv.Atoi(matches[2])

			if week < 1 || week > 53 {
				return time.Time{}, &ErrInvalidDate{
					Year:   year,
					Month:  0,
					Day:    0,
					Reason: fmt.Sprintf("week number %d out of range (1-53)", week),
				}
			}

			return getDateFromISOWeek(year, week, 1), nil
		},
	},
	// Week only (current year): "W42", "Week 15"
	{
		regex: regexp.MustCompile(`(?i)^w(eek\s+)?(\d{1,2})$`),
		parser: func(_ *parserContext, matches []string) (time.Time, error) {
			week, _ := strconv.Atoi(matches[2])

			if week < 1 || week > 53 {
				return time.Time{}, &ErrInvalidDate{
					Year:   0,
					Month:  0,
					Day:    0,
					Reason: fmt.Sprintf("week number %d out of range (1-53)", week),
				}
			}

			// Use current year or year from RelativeBase
			year := ctx.settings.RelativeBase.Year()
			return getDateFromISOWeek(year, week, 1), nil
		},
	},
	// ISO 8601 with weekday: "2024-W15-3" (Wednesday of week 15)
	{
		regex: regexp.MustCompile(`^(\d{4})-?W(\d{1,2})-?(\d)$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			year, _ := strconv.Atoi(matches[1])
			week, _ := strconv.Atoi(matches[2])
			weekday, _ := strconv.Atoi(matches[3])

			if week < 1 || week > 53 {
				return time.Time{}, &ErrInvalidDate{
					Year:   year,
					Month:  0,
					Day:    0,
					Reason: fmt.Sprintf("week number %d out of range (1-53)", week),
				}
			}

			if weekday < 1 || weekday > 7 {
				return time.Time{}, &ErrInvalidDate{
					Year:   year,
					Month:  0,
					Day:    0,
					Reason: fmt.Sprintf("weekday %d out of range (1-7)", weekday),
				}
			}

			return getDateFromISOWeek(year, week, weekday), nil
		},
	},
}

// getDateFromISOWeek converts an ISO week number to a date.
// ISO week date: Year, week number (1-53), and weekday (1=Monday, 7=Sunday)
// Algorithm from ISO 8601 standard.
func getDateFromISOWeek(year, week, weekday int) time.Time {
	// Find the Thursday of week 1 (ISO week starts on Monday, week 1 contains first Thursday)
	jan4 := time.Date(year, 1, 4, 0, 0, 0, 0, time.UTC) // Jan 4 is always in week 1

	// Find Monday of week 1
	weekdayJan4 := int(jan4.Weekday())
	if weekdayJan4 == 0 { // Sunday
		weekdayJan4 = 7
	}
	daysToMonday := 1 - weekdayJan4 // Days to go back to Monday
	mondayWeek1 := jan4.AddDate(0, 0, daysToMonday)

	// Calculate target date
	daysToAdd := (week-1)*7 + (weekday - 1)
	return mondayWeek1.AddDate(0, 0, daysToAdd)
}

// tryParseWeekNumber attempts to parse ISO week number patterns
func tryParseWeekNumber(ctx *parserContext) (time.Time, error) {
	input := strings.TrimSpace(ctx.input)

	for _, pattern := range weekPatterns {
		matches := pattern.regex.FindStringSubmatch(input)
		if matches != nil {
			return pattern.parser(ctx, matches)
		}
	}

	return time.Time{}, fmt.Errorf("no week number pattern matched")
}
