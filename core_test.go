package godateparser

import (
	"errors"
	"testing"
	"time"
)

// Tests for core functionality (API, settings, error handling, validation, extraction)

// API and Settings Tests

func TestDefaultSettings(t *testing.T) {
	settings := DefaultSettings()

	// DateOrder may be set to "MDY" by default
	if len(settings.EnableParsers) != 8 {
		t.Errorf("DefaultSettings() EnableParsers length = %d, want 8", len(settings.EnableParsers))
	}
	if settings.Strict {
		t.Error("DefaultSettings() Strict = true, want false")
	}
	if settings.PreferDatesFrom != "future" {
		t.Errorf("DefaultSettings() PreferDatesFrom = %q, want 'future'", settings.PreferDatesFrom)
	}
}

func TestParseDate_NilSettings(t *testing.T) {
	// Should work with nil settings (use defaults)
	result, err := ParseDate("2024-12-31", nil)
	if err != nil {
		t.Fatalf("ParseDate() with nil settings error = %v", err)
	}
	if result.Year() != 2024 || result.Month() != 12 || result.Day() != 31 {
		t.Errorf("ParseDate() = %v, want 2024-12-31", result)
	}
}

func TestParseDate_CustomRelativeBase(t *testing.T) {
	base := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	result, err := ParseDate("tomorrow", settings)
	if err != nil {
		t.Fatalf("ParseDate() error = %v", err)
	}

	expected := time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	if !result.Equal(expected) {
		t.Errorf("ParseDate() = %v, want %v", result, expected)
	}
}

// Error Handling Tests

func TestParseDate_EmptyInput(t *testing.T) {
	_, err := ParseDate("", nil)
	if err == nil {
		t.Fatal("ParseDate() with empty input should return error")
	}

	var emptyErr *ErrEmptyInput
	if !errors.As(err, &emptyErr) {
		t.Errorf("ParseDate() error type = %T, want *ErrEmptyInput", err)
	}
}

func TestParseDate_InvalidFormat(t *testing.T) {
	_, err := ParseDate("not a date at all", nil)
	if err == nil {
		t.Fatal("ParseDate() with invalid format should return error")
	}

	var formatErr *ErrInvalidFormat
	if !errors.As(err, &formatErr) {
		t.Errorf("ParseDate() error type = %T, want *ErrInvalidFormat", err)
	}
}

func TestParseDate_InvalidDate(t *testing.T) {
	tests := []string{
		"2024-02-30", // Feb doesn't have 30 days
		"2024-13-01", // Invalid month
		"2024-00-01", // Invalid month
		"2024-01-32", // Invalid day
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			_, err := ParseDate(input, nil)
			if err == nil {
				t.Fatalf("ParseDate(%q) should return error", input)
			}

			var dateErr *ErrInvalidDate
			if !errors.As(err, &dateErr) {
				t.Errorf("ParseDate(%q) error type = %T, want *ErrInvalidDate", input, err)
			}
		})
	}
}

func TestParseDate_AmbiguousDate_StrictMode(t *testing.T) {
	settings := &Settings{
		Strict: true,
		// Don't set DateOrder - let it try to auto-detect
	}

	// This date is ambiguous (could be Jan 2 or Feb 1)
	_, err := ParseDate("01/02/2024", settings)
	if err == nil {
		t.Fatal("ParseDate() with ambiguous date in strict mode should return error")
	}

	var ambigErr *ErrAmbiguousDate
	if !errors.As(err, &ambigErr) {
		t.Errorf("ParseDate() error type = %T, want *ErrAmbiguousDate", err)
	}
}

func TestParseDate_AmbiguousDate_NonStrictMode(t *testing.T) {
	settings := &Settings{
		Strict: false,
	}

	// In non-strict mode, should pick one interpretation without error
	_, err := ParseDate("01/02/2024", settings)
	if err != nil {
		t.Errorf("ParseDate() in non-strict mode should not error on ambiguous date, got: %v", err)
	}
}

// Validation Tests

func TestValidation_InvalidMonth(t *testing.T) {
	tests := []string{
		"2024-00-15",
		"2024-13-15",
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			_, err := ParseDate(input, nil)
			if err == nil {
				t.Errorf("ParseDate(%q) should return error for invalid month", input)
			}
		})
	}
}

func TestValidation_InvalidDay(t *testing.T) {
	tests := []struct {
		input  string
		reason string
	}{
		{"2024-01-00", "day 0"},
		{"2024-01-32", "day 32"},
		{"2024-02-30", "Feb 30"},
		{"2024-04-31", "Apr 31"},
		{"2024-06-31", "Jun 31"},
		{"2024-09-31", "Sep 31"},
		{"2024-11-31", "Nov 31"},
	}

	for _, tt := range tests {
		t.Run(tt.reason, func(t *testing.T) {
			_, err := ParseDate(tt.input, nil)
			if err == nil {
				t.Errorf("ParseDate(%q) should return error for %s", tt.input, tt.reason)
			}
		})
	}
}

func TestValidation_LeapYear(t *testing.T) {
	tests := []struct {
		input  string
		valid  bool
		isLeap bool
	}{
		{"2024-02-29", true, true},   // 2024 is leap year
		{"2023-02-29", false, false}, // 2023 is not leap year
		{"2000-02-29", true, true},   // 2000 is leap year
		{"1900-02-29", false, false}, // 1900 is not leap year (divisible by 100 but not 400)
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			_, err := ParseDate(tt.input, nil)
			if tt.valid && err != nil {
				t.Errorf("ParseDate(%q) error = %v, want nil (leap year: %v)", tt.input, err, tt.isLeap)
			}
			if !tt.valid && err == nil {
				t.Errorf("ParseDate(%q) should return error (not leap year)", tt.input)
			}
		})
	}
}

