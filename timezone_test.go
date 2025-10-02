package godateparser

import (
	"testing"
	"time"
)

// ============================================================================
// GoDateParser v0.3.0 Test Suite - Timezone Support
// ============================================================================
//
// This file contains tests for v0.3.0 features:
// - Timezone abbreviation parsing (EST, PST, GMT, etc.)
// - Timezone offset parsing (+05:00, -08:00)
// - Named offset parsing (UTC+5, GMT-8)
// - ISO 8601 with timezone
// - Date strings with timezone suffixes
// - Timezone conversion
// - Ambiguous timezone handling
//
// Total: 30+ tests (with sub-tests)
// ============================================================================

// ============================================================================
// TIMEZONE ABBREVIATION TESTS
// ============================================================================

func TestParseTimezone_Abbreviations(t *testing.T) {
	tests := []struct {
		name          string
		tz            string
		wantError     bool
		wantAmbiguous bool
	}{
		// UTC and GMT
		{"UTC", "UTC", false, false},
		{"GMT", "GMT", false, false},
		{"Z", "Z", false, false},

		// North American
		{"EST", "EST", false, false},
		{"EDT", "EDT", false, false},
		{"PST", "PST", false, false},
		{"PDT", "PDT", false, false},
		{"CST - ambiguous", "CST", false, true},
		{"MST", "MST", false, false},

		// European
		{"CET", "CET", false, false},
		{"BST", "BST", false, false},
		{"WET", "WET", false, false},

		// Asian
		{"JST", "JST", false, false},
		{"KST", "KST", false, false},
		{"SGT", "SGT", false, false},

		// Invalid
		{"Invalid abbreviation", "XYZ", true, false},
		{"Empty", "", true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tzInfo, err := ParseTimezone(tt.tz)

			if tt.wantError {
				if err == nil {
					t.Errorf("ParseTimezone(%q) expected error but got none", tt.tz)
				}
				return
			}

			if err != nil {
				t.Fatalf("ParseTimezone(%q) unexpected error: %v", tt.tz, err)
			}

			if tzInfo.Location == nil {
				t.Errorf("ParseTimezone(%q) returned nil location", tt.tz)
			}

			if tzInfo.Ambiguous != tt.wantAmbiguous {
				t.Errorf("ParseTimezone(%q) ambiguous = %v, want %v",
					tt.tz, tzInfo.Ambiguous, tt.wantAmbiguous)
			}
		})
	}
}

// ============================================================================
// TIMEZONE OFFSET TESTS
// ============================================================================

func TestParseTimezone_Offsets(t *testing.T) {
	tests := []struct {
		name        string
		offset      string
		wantSeconds int
		wantError   bool
	}{
		// Standard offsets
		{"+00:00", "+00:00", 0, false},
		{"+05:00", "+05:00", 5 * 3600, false},
		{"-08:00", "-08:00", -8 * 3600, false},
		{"+05:30", "+05:30", 5*3600 + 30*60, false},
		{"-03:30", "-03:30", -(3*3600 + 30*60), false},

		// Without colon
		{"+0530", "+0530", 5*3600 + 30*60, false},
		{"-0800", "-0800", -8 * 3600, false},

		// Edge cases
		{"+12:00", "+12:00", 12 * 3600, false},
		{"-12:00", "-12:00", -12 * 3600, false},
		{"+13:00", "+13:00", 13 * 3600, false},
		{"+14:00", "+14:00", 14 * 3600, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tzInfo, err := ParseTimezone(tt.offset)

			if tt.wantError {
				if err == nil {
					t.Errorf("ParseTimezone(%q) expected error but got none", tt.offset)
				}
				return
			}

			if err != nil {
				t.Fatalf("ParseTimezone(%q) unexpected error: %v", tt.offset, err)
			}

			if tzInfo.Offset != tt.wantSeconds {
				t.Errorf("ParseTimezone(%q) offset = %d seconds, want %d seconds",
					tt.offset, tzInfo.Offset, tt.wantSeconds)
			}
		})
	}
}

// ============================================================================
// NAMED OFFSET TESTS
// ============================================================================

func TestParseTimezone_NamedOffsets(t *testing.T) {
	tests := []struct {
		name        string
		offset      string
		wantSeconds int
	}{
		{"UTC+0", "UTC+0", 0},
		{"UTC+5", "UTC+5", 5 * 3600},
		{"UTC-8", "UTC-8", -8 * 3600},
		{"GMT+5", "GMT+5", 5 * 3600},
		{"GMT-8", "GMT-8", -8 * 3600},
		{"UTC+05:30", "UTC+05:30", 5*3600 + 30*60},
		{"UTC-03:30", "UTC-03:30", -(3*3600 + 30*60)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tzInfo, err := ParseTimezone(tt.offset)
			if err != nil {
				t.Fatalf("ParseTimezone(%q) unexpected error: %v", tt.offset, err)
			}

			if tzInfo.Offset != tt.wantSeconds {
				t.Errorf("ParseTimezone(%q) offset = %d seconds, want %d seconds",
					tt.offset, tzInfo.Offset, tt.wantSeconds)
			}
		})
	}
}

