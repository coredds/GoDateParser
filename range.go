package godateparser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// DateRange represents a parsed date range with start and end dates
type DateRange struct {
	Start time.Time
	End   time.Time
	// MatchedText is the original text that was parsed
	MatchedText string
}

// rangePattern represents a date range parsing pattern
type rangePattern struct {
	regex  *regexp.Regexp
	parser func(*parserContext, []string) (*DateRange, error)
}

// Range patterns for Phase 3B
var rangePatterns = []*rangePattern{
	// "from X to Y" pattern
	{
		regex: regexp.MustCompile(`(?i)^from\s+(.+)\s+to\s+(.+)$`),
		parser: func(ctx *parserContext, matches []string) (*DateRange, error) {
			// Use smart splitting to handle multi-word dates
			startStr, endStr, err := splitRangeOnKeyword(ctx.input, "from", "to", ctx.settings)
			if err != nil {
				return nil, err
			}

			// Parse start date
			startDate, err := ParseDate(startStr, ctx.settings)
			if err != nil {
				return nil, fmt.Errorf("failed to parse start date '%s': %w", startStr, err)
			}

			// Parse end date
			endDate, err := ParseDate(endStr, ctx.settings)
			if err != nil {
				return nil, fmt.Errorf("failed to parse end date '%s': %w", endStr, err)
			}

			// Ensure start is before end
			if startDate.After(endDate) {
				return nil, &ErrInvalidDate{
					Reason: fmt.Sprintf("start date %v is after end date %v", startDate, endDate),
				}
			}

			return &DateRange{
				Start:       startDate,
				End:         endDate,
				MatchedText: ctx.input,
			}, nil
		},
	},
	// "between X and Y" pattern
	{
		regex: regexp.MustCompile(`(?i)^between\s+(.+)\s+and\s+(.+)$`),
		parser: func(ctx *parserContext, matches []string) (*DateRange, error) {
			// Use smart splitting to handle multi-word dates
			startStr, endStr, err := splitRangeOnKeyword(ctx.input, "between", "and", ctx.settings)
			if err != nil {
				return nil, err
			}

			// Parse start date
			startDate, err := ParseDate(startStr, ctx.settings)
			if err != nil {
				return nil, fmt.Errorf("failed to parse start date '%s': %w", startStr, err)
			}

			// Parse end date
			endDate, err := ParseDate(endStr, ctx.settings)
			if err != nil {
				return nil, fmt.Errorf("failed to parse end date '%s': %w", endStr, err)
			}

			// Ensure start is before end
			if startDate.After(endDate) {
				return nil, &ErrInvalidDate{
					Reason: fmt.Sprintf("start date %v is after end date %v", startDate, endDate),
				}
			}

			return &DateRange{
				Start:       startDate,
				End:         endDate,
				MatchedText: ctx.input,
			}, nil
		},
	},
	// "X - Y" pattern (with dashes)
	{
		regex: regexp.MustCompile(`^(.+?)\s+-\s+(.+)$`),
		parser: func(ctx *parserContext, matches []string) (*DateRange, error) {
			startStr := strings.TrimSpace(matches[1])
			endStr := strings.TrimSpace(matches[2])

			// Skip if it looks like a negative timezone offset or ISO date
			if regexp.MustCompile(`\d{4}-\d{2}-\d{2}`).MatchString(ctx.input) {
				return nil, fmt.Errorf("not a range, looks like ISO date")
			}

			// Parse start date
			startDate, err := ParseDate(startStr, ctx.settings)
			if err != nil {
				return nil, fmt.Errorf("failed to parse start date '%s': %w", startStr, err)
			}

			// Parse end date
			endDate, err := ParseDate(endStr, ctx.settings)
			if err != nil {
				return nil, fmt.Errorf("failed to parse end date '%s': %w", endStr, err)
			}

			// Ensure start is before end
			if startDate.After(endDate) {
				return nil, &ErrInvalidDate{
					Reason: fmt.Sprintf("start date %v is after end date %v", startDate, endDate),
				}
			}

			return &DateRange{
				Start:       startDate,
				End:         endDate,
				MatchedText: ctx.input,
			}, nil
		},
	},
	// "next N days/weeks/months/years" - returns range from now to now+N
	{
		regex: regexp.MustCompile(`(?i)^next\s+(\d+)\s+(day|week|month|year)s?$`),
		parser: func(ctx *parserContext, matches []string) (*DateRange, error) {
			amount, _ := strconv.Atoi(matches[1])
			unit := strings.ToLower(matches[2])

			base := ctx.settings.RelativeBase
			start := base
			var end time.Time

			switch unit {
			case "day":
				end = base.AddDate(0, 0, amount)
			case "week":
				end = base.AddDate(0, 0, amount*7)
			case "month":
				end = base.AddDate(0, amount, 0)
			case "year":
				end = base.AddDate(amount, 0, 0)
			}

			return &DateRange{
				Start:       start,
				End:         end,
				MatchedText: ctx.input,
			}, nil
		},
	},
	// "last N days/weeks/months/years" - returns range from now-N to now
	{
		regex: regexp.MustCompile(`(?i)^last\s+(\d+)\s+(day|week|month|year)s?$`),
		parser: func(ctx *parserContext, matches []string) (*DateRange, error) {
			amount, _ := strconv.Atoi(matches[1])
			unit := strings.ToLower(matches[2])

			base := ctx.settings.RelativeBase
			end := base
			var start time.Time

			switch unit {
			case "day":
				start = base.AddDate(0, 0, -amount)
			case "week":
				start = base.AddDate(0, 0, -amount*7)
			case "month":
				start = base.AddDate(0, -amount, 0)
			case "year":
				start = base.AddDate(-amount, 0, 0)
			}

			return &DateRange{
				Start:       start,
				End:         end,
				MatchedText: ctx.input,
			}, nil
		},
	},
}

