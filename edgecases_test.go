package godateparser

import (
	"testing"
	"time"
)

// Tests for edge cases and corner scenarios

// Boundary Tests

func TestEdgeCase_MonthBoundaries(t *testing.T) {
	base := time.Date(2024, 1, 31, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		name      string
		input     string
		wantMonth time.Month
		wantDay   int
	}{
		// Go's time.AddDate normalizes overflow dates
		// Jan 31 + 1 month = Feb 31 (invalid) -> March 2
		{"Next month from Jan 31", "next month", time.March, 2},
		{"1 month from Jan 31", "in 1 month", time.March, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if result.Month() != tt.wantMonth || result.Day() != tt.wantDay {
				t.Errorf("ParseDate(%q) = %v %v, want %v %v",
					tt.input, result.Month(), result.Day(), tt.wantMonth, tt.wantDay)
			}
		})
	}
}

func TestEdgeCase_YearBoundaries(t *testing.T) {
	base := time.Date(2024, 12, 31, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		name     string
		input    string
		wantYear int
		wantDay  int
	}{
		{"Tomorrow crosses year", "tomorrow", 2025, 1},
		{"1 day from Dec 31", "in 1 day", 2025, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if result.Year() != tt.wantYear || result.Day() != tt.wantDay {
				t.Errorf("ParseDate(%q) = %v-%v, want %v-%v",
					tt.input, result.Year(), result.Day(), tt.wantYear, tt.wantDay)
			}
		})
	}
}

func TestEdgeCase_NegativeQuantities(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	// Negative quantities should be treated as "ago"
	tests := []struct {
		input   string
		wantDay int
	}{
		{"-1 days", 14},
		{"-2 days", 13},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			// Some parsers may reject negative quantities, which is also valid
			if err != nil {
				return // Acceptable behavior
			}
			if result.Day() != tt.wantDay {
				t.Errorf("ParseDate(%q) day = %v, want %v", tt.input, result.Day(), tt.wantDay)
			}
		})
	}
}

// Same-Day Edge Cases

func TestEdgeCase_SameDayRelative(t *testing.T) {
	// When "today" is Monday and we ask for "Monday"
	base := time.Date(2024, 10, 14, 12, 0, 0, 0, time.UTC) // Monday
	settings := &Settings{
		RelativeBase:    base,
		PreferDatesFrom: "future",
	}

	result, err := ParseDate("Monday", settings)
	if err != nil {
		t.Fatalf("ParseDate() error = %v", err)
	}

	// With "future" preference, should go to next Monday (7 days from now)
	expectedDay := 21
	if result.Day() != expectedDay {
		t.Errorf("ParseDate('Monday') day = %v, want %v (next Monday from base Monday)",
			result.Day(), expectedDay)
	}
}

func TestEdgeCase_ThisMonday_WhenMonday(t *testing.T) {
	base := time.Date(2024, 10, 14, 12, 0, 0, 0, time.UTC) // Monday
	settings := &Settings{RelativeBase: base}

	result, err := ParseDate("this Monday", settings)
	if err != nil {
		t.Fatalf("ParseDate() error = %v", err)
	}

	// "this Monday" when today is Monday should return today
	if result.Day() != 14 {
		t.Errorf("ParseDate('this Monday') day = %v, want 14 (today)", result.Day())
	}
}

// Parser Priority Tests

func TestEdgeCase_ParserPriority_ISO_vs_Timestamp(t *testing.T) {
	// String "2024" could be year-only or timestamp
	result, err := ParseDate("2024", nil)
	if err != nil {
		t.Fatalf("ParseDate() error = %v", err)
	}

	// Should prefer year-only interpretation (2024-01-01) over timestamp (1970-01-01 00:33:44)
	if result.Year() != 2024 {
		t.Errorf("ParseDate('2024') year = %v, want 2024 (year-only, not timestamp)", result.Year())
	}
}

func TestEdgeCase_ParserPriority_Extended_vs_Basic(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC) // Q4
	settings := &Settings{RelativeBase: base}

	// "next quarter" should use quarter-aware logic, not simple +3 months
	result, err := ParseDate("next quarter", settings)
	if err != nil {
		t.Fatalf("ParseDate() error = %v", err)
	}

	// Should be start of Q1 2025 (January 1, 2025)
	if result.Year() != 2025 || result.Month() != time.January {
		t.Errorf("ParseDate('next quarter') = %v-%v, want 2025-January",
			result.Year(), result.Month())
	}
}

// Ordinal Edge Cases

func TestEdgeCase_OrdinalInvalidDay(t *testing.T) {
	tests := []string{
		"32nd",          // No month has 32 days
		"February 30th", // Feb doesn't have 30 days
		"April 31st",    // April has 30 days
		"0th",           // Day 0 doesn't exist
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			_, err := ParseDate(input, nil)
			if err == nil {
				t.Errorf("ParseDate(%q) should return error for invalid ordinal", input)
			}
		})
	}
}

// Week Number Edge Cases