func TestValidation_TimeComponents(t *testing.T) {
	tests := []struct {
		input     string
		expectErr bool
	}{
		{"2024-12-31T00:00:00", false},
		{"2024-12-31T23:59:59", false},
		{"2024-12-31T24:00:00", true}, // Invalid hour
		{"2024-12-31T12:60:00", true}, // Invalid minute
		{"2024-12-31T12:00:60", true}, // Invalid second
		// Note: Negative hour in ISO format may be parsed as timezone offset
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			_, err := ParseDate(tt.input, nil)
			if tt.expectErr && err == nil {
				t.Errorf("ParseDate(%q) should return error for invalid time", tt.input)
			}
			if !tt.expectErr && err != nil {
				t.Errorf("ParseDate(%q) error = %v, want nil", tt.input, err)
			}
		})
	}
}

// Timestamp Tests

func TestParseDate_UnixTimestamp(t *testing.T) {
	tests := []struct {
		input string
		want  time.Time
	}{
		{"1609459200", time.Unix(1609459200, 0).UTC()},    // 2021-01-01 00:00:00 UTC
		{"1640995200", time.Unix(1640995200, 0).UTC()},    // 2022-01-01 00:00:00 UTC
		{"1672531200", time.Unix(1672531200, 0).UTC()},    // 2023-01-01 00:00:00 UTC
		{"1609459200000", time.Unix(1609459200, 0).UTC()}, // Milliseconds
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, nil)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if !result.Equal(tt.want) {
				t.Errorf("ParseDate() = %v, want %v", result, tt.want)
			}
		})
	}
}

// Extraction Tests

func TestExtractDates_MultipleOccurrences(t *testing.T) {
	text := "Meeting on 2024-12-31 and follow-up on 2025-01-15."
	results, err := ExtractDates(text, nil)
	if err != nil {
		t.Fatalf("ExtractDates() error = %v", err)
	}

	if len(results) != 2 {
		t.Fatalf("ExtractDates() found %d dates, want 2", len(results))
	}

	// Check first date
	if results[0].Date.Year() != 2024 || results[0].Date.Month() != 12 || results[0].Date.Day() != 31 {
		t.Errorf("First date = %v, want 2024-12-31", results[0].Date)
	}

	// Check second date
	if results[1].Date.Year() != 2025 || results[1].Date.Month() != 1 || results[1].Date.Day() != 15 {
		t.Errorf("Second date = %v, want 2025-01-15", results[1].Date)
	}
}

func TestExtractDates_NoDates(t *testing.T) {
	text := "This text has no dates in it at all."
	results, err := ExtractDates(text, nil)
	if err != nil {
		t.Fatalf("ExtractDates() error = %v", err)
	}

	if len(results) != 0 {
		t.Errorf("ExtractDates() found %d dates, want 0", len(results))
	}
}

func TestExtractDates_VariousFormats(t *testing.T) {
	text := "ISO: 2024-12-31, US: 12/31/2024, Named: December 31, 2024"
	results, err := ExtractDates(text, nil)
	if err != nil {
		t.Fatalf("ExtractDates() error = %v", err)
	}

	if len(results) < 3 {
		t.Errorf("ExtractDates() found %d dates, want at least 3", len(results))
	}
}

// Case Sensitivity Tests

func TestParseDate_CaseInsensitive(t *testing.T) {
	tests := []string{
		"DECEMBER 31, 2024",
		"december 31, 2024",
		"December 31, 2024",
		"YESTERDAY",
		"yesterday",
		"YeStErDaY",
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			_, err := ParseDate(input, nil)
			if err != nil {
				t.Errorf("ParseDate(%q) should be case insensitive, got error: %v", input, err)
			}
		})
	}
}

// Whitespace Tests

func TestParseDate_ExtraWhitespace(t *testing.T) {
	tests := []struct {
		input string
		valid bool
	}{
		{"  2024-12-31  ", true},
		{"\t2024-12-31\t", true},
		{"2024 - 12 - 31", false}, // Too much internal whitespace for ISO format
		{"December  31,  2024", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			_, err := ParseDate(tt.input, nil)
			if tt.valid && err != nil {
				t.Errorf("ParseDate(%q) error = %v, want nil", tt.input, err)
			}
			if !tt.valid && err == nil {
				t.Errorf("ParseDate(%q) should return error", tt.input)
			}
		})
	}
}

// Benchmarks

func BenchmarkParseDate_ISO8601(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ParseDate("2024-12-31", nil)
	}
}

func BenchmarkParseDate_Relative(b *testing.B) {
	settings := DefaultSettings()
	for i := 0; i < b.N; i++ {
		_, _ = ParseDate("yesterday", settings)
	}
}

func BenchmarkParseDate_Timestamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ParseDate("1609459200", nil)
	}
}

func BenchmarkExtractDates(b *testing.B) {
	text := "Meeting on 2024-12-31 and follow-up on 2025-01-15."
	for i := 0; i < b.N; i++ {
		_, _ = ExtractDates(text, nil)
	}
}
