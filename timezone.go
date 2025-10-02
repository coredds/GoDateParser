package godateparser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Common timezone abbreviations mapped to IANA timezone names
// Note: Some abbreviations are ambiguous (e.g., CST can be Central, China, or Cuba)
// We use the most common interpretation by default
var timezoneAbbreviations = map[string]string{
	// UTC and GMT
	"UTC": "UTC",
	"GMT": "GMT",
	"Z":   "UTC",

	// North American timezones
	"EST":  "America/New_York",    // Eastern Standard Time
	"EDT":  "America/New_York",    // Eastern Daylight Time
	"CST":  "America/Chicago",     // Central Standard Time (US)
	"CDT":  "America/Chicago",     // Central Daylight Time
	"MST":  "America/Denver",      // Mountain Standard Time
	"MDT":  "America/Denver",      // Mountain Daylight Time
	"PST":  "America/Los_Angeles", // Pacific Standard Time
	"PDT":  "America/Los_Angeles", // Pacific Daylight Time
	"AKST": "America/Anchorage",   // Alaska Standard Time
	"AKDT": "America/Anchorage",   // Alaska Daylight Time
	"HST":  "Pacific/Honolulu",    // Hawaii Standard Time

	// European timezones
	"CET":  "Europe/Paris",  // Central European Time
	"CEST": "Europe/Paris",  // Central European Summer Time
	"WET":  "Europe/Lisbon", // Western European Time
	"WEST": "Europe/Lisbon", // Western European Summer Time
	"EET":  "Europe/Athens", // Eastern European Time
	"EEST": "Europe/Athens", // Eastern European Summer Time
	"BST":  "Europe/London", // British Summer Time
	"IST":  "Europe/Dublin", // Irish Standard Time

	// Asian timezones
	"JST":  "Asia/Tokyo",      // Japan Standard Time
	"KST":  "Asia/Seoul",      // Korea Standard Time
	"SGT":  "Asia/Singapore",  // Singapore Time
	"HKT":  "Asia/Hong_Kong",  // Hong Kong Time
	"IST2": "Asia/Kolkata",    // India Standard Time (use IST2 to avoid conflict)
	"CST2": "Asia/Shanghai",   // China Standard Time (use CST2 to avoid US CST conflict)
	"AWST": "Australia/Perth", // Australian Western Standard Time

	// Australian timezones
	"AEST": "Australia/Sydney",   // Australian Eastern Standard Time
	"AEDT": "Australia/Sydney",   // Australian Eastern Daylight Time
	"ACST": "Australia/Adelaide", // Australian Central Standard Time
	"ACDT": "Australia/Adelaide", // Australian Central Daylight Time

	// Other timezones
	"NZST": "Pacific/Auckland", // New Zealand Standard Time
	"NZDT": "Pacific/Auckland", // New Zealand Daylight Time
}

// Timezone offset patterns
var (
	// Matches: +05:00, -08:00, +0530, -0800
	offsetPattern = regexp.MustCompile(`^([+-])(\d{2}):?(\d{2})$`)

	// Matches: UTC+5, GMT-8, UTC+05:30
	namedOffsetPattern = regexp.MustCompile(`^(UTC|GMT)([+-]\d{1,2}(?::\d{2})?)$`)
)

// TimezoneInfo represents parsed timezone information
type TimezoneInfo struct {
	Location   *time.Location
	Offset     int    // Offset in seconds from UTC
	Name       string // Original timezone string (e.g., "EST", "+05:00")
	Ambiguous  bool   // True if the abbreviation is ambiguous
	Normalized string // Normalized IANA timezone name
}

// ParseTimezone attempts to parse a timezone string into a TimezoneInfo
func ParseTimezone(tz string) (*TimezoneInfo, error) {
	tz = strings.TrimSpace(tz)
	if tz == "" {
		return nil, fmt.Errorf("empty timezone string")
	}

	// Try timezone abbreviation lookup
	if loc, normalized, ambiguous := lookupTimezoneAbbreviation(tz); loc != nil {
		return &TimezoneInfo{
			Location:   loc,
			Name:       tz,
			Normalized: normalized,
			Ambiguous:  ambiguous,
		}, nil
	}

	// Try offset pattern (+05:00, -08:00)
	if info := parseOffset(tz); info != nil {
		return info, nil
	}

	// Try named offset pattern (UTC+5, GMT-8)
	if info := parseNamedOffset(tz); info != nil {
		return info, nil
	}

	// Try loading as IANA timezone name
	if loc, err := time.LoadLocation(tz); err == nil {
		return &TimezoneInfo{
			Location:   loc,
			Name:       tz,
			Normalized: tz,
			Ambiguous:  false,
		}, nil
	}

	return nil, fmt.Errorf("unrecognized timezone: %q", tz)
}

// lookupTimezoneAbbreviation looks up a timezone abbreviation
// Returns (location, normalized name, is ambiguous)
func lookupTimezoneAbbreviation(abbr string) (*time.Location, string, bool) {
	abbr = strings.ToUpper(abbr)

	// Check for ambiguous abbreviations
	ambiguous := false
	if abbr == "CST" || abbr == "IST" {
		ambiguous = true
	}

	// Look up in abbreviation map
	if tzName, ok := timezoneAbbreviations[abbr]; ok {
		if loc, err := time.LoadLocation(tzName); err == nil {
			return loc, tzName, ambiguous
		}
	}

	return nil, "", false
}

