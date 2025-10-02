package godateparser

import (
	"testing"
	"time"
)

// Tests for absolute date parsing (ISO 8601, numeric formats, month names)

func TestParseAbsolute_ISO8601(t *testing.T) {
	tests := []struct {
		input string
		want  time.Time
	}{
		{"2024-12-31", time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)},
		{"2024-01-01", time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"2024-06-15", time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC)},
		{"1999-12-31", time.Date(1999, 12, 31, 0, 0, 0, 0, time.UTC)},
		{"2000-01-01", time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)},
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

func TestParseAbsolute_ISO8601WithTime(t *testing.T) {
	tests := []struct {
		input string
		want  time.Time
	}{
		{"2024-12-31T10:30:00", time.Date(2024, 12, 31, 10, 30, 0, 0, time.UTC)},
		{"2024-01-01T00:00:00", time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"2024-06-15T23:59:59", time.Date(2024, 6, 15, 23, 59, 59, 0, time.UTC)},
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

func TestParseAbsolute_TwoDigitYear(t *testing.T) {
	tests := []struct {
		input string
		want  time.Time
	}{
		// Years 00-69 -> 2000-2069
		{"24-12-31", time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)},
		{"00-01-01", time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"69-12-31", time.Date(2069, 12, 31, 0, 0, 0, 0, time.UTC)},
		// Years 70-99 -> 1970-1999
		{"70-01-01", time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"99-12-31", time.Date(1999, 12, 31, 0, 0, 0, 0, time.UTC)},
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

func TestParseAbsolute_NumericMDY(t *testing.T) {
	settings := &Settings{DateOrder: "MDY"}

	tests := []struct {
		input string
		want  time.Time
	}{
		{"12/31/2024", time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)},
		{"01/01/2024", time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"06/15/2024", time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC)},
		{"12-31-2024", time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if !result.Equal(tt.want) {
				t.Errorf("ParseDate() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestParseAbsolute_NumericDMY(t *testing.T) {
	settings := &Settings{DateOrder: "DMY"}

	tests := []struct {
		input string
		want  time.Time
	}{
		{"31/12/2024", time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)},
		{"01/01/2024", time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"15/06/2024", time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC)},
		{"31-12-2024", time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if !result.Equal(tt.want) {
				t.Errorf("ParseDate() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestParseAbsolute_MonthNames(t *testing.T) {
	tests := []struct {
		input string
		want  time.Time
	}{
		{"December 31, 2024", time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)},
		{"31 Dec 2024", time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)},
		{"Dec 31 2024", time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)},
		{"January 1, 2024", time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"1 Jan 2024", time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
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

func TestParseAbsolute_AutoDetectDateOrder(t *testing.T) {
	// When DateOrder is explicitly unset (empty string), should auto-detect from input
	tests := []struct {
		name  string
		input string
		want  time.Time
	}{
		{"DMY detected (day > 12)", "25/06/2024", time.Date(2024, 6, 25, 0, 0, 0, 0, time.UTC)},
		{"MDY detected (month > 12 impossible)", "06/25/2024", time.Date(2024, 6, 25, 0, 0, 0, 0, time.UTC)},
		{"DMY detected", "13/01/2024", time.Date(2024, 1, 13, 0, 0, 0, 0, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings := &Settings{DateOrder: ""} // Explicitly empty
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if !result.Equal(tt.want) {
				t.Errorf("ParseDate() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestParseAbsolute_InvalidDates(t *testing.T) {
	tests := []string{
		"2024-13-01", // Invalid month
		"2024-00-01", // Invalid month
		"2024-01-32", // Invalid day
		"2024-02-30", // Feb doesn't have 30 days
		"2023-02-29", // Not a leap year
		"2024-04-31", // April has 30 days
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			_, err := ParseDate(input, nil)
			if err == nil {
				t.Errorf("ParseDate(%q) should return error for invalid date", input)
			}
		})
	}
}

func TestParseAbsolute_LeapYear(t *testing.T) {
	tests := []struct {
		input       string
		expectError bool
	}{
		{"2024-02-29", false}, // 2024 is a leap year
		{"2023-02-29", true},  // 2023 is not a leap year
		{"2000-02-29", false}, // 2000 is a leap year
		{"1900-02-29", true},  // 1900 is not a leap year
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

func BenchmarkParseAbsolute_ISO8601(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ParseDate("2024-12-31", nil)
	}
}

func BenchmarkParseAbsolute_Numeric(b *testing.B) {
	settings := &Settings{DateOrder: "MDY"}
	for i := 0; i < b.N; i++ {
		_, _ = ParseDate("12/31/2024", settings)
	}
}

func BenchmarkParseAbsolute_MonthName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ParseDate("December 31, 2024", nil)
	}
}
