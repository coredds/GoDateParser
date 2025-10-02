package godateparser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	// Unix timestamp patterns (seconds or milliseconds)
	timestampRegex = regexp.MustCompile(`^\s*(\d{10,13})\s*$`)
)

// parseTimestamp attempts to parse Unix timestamps (seconds or milliseconds).
func parseTimestamp(ctx *parserContext) (time.Time, error) {
	input := strings.TrimSpace(ctx.input)

	matches := timestampRegex.FindStringSubmatch(input)
	if matches == nil {
		return time.Time{}, fmt.Errorf("not a valid timestamp")
	}

	timestampStr := matches[1]
	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid timestamp format: %w", err)
	}

	// Determine if it's seconds or milliseconds based on magnitude
	var result time.Time
	if timestamp > 1e12 {
		// Milliseconds
		result = time.Unix(0, timestamp*int64(time.Millisecond))
	} else {
		// Seconds
		result = time.Unix(timestamp, 0)
	}

	// Return in UTC to ensure consistent timezone handling
	return result.UTC(), nil
}