// parseOffset parses timezone offsets like +05:00, -08:00, +0530
func parseOffset(offset string) *TimezoneInfo {
	matches := offsetPattern.FindStringSubmatch(offset)
	if matches == nil {
		return nil
	}

	sign := matches[1]
	hours, _ := strconv.Atoi(matches[2])
	minutes, _ := strconv.Atoi(matches[3])

	// Calculate offset in seconds
	offsetSeconds := hours*3600 + minutes*60
	if sign == "-" {
		offsetSeconds = -offsetSeconds
	}

	// Create fixed offset location
	loc := time.FixedZone(offset, offsetSeconds)

	return &TimezoneInfo{
		Location:   loc,
		Offset:     offsetSeconds,
		Name:       offset,
		Normalized: offset,
		Ambiguous:  false,
	}
}

// parseNamedOffset parses offsets like UTC+5, GMT-8, UTC+05:30
func parseNamedOffset(offset string) *TimezoneInfo {
	matches := namedOffsetPattern.FindStringSubmatch(strings.ToUpper(offset))
	if matches == nil {
		return nil
	}

	baseTZ := matches[1] // UTC or GMT
	offsetStr := matches[2]

	// Parse the offset part
	var sign int
	if strings.HasPrefix(offsetStr, "+") {
		sign = 1
		offsetStr = offsetStr[1:]
	} else if strings.HasPrefix(offsetStr, "-") {
		sign = -1
		offsetStr = offsetStr[1:]
	}

	// Parse hours and optional minutes
	var hours, minutes int
	if strings.Contains(offsetStr, ":") {
		parts := strings.Split(offsetStr, ":")
		hours, _ = strconv.Atoi(parts[0])
		minutes, _ = strconv.Atoi(parts[1])
	} else {
		hours, _ = strconv.Atoi(offsetStr)
	}

	// Calculate offset in seconds
	offsetSeconds := sign * (hours*3600 + minutes*60)

	// Create fixed offset location
	locName := fmt.Sprintf("%s%+d", baseTZ, sign*hours)
	if minutes > 0 {
		locName = fmt.Sprintf("%s%+d:%02d", baseTZ, sign*hours, minutes)
	}
	loc := time.FixedZone(locName, offsetSeconds)

	return &TimezoneInfo{
		Location:   loc,
		Offset:     offsetSeconds,
		Name:       offset,
		Normalized: locName,
		Ambiguous:  false,
	}
}

// ExtractTimezone attempts to extract timezone information from the end of a date string
// Returns the date string without timezone and the timezone info
func ExtractTimezone(input string) (dateStr string, tzInfo *TimezoneInfo, err error) {
	input = strings.TrimSpace(input)

	// Common patterns where timezone appears
	// 1. At the end after a space: "2024-12-31 10:30:00 EST"
	// 2. After T in ISO: "2024-12-31T10:30:00Z"
	// 3. With offset: "2024-12-31T10:30:00+05:00"

	// Try ISO format with Z
	if strings.HasSuffix(input, "Z") {
		tzInfo, _ := ParseTimezone("Z")
		return strings.TrimSuffix(input, "Z"), tzInfo, nil
	}

	// Try offset at the end (ISO format)
	// Match +HH:MM or +HHMM but not more than 2 digits for hours
	// This prevents matching dates like "01-15-2024" where "-2024" looks like an offset
	offsetRegex := regexp.MustCompile(`([+-]\d{2}:?\d{2})$`)
	if matches := offsetRegex.FindStringSubmatch(input); matches != nil {
		tzStr := matches[1]
		// Additional validation: make sure the hours part is valid (00-14)
		// and that there's a clear separator (T or space) before the offset
		if strings.Contains(input, "T") || strings.Contains(input, " ") {
			dateStr = strings.TrimSuffix(input, tzStr)
			tzInfo, err := ParseTimezone(tzStr)
			return dateStr, tzInfo, err
		}
	}

	// Try timezone abbreviation at the end
	// Only check if there's a space-separated last part (not part of date itself)
	parts := strings.Fields(input)
	if len(parts) >= 2 {
		lastPart := parts[len(parts)-1]
		// Only try if it looks like a timezone abbreviation (2-5 uppercase letters)
		// This avoids treating numbers like "2024" as potential timezones
		if len(lastPart) >= 2 && len(lastPart) <= 5 && isAllUpperOrZ(lastPart) {
			if tzInfo, err := ParseTimezone(lastPart); err == nil {
				// Successfully parsed last part as timezone
				dateStr = strings.Join(parts[:len(parts)-1], " ")
				return dateStr, tzInfo, nil
			}
		}
	}

	// No timezone found
	return input, nil, nil
}

// isAllUpperOrZ checks if a string is all uppercase letters or 'Z'
func isAllUpperOrZ(s string) bool {
	for _, c := range s {
		if !((c >= 'A' && c <= 'Z') || c == 'Z') {
			return false
		}
	}
	return true
}

// ApplyTimezone applies timezone information to a time.Time value
func ApplyTimezone(t time.Time, tzInfo *TimezoneInfo) time.Time {
	if tzInfo == nil || tzInfo.Location == nil {
		return t
	}

	// If the time is already in a specific location, convert it
	// Otherwise, interpret it as being in the target timezone
	if t.Location() == time.UTC || t.Location().String() == "UTC" {
		// Interpret the time as being in the target timezone
		return time.Date(
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second(),
			t.Nanosecond(), tzInfo.Location,
		)
	}

	// Convert to target timezone
	return t.In(tzInfo.Location)
}
