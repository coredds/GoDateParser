package godateparser

import (
	"testing"
	"time"
)

// Tests for relative date parsing (basic and extended)

func TestParseRelative_Simple(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"yesterday", 14},
		{"today", 15},
		{"tomorrow", 16},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("ParseDate(%q) day = %v, want %v", tt.input, result.Day(), tt.wantDay)
			}
		})
	}
}

func TestParseRelative_Ago(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"1 day ago", 14},
		{"2 days ago", 13},
		{"1 week ago", 8},
		{"2 weeks ago", 1},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("ParseDate(%q) day = %v, want %v", tt.input, result.Day(), tt.wantDay)
			}
		})
	}
}

func TestParseRelative_In(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"in 1 day", 16},
		{"in 2 days", 17},
		{"in 1 week", 22},
		{"in 2 weeks", 29},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("ParseDate(%q) day = %v, want %v", tt.input, result.Day(), tt.wantDay)
			}
		})
	}
}

func TestParseRelative_LastNext(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
	}{
		{"last week", time.October, 8},
		{"next week", time.October, 22},
		{"last month", time.September, 15},
		{"next month", time.November, 15},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if result.Month() != tt.wantMonth {
				t.Errorf("ParseDate(%q) month = %v, want %v", tt.input, result.Month(), tt.wantMonth)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("ParseDate(%q) day = %v, want %v", tt.input, result.Day(), tt.wantDay)
			}
		})
	}
}

func TestParseRelative_Weekdays(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC) // Tuesday
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"next Monday", 21},
		{"last Monday", 14},
		{"next Friday", 18},
		{"last Friday", 11},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("ParseDate(%q) day = %v, want %v", tt.input, result.Day(), tt.wantDay)
			}
		})
	}
}

func TestParseRelative_AdditionalTerms(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"a fortnight ago", 1}, // 14 days ago
		{"in a fortnight", 29}, // 14 days from now
		{"a week ago", 8},
		{"an hour ago", 15}, // Same day
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("ParseDate(%q) day = %v, want %v (full: %v)", tt.input, result.Day(), tt.wantDay, result)
			}
		})
	}
}

func TestParseRelative_Decade(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		input    string
		wantYear int
	}{
		{"a decade ago", 2014},
		{"in a decade", 2034},
		{"last decade", 2014},
		{"next decade", 2034},
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

func TestParseRelative_PeriodBoundaries(t *testing.T) {
	base := time.Date(2024, 10, 15, 14, 30, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		name      string
		input     string
		wantMonth time.Month
		wantDay   int
	}{
		{"beginning of month", "beginning of month", time.October, 1},
		{"end of month", "end of month", time.October, 31},
		{"beginning of year", "beginning of year", time.January, 1},
		{"end of year", "end of year", time.December, 31},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate(%q) error = %v", tt.input, err)
			}
			if result.Month() != tt.wantMonth {
				t.Errorf("ParseDate(%q) month = %v, want %v", tt.input, result.Month(), tt.wantMonth)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("ParseDate(%q) day = %v, want %v", tt.input, result.Day(), tt.wantDay)
			}
		})
	}
}

func TestParseRelative_ComplexExpressions(t *testing.T) {
	base := time.Date(2024, 10, 15, 14, 30, 0, 0, time.UTC) // Tuesday
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		name    string
		input   string
		wantDay int
	}{
		{"a week from Tuesday", "a week from Tuesday", 29},     // Next Tuesday (Oct 22) + 7 = Oct 29
		{"2 days from Monday", "2 days from Monday", 23},       // Next Monday (Oct 21) + 2 = Oct 23
		{"3 days after tomorrow", "3 days after tomorrow", 19}, // Tomorrow (Oct 16) + 3 = Oct 19
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate(%q) error = %v", tt.input, err)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("ParseDate(%q) day = %v, want %v", tt.input, result.Day(), tt.wantDay)
			}
		})
	}
}

func TestParseRelative_Quarters(t *testing.T) {
	base := time.Date(2024, 10, 15, 14, 30, 0, 0, time.UTC) // Q4
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		name      string
		input     string
		wantMonth time.Month
		wantYear  int
	}{
		{"Q1", "Q1", time.January, 2024},
		{"Q2", "Q2", time.April, 2024},
		{"Q3", "Q3", time.July, 2024},
		{"Q4", "Q4", time.October, 2024},
		{"Q4 2025", "Q4 2025", time.October, 2025},
		{"this quarter", "this quarter", time.October, 2024},
		{"next quarter", "next quarter", time.January, 2025},
		{"last quarter", "last quarter", time.July, 2024},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate(%q) error = %v", tt.input, err)
			}
			if result.Month() != tt.wantMonth {
				t.Errorf("ParseDate(%q) month = %v, want %v", tt.input, result.Month(), tt.wantMonth)
			}
			if result.Year() != tt.wantYear {
				t.Errorf("ParseDate(%q) year = %v, want %v", tt.input, result.Year(), tt.wantYear)
			}
		})
	}
}

func TestParseRelative_ZeroQuantity(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []string{
		"0 days ago",
		"0 hours ago",
		"0 weeks ago",
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			result, err := ParseDate(input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if !result.Equal(base) {
				t.Errorf("ParseDate(%q) = %v, want %v (should be same as base)", input, result, base)
			}
		})
	}
}

func BenchmarkParseRelative_Simple(b *testing.B) {
	settings := DefaultSettings()
	for i := 0; i < b.N; i++ {
		ParseDate("yesterday", settings)
	}
}

func BenchmarkParseRelative_Ago(b *testing.B) {
	settings := DefaultSettings()
	for i := 0; i < b.N; i++ {
		ParseDate("2 days ago", settings)
	}
}

func BenchmarkParseRelative_Complex(b *testing.B) {
	settings := DefaultSettings()
	for i := 0; i < b.N; i++ {
		ParseDate("a week from Tuesday", settings)
	}
}
