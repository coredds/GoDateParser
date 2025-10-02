package godateparser

import (
	"errors"
	"testing"
	"time"
)

// ============================================================================
// GoDateParser v1.0.0 Phase 3B Test Suite - Date Range Parsing
// ============================================================================
//
// This file contains tests for Phase 3B range parsing features:
// - from...to patterns
// - between...and patterns
// - X - Y patterns
// - Duration-based ranges (next/last N days/weeks/months)
// - Helper functions for range manipulation
//
// Total: 25+ tests
// ============================================================================

// ============================================================================
// FROM...TO PATTERN TESTS
// ============================================================================

func TestParseRange_FromTo(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		name      string
		input     string
		wantStart time.Time
		wantEnd   time.Time
	}{
		{
			"from yesterday to tomorrow",
			"from yesterday to tomorrow",
			time.Date(2024, 10, 14, 12, 0, 0, 0, time.UTC),
			time.Date(2024, 10, 16, 12, 0, 0, 0, time.UTC),
		},
		{
			"from 2024-01-01 to 2024-12-31",
			"from 2024-01-01 to 2024-12-31",
			time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			"from Jan 1 to Dec 31",
			"from Jan 1 2024 to Dec 31 2024",
			time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			"from last Monday to next Friday",
			"from last Monday to next Friday",
			time.Date(2024, 10, 14, 12, 0, 0, 0, time.UTC), // Last Monday from Tuesday (Oct 14)
			time.Date(2024, 10, 18, 12, 0, 0, 0, time.UTC), // Next Friday from Tuesday (Oct 18)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDateRange(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDateRange(%q) error = %v", tt.input, err)
			}

			if !result.Start.Equal(tt.wantStart) {
				t.Errorf("ParseDateRange(%q) start = %v, want %v", tt.input, result.Start, tt.wantStart)
			}
			if !result.End.Equal(tt.wantEnd) {
				t.Errorf("ParseDateRange(%q) end = %v, want %v", tt.input, result.End, tt.wantEnd)
			}
		})
	}
}

// ============================================================================
// BETWEEN...AND PATTERN TESTS
// ============================================================================

func TestParseRange_BetweenAnd(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		name      string
		input     string
		wantStart time.Time
		wantEnd   time.Time
	}{
		{
			"between yesterday and tomorrow",
			"between yesterday and tomorrow",
			time.Date(2024, 10, 14, 12, 0, 0, 0, time.UTC),
			time.Date(2024, 10, 16, 12, 0, 0, 0, time.UTC),
		},
		{
			"between 2024-01-01 and 2024-06-30",
			"between 2024-01-01 and 2024-06-30",
			time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC),
		},
		{
			"between Q1 and Q2",
			"between Q1 2024 and Q2 2024",
			time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDateRange(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDateRange(%q) error = %v", tt.input, err)
			}

			if !result.Start.Equal(tt.wantStart) {
				t.Errorf("ParseDateRange(%q) start = %v, want %v", tt.input, result.Start, tt.wantStart)
			}
			if !result.End.Equal(tt.wantEnd) {
				t.Errorf("ParseDateRange(%q) end = %v, want %v", tt.input, result.End, tt.wantEnd)
			}
		})
	}
}

// ============================================================================
// DURATION-BASED RANGE TESTS
// ============================================================================

func TestParseRange_NextDuration(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		name      string
		input     string
		wantStart time.Time
		wantEnd   time.Time
	}{
		{
			"next 7 days",
			"next 7 days",
			time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC),
			time.Date(2024, 10, 22, 12, 0, 0, 0, time.UTC),
		},
		{
			"next 2 weeks",
			"next 2 weeks",
			time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC),
			time.Date(2024, 10, 29, 12, 0, 0, 0, time.UTC),
		},
		{
			"next 3 months",
			"next 3 months",
			time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC),
			time.Date(2025, 1, 15, 12, 0, 0, 0, time.UTC),
		},
		{
			"next 1 year",
			"next 1 year",
			time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC),
			time.Date(2025, 10, 15, 12, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDateRange(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDateRange(%q) error = %v", tt.input, err)
			}

			if !result.Start.Equal(tt.wantStart) {
				t.Errorf("ParseDateRange(%q) start = %v, want %v", tt.input, result.Start, tt.wantStart)
			}
			if !result.End.Equal(tt.wantEnd) {
				t.Errorf("ParseDateRange(%q) end = %v, want %v", tt.input, result.End, tt.wantEnd)
			}
		})
	}
}

func TestParseRange_LastDuration(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		name      string
		input     string
		wantStart time.Time
		wantEnd   time.Time
	}{
		{
			"last 7 days",
			"last 7 days",
			time.Date(2024, 10, 8, 12, 0, 0, 0, time.UTC),
			time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC),
		},
		{
			"last 2 weeks",
			"last 2 weeks",
			time.Date(2024, 10, 1, 12, 0, 0, 0, time.UTC),
			time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC),
		},
		{
			"last 3 months",
			"last 3 months",
			time.Date(2024, 7, 15, 12, 0, 0, 0, time.UTC),
			time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC),
		},
		{
			"last 1 year",
			"last 1 year",
			time.Date(2023, 10, 15, 12, 0, 0, 0, time.UTC),
			time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDateRange(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDateRange(%q) error = %v", tt.input, err)
			}

			if !result.Start.Equal(tt.wantStart) {
				t.Errorf("ParseDateRange(%q) start = %v, want %v", tt.input, result.Start, tt.wantStart)
			}
			if !result.End.Equal(tt.wantEnd) {
				t.Errorf("ParseDateRange(%q) end = %v, want %v", tt.input, result.End, tt.wantEnd)
			}
		})
	}
}