// ============================================================================
// TIMEZONE EXTRACTION TESTS
// ============================================================================

func TestExtractTimezone(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantDateStr string
		wantTZFound bool
	}{
		{
			"ISO with Z",
			"2024-12-31T10:30:00Z",
			"2024-12-31T10:30:00",
			true,
		},
		{
			"ISO with offset",
			"2024-12-31T10:30:00+05:00",
			"2024-12-31T10:30:00",
			true,
		},
		{
			"Date with abbreviation",
			"2024-12-31 10:30:00 EST",
			"2024-12-31 10:30:00",
			true,
		},
		{
			"Date with PST",
			"December 31, 2024 3:00 PM PST",
			"December 31, 2024 3:00 PM",
			true,
		},
		{
			"No timezone",
			"2024-12-31 10:30:00",
			"2024-12-31 10:30:00",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dateStr, tzInfo, _ := ExtractTimezone(tt.input)

			if dateStr != tt.wantDateStr {
				t.Errorf("ExtractTimezone(%q) dateStr = %q, want %q",
					tt.input, dateStr, tt.wantDateStr)
			}

			if tt.wantTZFound && tzInfo == nil {
				t.Errorf("ExtractTimezone(%q) expected timezone info but got nil", tt.input)
			}

			if !tt.wantTZFound && tzInfo != nil {
				t.Errorf("ExtractTimezone(%q) expected no timezone but got %v", tt.input, tzInfo)
			}
		})
	}
}

// ============================================================================
// INTEGRATION TESTS - PARSING DATES WITH TIMEZONES
// ============================================================================

func TestParseDate_WithTimezones(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		wantError    bool
		checkTZ      bool
		wantTZOffset int // Offset in seconds
	}{
		{
			"ISO with Z",
			"2024-12-31T10:30:00Z",
			false,
			true,
			0,
		},
		{
			"ISO with +05:00",
			"2024-12-31T10:30:00+05:00",
			false,
			true,
			5 * 3600,
		},
		{
			"ISO with -08:00",
			"2024-12-31T10:30:00-08:00",
			false,
			true,
			-8 * 3600,
		},
		{
			"Date with EST",
			"2024-12-31 10:30:00 EST",
			false,
			true,
			-5 * 3600, // EST is UTC-5
		},
		{
			"Date with PST",
			"2024-12-31 10:30:00 PST",
			false,
			true,
			-8 * 3600, // PST is UTC-8
		},
		{
			"Date with GMT",
			"2024-12-31 10:30:00 GMT",
			false,
			true,
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDate(tt.input, nil)

			if tt.wantError {
				if err == nil {
					t.Errorf("ParseDate(%q) expected error but got none", tt.input)
				}
				return
			}

			if err != nil {
				t.Fatalf("ParseDate(%q) unexpected error: %v", tt.input, err)
			}

			if tt.checkTZ {
				// Check that the timezone is correctly set
				_, offset := result.Zone()
				if offset != tt.wantTZOffset {
					t.Errorf("ParseDate(%q) timezone offset = %d seconds, want %d seconds",
						tt.input, offset, tt.wantTZOffset)
				}
			}
		})
	}
}

// ============================================================================
// TIMEZONE CONVERSION TESTS
// ============================================================================

func TestApplyTimezone(t *testing.T) {
	// Create a test time
	baseTime := time.Date(2024, 12, 31, 10, 30, 0, 0, time.UTC)

	// Test applying EST timezone
	estInfo, _ := ParseTimezone("EST")
	result := ApplyTimezone(baseTime, estInfo)

	// The time should now be interpreted as 10:30 EST
	if result.Location().String() != estInfo.Location.String() {
		t.Errorf("ApplyTimezone() location = %v, want %v",
			result.Location(), estInfo.Location)
	}

	// Test with nil timezone
	result2 := ApplyTimezone(baseTime, nil)
	if !result2.Equal(baseTime) {
		t.Errorf("ApplyTimezone() with nil should return original time")
	}
}

// ============================================================================
// BENCHMARKS
// ============================================================================

func BenchmarkParseTimezone_Abbreviation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ParseTimezone("EST")
	}
}

func BenchmarkParseTimezone_Offset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ParseTimezone("+05:00")
	}
}

func BenchmarkParseDate_WithTimezone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ParseDate("2024-12-31T10:30:00Z", nil)
	}
}

func BenchmarkExtractTimezone(b *testing.B) {
	input := "2024-12-31T10:30:00+05:00"
	for i := 0; i < b.N; i++ {
		_, _, _ = ExtractTimezone(input)
	}
}