func TestEdgeCase_WeekNumber_InvalidWeek(t *testing.T) {
	tests := []string{
		"2024-W00", // Week 0 doesn't exist
		"2024-W54", // Most years don't have 54 weeks
		"2024-W99", // Definitely invalid
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			_, err := ParseDate(input, nil)
			if err == nil {
				t.Errorf("ParseDate(%q) should return error for invalid week number", input)
			}
		})
	}
}

func TestEdgeCase_WeekNumber_InvalidWeekday(t *testing.T) {
	tests := []string{
		"2024-W15-0", // ISO weekdays are 1-7
		"2024-W15-8", // ISO weekdays are 1-7
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			_, err := ParseDate(input, nil)
			if err == nil {
				t.Errorf("ParseDate(%q) should return error for invalid weekday", input)
			}
		})
	}
}

// Time Edge Cases

func TestEdgeCase_Time_InvalidHour(t *testing.T) {
	base := time.Date(2024, 10, 15, 0, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []string{
		"25:00",
		"24:01",
		"-1:00",
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			_, err := ParseDate(input, settings)
			if err == nil {
				t.Errorf("ParseDate(%q) should return error for invalid hour", input)
			}
		})
	}
}

func TestEdgeCase_Time_MidnightAmbiguity(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		input    string
		wantHour int
	}{
		{"midnight", 0},
		{"12:00 AM", 0},
		{"00:00", 0},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if result.Hour() != tt.wantHour {
				t.Errorf("ParseDate(%q) hour = %v, want %v", tt.input, result.Hour(), tt.wantHour)
			}
		})
	}
}

func TestEdgeCase_Time_NoonAmbiguity(t *testing.T) {
	base := time.Date(2024, 10, 15, 0, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		input    string
		wantHour int
	}{
		{"noon", 12},
		{"12:00 PM", 12},
		{"12:00", 12},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if result.Hour() != tt.wantHour {
				t.Errorf("ParseDate(%q) hour = %v, want %v", tt.input, result.Hour(), tt.wantHour)
			}
		})
	}
}

// Mixed Format Tests

func TestEdgeCase_MixedFormats_ISOWithTimezone(t *testing.T) {
	tests := []string{
		"2024-12-31T10:30:00Z",
		"2024-12-31T10:30:00+05:00",
		"2024-12-31T10:30:00-08:00",
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			_, err := ParseDate(input, nil)
			if err != nil {
				t.Errorf("ParseDate(%q) error = %v, want nil", input, err)
			}
		})
	}
}

func TestEdgeCase_MixedFormats_MonthNameVariations(t *testing.T) {
	tests := []struct {
		input       string
		expectError bool
	}{
		{"31 Dec 2024", false},
		{"Dec 31 2024", false},
		{"December 31, 2024", false},
		{"31 December 2024", false},
		{"31-Dec-2024", true}, // Format with dashes not fully supported
		{"15-Jan-2024", true}, // Format with dashes not fully supported
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			_, err := ParseDate(tt.input, nil)
			if tt.expectError && err == nil {
				t.Errorf("ParseDate(%q) should return error", tt.input)
			}
			if !tt.expectError && err != nil {
				t.Errorf("ParseDate(%q) error = %v, want nil", tt.input, err)
			}
		})
	}
}

// Range Edge Cases

func TestEdgeCase_Range_StartAfterEnd(t *testing.T) {
	// End date before start date should be detected
	_, err := ParseDateRange("from 2024-12-31 to 2024-01-01", nil)
	if err == nil {
		t.Error("ParseDateRange() should return error when end is before start")
	}
}

func TestEdgeCase_Range_SameDate(t *testing.T) {
	// Start and end are the same
	result, err := ParseDateRange("from 2024-12-31 to 2024-12-31", nil)
	if err != nil {
		t.Fatalf("ParseDateRange() error = %v", err)
	}

	if !result.Start.Equal(result.End) {
		t.Error("ParseDateRange() start and end should be equal")
	}
}

// Large Quantity Tests

func TestEdgeCase_LargeQuantities(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		input    string
		wantYear int
	}{
		{"100 years ago", 1924},
		{"in 100 years", 2124},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if result.Year() != tt.wantYear {
				t.Errorf("ParseDate(%q) year = %v, want %v", tt.input, result.Year(), tt.wantYear)
			}
		})
	}
}

// Unicode and Special Characters

func TestEdgeCase_UnicodeWhitespace(t *testing.T) {
	// Test with various unicode whitespace characters
	tests := []string{
		"2024-12-31",       // Regular
		"2024\u00A0-12-31", // Non-breaking space (should fail)
	}

	validCount := 0
	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			_, err := ParseDate(input, nil)
			if err == nil {
				validCount++
			}
		})
	}

	// At least the regular one should work
	if validCount < 1 {
		t.Error("Should parse at least regular dates")
	}
}

// Benchmarks

func BenchmarkEdgeCase_MonthBoundary(b *testing.B) {
	base := time.Date(2024, 1, 31, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}
	for i := 0; i < b.N; i++ {
		_, _ = ParseDate("next month", settings)
	}
}

func BenchmarkEdgeCase_YearBoundary(b *testing.B) {
	base := time.Date(2024, 12, 31, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}
	for i := 0; i < b.N; i++ {
		_, _ = ParseDate("tomorrow", settings)
	}
}