// ParseDateRange parses a date range string and returns a DateRange
func ParseDateRange(input string, opts *Settings) (*DateRange, error) {
	if input == "" {
		return nil, &ErrEmptyInput{}
	}

	if opts == nil {
		opts = DefaultSettings()
	}

	settings := normalizeSettings(opts)

	ctx := &parserContext{
		input:    input,
		settings: settings,
	}

	// Try each range pattern
	var lastErr error
	for _, pattern := range rangePatterns {
		matches := pattern.regex.FindStringSubmatch(input)
		if matches != nil {
			result, err := pattern.parser(ctx, matches)
			if err == nil {
				return result, nil
			}
			// Save the error for debugging
			lastErr = err
			// Continue to next pattern if this one failed
		}
	}

	// If we had a parse error, return it with context
	if lastErr != nil {
		return nil, fmt.Errorf("range parsing failed: %w", lastErr)
	}

	return nil, &ErrInvalidFormat{
		Input:      input,
		Suggestion: "supported range formats: 'from X to Y', 'between X and Y', 'X - Y', 'next N days', 'last N weeks'",
	}
}

// GetDatesInRange returns all dates between start and end (inclusive) with the given step
// step is in days (e.g., 1 for every day, 7 for every week)
func GetDatesInRange(start, end time.Time, stepDays int) []time.Time {
	if stepDays <= 0 {
		stepDays = 1
	}

	var dates []time.Time
	current := start

	for !current.After(end) {
		dates = append(dates, current)
		current = current.AddDate(0, 0, stepDays)
	}

	return dates
}

// GetBusinessDaysInRange returns all business days (Monday-Friday) in the range
func GetBusinessDaysInRange(start, end time.Time) []time.Time {
	var dates []time.Time
	current := start

	for !current.After(end) {
		// Skip weekends (Saturday = 6, Sunday = 0)
		if current.Weekday() != time.Saturday && current.Weekday() != time.Sunday {
			dates = append(dates, current)
		}
		current = current.AddDate(0, 0, 1)
	}

	return dates
}

// DurationBetween returns the duration between two dates
func DurationBetween(start, end time.Time) time.Duration {
	return end.Sub(start)
}

// DaysBetween returns the number of days between two dates
func DaysBetween(start, end time.Time) int {
	duration := end.Sub(start)
	return int(duration.Hours() / 24)
}

// splitRangeOnKeyword intelligently splits a range string on the separator keyword
// by trying different split points and finding valid dates
func splitRangeOnKeyword(input, startKeyword, separatorKeyword string, settings *Settings) (string, string, error) {
	input = strings.TrimSpace(input)

	// Remove case-insensitive prefix
	inputLower := strings.ToLower(input)
	startKeywordLower := strings.ToLower(startKeyword)
	separatorKeywordLower := strings.ToLower(separatorKeyword)

	// Remove the start keyword (e.g., "from " or "between ")
	if !strings.HasPrefix(inputLower, startKeywordLower) {
		return "", "", fmt.Errorf("input doesn't start with '%s'", startKeyword)
	}

	// Find the start keyword and remove it
	prefixLen := len(startKeyword)
	remaining := strings.TrimSpace(input[prefixLen:])

	// Find all occurrences of the separator keyword (case-insensitive)
	remainingLower := strings.ToLower(remaining)
	separatorWithSpaces := " " + separatorKeywordLower + " "
	separatorPositions := []int{}
	searchStart := 0
	for {
		pos := strings.Index(remainingLower[searchStart:], separatorWithSpaces)
		if pos == -1 {
			break
		}
		actualPos := searchStart + pos
		separatorPositions = append(separatorPositions, actualPos)
		searchStart = actualPos + len(separatorWithSpaces)
	}

	if len(separatorPositions) == 0 {
		return "", "", fmt.Errorf("separator keyword '%s' not found", separatorKeyword)
	}

	// Try each split point from right to left (prefer later occurrences)
	// This handles cases like "from next Monday to next Friday" where "to" appears once
	for i := len(separatorPositions) - 1; i >= 0; i-- {
		splitPos := separatorPositions[i]
		startStr := strings.TrimSpace(remaining[:splitPos])
		// Skip the separator with spaces: splitPos points to the space before separator
		endPos := splitPos + len(separatorWithSpaces)
		endStr := strings.TrimSpace(remaining[endPos:])

		// Try to parse both parts - if both succeed, we found the right split
		_, startErr := ParseDate(startStr, settings)
		_, endErr := ParseDate(endStr, settings)

		if startErr == nil && endErr == nil {
			return startStr, endStr, nil
		}
	}

	// If no valid split found, return the first attempt
	if len(separatorPositions) > 0 {
		splitPos := separatorPositions[0]
		startStr := strings.TrimSpace(remaining[:splitPos])
		endPos := splitPos + len(separatorWithSpaces)
		endStr := strings.TrimSpace(remaining[endPos:])
		return startStr, endStr, nil
	}

	return "", "", fmt.Errorf("could not find valid date split")
}
