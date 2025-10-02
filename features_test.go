package godateparser

import (
	"testing"
	"time"
)

// Tests for advanced features (incomplete dates, ordinals, week numbers, natural time, PREFER_DATES_FROM)

// PREFER_DATES_FROM Tests

func TestPreferDatesFrom_Weekdays(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC) // Tuesday

	tests := []struct {
		name            string
		input           string
		preferDatesFrom string
		wantDay         int
	}{
		{"Monday future", "Monday", "future", 21},
		{"Monday past", "Monday", "past", 14},
		{"Friday future", "Friday", "future", 18},
		{"Friday past", "Friday", "past", 11},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings := &Settings{
				RelativeBase:    base,
				PreferDatesFrom: tt.preferDatesFrom,
			}
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

// Incomplete Date Tests

func TestIncompleteDate_YearOnly(t *testing.T) {
	tests := []struct {
		input string
		want  time.Time
	}{
		{"2024", time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"2000", time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"1999", time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, nil)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if !result.Equal(tt.want) {
				t.Errorf("ParseDate(%q) = %v, want %v", tt.input, result, tt.want)
			}
		})
	}
}

func TestIncompleteDate_MonthOnly(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantYear  int
	}{
		{"January", time.January, 2025},
		{"May", time.May, 2025},
		{"October", time.October, 2024},
		{"December", time.December, 2024},
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
			if result.Year() != tt.wantYear {
				t.Errorf("ParseDate(%q) year = %v, want %v", tt.input, result.Year(), tt.wantYear)
			}
		})
	}
}

func TestIncompleteDate_MonthAndDay(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
		wantYear  int
	}{
		{"June 15", time.June, 15, 2025},
		{"October 20", time.October, 20, 2024},
		{"December 25", time.December, 25, 2024},
		{"15 June", time.June, 15, 2025},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if result.Month() != tt.wantMonth || result.Day() != tt.wantDay || result.Year() != tt.wantYear {
				t.Errorf("ParseDate(%q) = %v-%v-%v, want %v-%v-%v",
					tt.input, result.Year(), result.Month(), result.Day(),
					tt.wantYear, tt.wantMonth, tt.wantDay)
			}
		})
	}
}

// Ordinal Date Tests

func TestOrdinalDate_Basic(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"1st", 1},
		{"15th", 15},
		{"20th", 20},
		{"31st", 31},
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

func TestOrdinalDate_WithMonth(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
	}{
		{"June 3rd", time.June, 3},
		{"3rd June", time.June, 3},
		{"3rd of June", time.June, 3},
		{"December 25th", time.December, 25},
		{"21st March", time.March, 21},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
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

func TestOrdinalDate_FullDate(t *testing.T) {
	tests := []struct {
		input string
		want  time.Time
	}{
		{"June 3rd 2024", time.Date(2024, 6, 3, 0, 0, 0, 0, time.UTC)},
		{"3rd of June 2024", time.Date(2024, 6, 3, 0, 0, 0, 0, time.UTC)},
		{"December 25th 2023", time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC)},
		{"1st January 2025", time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, nil)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if !result.Equal(tt.want) {
				t.Errorf("ParseDate(%q) = %v, want %v", tt.input, result, tt.want)
			}
		})
	}
}

// Week Number Tests

func TestWeekNumber_ISO8601(t *testing.T) {
	tests := []struct {
		input     string
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{"2024-W01", 2024, 1, 1},
		{"2024W01", 2024, 1, 1},
		{"2024-W15", 2024, 4, 8},
		{"2024-W52", 2024, 12, 23},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, nil)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if result.Year() != tt.wantYear || result.Month() != tt.wantMonth || result.Day() != tt.wantDay {
				t.Errorf("ParseDate(%q) = %v-%v-%v, want %v-%v-%v",
					tt.input, result.Year(), result.Month(), result.Day(),
					tt.wantYear, tt.wantMonth, tt.wantDay)
			}
		})
	}
}

func TestWeekNumber_WithWeekday(t *testing.T) {
	tests := []struct {
		input       string
		wantWeekday time.Weekday
	}{
		{"2024-W15-1", time.Monday},
		{"2024-W15-3", time.Wednesday},
		{"2024-W15-5", time.Friday},
		{"2024-W15-7", time.Sunday},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, nil)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if result.Weekday() != tt.wantWeekday {
				t.Errorf("ParseDate(%q) weekday = %v, want %v", tt.input, result.Weekday(), tt.wantWeekday)
			}
		})
	}
}

func TestWeekNumber_NaturalLanguage(t *testing.T) {
	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
	}{
		{"Week 15 2024", time.April, 8},
		{"2024 Week 15", time.April, 8},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, nil)
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

// Natural Time Expression Tests

func TestNaturalTime_QuarterPast(t *testing.T) {
	base := time.Date(2024, 10, 15, 0, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		input      string
		wantHour   int
		wantMinute int
	}{
		{"quarter past 3", 3, 15},
		{"quarter past 9", 9, 15},
		{"quarter past noon", 12, 15},
		{"quarter after 3", 3, 15},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if result.Hour() != tt.wantHour || result.Minute() != tt.wantMinute {
				t.Errorf("ParseDate(%q) = %v:%v, want %v:%v",
					tt.input, result.Hour(), result.Minute(), tt.wantHour, tt.wantMinute)
			}
		})
	}
}

func TestNaturalTime_HalfPast(t *testing.T) {
	base := time.Date(2024, 10, 15, 0, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		input      string
		wantHour   int
		wantMinute int
	}{
		{"half past 3", 3, 30},
		{"half past 9", 9, 30},
		{"half past midnight", 0, 30},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if result.Hour() != tt.wantHour || result.Minute() != tt.wantMinute {
				t.Errorf("ParseDate(%q) = %v:%v, want %v:%v",
					tt.input, result.Hour(), result.Minute(), tt.wantHour, tt.wantMinute)
			}
		})
	}
}

func TestNaturalTime_QuarterTo(t *testing.T) {
	base := time.Date(2024, 10, 15, 0, 0, 0, 0, time.UTC)
	settings := &Settings{RelativeBase: base}

	tests := []struct {
		input      string
		wantHour   int
		wantMinute int
	}{
		{"quarter to 3", 2, 45},
		{"quarter to 12", 11, 45},
		{"quarter to midnight", 23, 45},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if result.Hour() != tt.wantHour || result.Minute() != tt.wantMinute {
				t.Errorf("ParseDate(%q) = %v:%v, want %v:%v",
					tt.input, result.Hour(), result.Minute(), tt.wantHour, tt.wantMinute)
			}
		})
	}
}

func BenchmarkFeatures_IncompleteDate(b *testing.B) {
	settings := DefaultSettings()
	for i := 0; i < b.N; i++ {
		_, _ = ParseDate("June 15", settings)
	}
}

func BenchmarkFeatures_OrdinalDate(b *testing.B) {
	settings := DefaultSettings()
	for i := 0; i < b.N; i++ {
		_, _ = ParseDate("3rd June", settings)
	}
}

func BenchmarkFeatures_WeekNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ParseDate("2024-W15", nil)
	}
}

func BenchmarkFeatures_NaturalTime(b *testing.B) {
	settings := DefaultSettings()
	for i := 0; i < b.N; i++ {
		_, _ = ParseDate("quarter past 3", settings)
	}
}