// ============================================================================
// HELPER FUNCTION TESTS
// ============================================================================

func TestGetDatesInRange(t *testing.T) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC)

	// Every day
	dates := GetDatesInRange(start, end, 1)
	if len(dates) != 5 {
		t.Errorf("GetDatesInRange() returned %d dates, want 5", len(dates))
	}

	// Every 2 days
	dates = GetDatesInRange(start, end, 2)
	if len(dates) != 3 {
		t.Errorf("GetDatesInRange(step=2) returned %d dates, want 3", len(dates))
	}
}

func TestGetBusinessDaysInRange(t *testing.T) {
	// October 14, 2024 (Monday) to October 20, 2024 (Sunday)
	start := time.Date(2024, 10, 14, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 10, 20, 0, 0, 0, 0, time.UTC)

	businessDays := GetBusinessDaysInRange(start, end)

	// Should be Mon, Tue, Wed, Thu, Fri = 5 days
	if len(businessDays) != 5 {
		t.Errorf("GetBusinessDaysInRange() returned %d days, want 5", len(businessDays))
	}

	// Verify no weekends
	for _, date := range businessDays {
		if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
			t.Errorf("GetBusinessDaysInRange() included weekend: %v", date)
		}
	}
}

func TestDaysBetween(t *testing.T) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC)

	days := DaysBetween(start, end)
	if days != 7 {
		t.Errorf("DaysBetween() = %d, want 7", days)
	}
}

func TestDurationBetween(t *testing.T) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)

	duration := DurationBetween(start, end)
	expected := 12 * time.Hour

	if duration != expected {
		t.Errorf("DurationBetween() = %v, want %v", duration, expected)
	}
}

// ============================================================================
// ERROR HANDLING TESTS
// ============================================================================

func TestParseRange_InvalidOrder(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	// Start after end should fail
	_, err := ParseDateRange("from tomorrow to yesterday", settings)
	if err == nil {
		t.Error("ParseDateRange() expected error for reversed dates, got nil")
	}
}

func TestParseRange_EmptyInput(t *testing.T) {
	_, err := ParseDateRange("", nil)
	if err == nil {
		t.Error("ParseDateRange() expected error for empty input, got nil")
	}

	var emptyErr *ErrEmptyInput
	if !errors.As(err, &emptyErr) {
		t.Errorf("ParseDateRange() error type = %T, want *ErrEmptyInput", err)
	}
}

func TestParseRange_InvalidFormat(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []string{
		"not a range",
		"random text",
		"just one date 2024-01-01",
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			_, err := ParseDateRange(input, settings)
			if err == nil {
				t.Errorf("ParseDateRange(%q) expected error, got nil", input)
			}
		})
	}
}

// ============================================================================
// CASE INSENSITIVITY TESTS
// ============================================================================

func TestParseRange_CaseInsensitive(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []string{
		"from yesterday to tomorrow",
		"FROM yesterday TO tomorrow",
		"From Yesterday To Tomorrow",
		"between yesterday and tomorrow",
		"BETWEEN yesterday AND tomorrow",
		"Between Yesterday And Tomorrow",
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			result, err := ParseDateRange(input, settings)
			if err != nil {
				t.Errorf("ParseDateRange(%q) error = %v", input, err)
			}
			if result == nil {
				t.Errorf("ParseDateRange(%q) returned nil", input)
			}
		})
	}
}

// ============================================================================
// SPLIT FUNCTION TESTS
// ============================================================================

func TestSplitRangeOnKeyword(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		name         string
		input        string
		startKeyword string
		sepKeyword   string
		wantStart    string
		wantEnd      string
		expectError  bool
	}{
		{
			"simple from...to",
			"from yesterday to tomorrow",
			"from",
			"to",
			"yesterday",
			"tomorrow",
			false,
		},
		{
			"multi-word from...to",
			"from next Monday to next Friday",
			"from",
			"to",
			"next Monday",
			"next Friday",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startStr, endStr, err := splitRangeOnKeyword(tt.input, tt.startKeyword, tt.sepKeyword, settings)

			if tt.expectError {
				if err == nil {
					t.Errorf("splitRangeOnKeyword() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("splitRangeOnKeyword() error = %v", err)
			}

			if startStr != tt.wantStart {
				t.Errorf("splitRangeOnKeyword() start = %q, want %q", startStr, tt.wantStart)
			}
			if endStr != tt.wantEnd {
				t.Errorf("splitRangeOnKeyword() end = %q, want %q", endStr, tt.wantEnd)
			}
		})
	}
}

// ============================================================================
// BENCHMARKS
// ============================================================================

func BenchmarkParseRange_FromTo(b *testing.B) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}
	for i := 0; i < b.N; i++ {
		_, _ = ParseDateRange("from yesterday to tomorrow", settings)
	}
}

func BenchmarkParseRange_BetweenAnd(b *testing.B) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}
	for i := 0; i < b.N; i++ {
		_, _ = ParseDateRange("between 2024-01-01 and 2024-12-31", settings)
	}
}

func BenchmarkParseRange_Next7Days(b *testing.B) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}
	for i := 0; i < b.N; i++ {
		_, _ = ParseDateRange("next 7 days", settings)
	}
}

func BenchmarkGetBusinessDaysInRange(b *testing.B) {
	start := time.Date(2024, 10, 14, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 10, 20, 0, 0, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		GetBusinessDaysInRange(start, end)
	}
}
