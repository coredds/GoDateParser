package translations_test

import (
	"testing"
	"time"

	"github.com/coredds/godateparser"
)

// Tests for Russian language support (ru-RU)

func TestRussian_Months(t *testing.T) {
	settings := &godateparser.Settings{Languages: []string{"ru"}}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
		wantYear  int
	}{
		// Genitive case (standard for dates)
		{"31 декабря 2024", time.December, 31, 2024},
		{"1 января 2025", time.January, 1, 2025},
		{"15 июня 2024", time.June, 15, 2024},
		{"25 дек 2024", time.December, 25, 2024},
		{"15 марта 2024", time.March, 15, 2024},
		{"10 августа 2024", time.August, 10, 2024},
		{"10 авг 2024", time.August, 10, 2024},
		// Nominative case (month names alone)
		{"март 2024", time.March, 1, 2024},
		{"декабрь 2024", time.December, 1, 2024},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("godateparser.ParseDate() error = %v", err)
			}
			if result.Month() != tt.wantMonth || result.Day() != tt.wantDay || result.Year() != tt.wantYear {
				t.Errorf("godateparser.ParseDate(%q) = %v-%v-%v, want %v-%v-%v",
					tt.input, result.Year(), result.Month(), result.Day(),
					tt.wantYear, tt.wantMonth, tt.wantDay)
			}
		})
	}
}

func TestRussian_Weekdays(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC) // Tuesday
	settings := &godateparser.Settings{
		Languages:    []string{"ru"},
		RelativeBase: base,
	}

	tests := []struct {
		input       string
		wantWeekday time.Weekday
	}{
		// Nominative case
		{"понедельник", time.Monday},
		{"вторник", time.Tuesday},
		{"среда", time.Wednesday},
		{"четверг", time.Thursday},
		{"пятница", time.Friday},
		{"суббота", time.Saturday},
		{"воскресенье", time.Sunday},
		// Prepositional case (with "в")
		{"в понедельнике", time.Monday},
		{"в вторнике", time.Tuesday},
		{"в среде", time.Wednesday},
		{"в четверге", time.Thursday},
		{"в пятнице", time.Friday},
		{"в субботе", time.Saturday},
		// Abbreviations
		{"пн", time.Monday},
		{"вт", time.Tuesday},
		{"ср", time.Wednesday},
		{"чт", time.Thursday},
		{"пт", time.Friday},
		{"сб", time.Saturday},
		{"вс", time.Sunday},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("godateparser.ParseDate() error = %v", err)
			}
			if result.Weekday() != tt.wantWeekday {
				t.Errorf("godateparser.ParseDate(%q) weekday = %v, want %v",
					tt.input, result.Weekday(), tt.wantWeekday)
			}
		})
	}
}

func TestRussian_RelativeSimple(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"ru"},
		RelativeBase: base,
	}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"вчера", 14},
		{"сегодня", 15},
		{"завтра", 16},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("godateparser.ParseDate() error = %v", err)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("godateparser.ParseDate(%q) day = %v, want %v", tt.input, result.Day(), tt.wantDay)
			}
		})
	}
}

func TestRussian_RelativeAgo(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"ru"},
		RelativeBase: base,
	}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"1 день назад", 14},
		{"2 дня назад", 13},
		{"1 неделя назад", 8},
		{"2 недели назад", 1},
		{"3 дня тому назад", 12},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("godateparser.ParseDate() error = %v", err)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("godateparser.ParseDate(%q) day = %v, want %v", tt.input, result.Day(), tt.wantDay)
			}
		})
	}
}

func TestRussian_RelativeIn(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"ru"},
		RelativeBase: base,
	}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"через 1 день", 16},
		{"через 2 дня", 17},
		{"через 1 неделю", 22},
		{"через 2 недели", 29},
		{"спустя 3 дня", 18},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("godateparser.ParseDate() error = %v", err)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("godateparser.ParseDate(%q) day = %v, want %v", tt.input, result.Day(), tt.wantDay)
			}
		})
	}
}

func TestRussian_RelativeNextLast(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"ru"},
		RelativeBase: base,
	}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
	}{
		{"следующая неделя", time.October, 22},
		{"прошлая неделя", time.October, 8},
		{"следующий месяц", time.November, 15},
		{"прошлый месяц", time.September, 15},
		{"будущая неделя", time.October, 22},
		{"предыдущая неделя", time.October, 8},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("godateparser.ParseDate() error = %v", err)
			}
			if result.Month() != tt.wantMonth || result.Day() != tt.wantDay {
				t.Errorf("godateparser.ParseDate(%q) = %v-%v, want %v-%v",
					tt.input, result.Month(), result.Day(), tt.wantMonth, tt.wantDay)
			}
		})
	}
}

func TestRussian_WeekdaysWithModifiers(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC) // Tuesday
	settings := &godateparser.Settings{
		Languages:    []string{"ru"},
		RelativeBase: base,
	}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"следующий понедельник", 21},
		{"прошлый понедельник", 14},
		{"следующая пятница", 18},
		{"прошлая пятница", 11},
		{"будущий понедельник", 21},
		{"предыдущая пятница", 11},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("godateparser.ParseDate() error = %v", err)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("godateparser.ParseDate(%q) day = %v, want %v", tt.input, result.Day(), tt.wantDay)
			}
		})
	}
}

func TestRussian_TimeExpressions(t *testing.T) {
	base := time.Date(2024, 10, 15, 0, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"ru"},
		RelativeBase: base,
	}

	tests := []struct {
		input      string
		wantHour   int
		wantMinute int
	}{
		{"полдень", 12, 0},
		{"полночь", 0, 0},
		{"3 часа дня", 15, 0},
		{"3 часа ночи", 3, 0},
		{"9 часов утра", 9, 0},
		{"7 часов вечера", 19, 0},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("godateparser.ParseDate() error = %v", err)
			}
			if result.Hour() != tt.wantHour || result.Minute() != tt.wantMinute {
				t.Errorf("godateparser.ParseDate(%q) = %v:%v, want %v:%v",
					tt.input, result.Hour(), result.Minute(), tt.wantHour, tt.wantMinute)
			}
		})
	}
}

func TestRussian_IncompleteDates(t *testing.T) {
	base := time.Date(2024, 10, 15, 14, 30, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:       []string{"ru"},
		RelativeBase:    base,
		PreferDatesFrom: "future",
	}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
		wantYear  int
	}{
		// Month only (nominative case)
		{"май", time.May, 1, 2025},          // May is before Oct, so next year
		{"декабрь", time.December, 1, 2024}, // Dec is after Oct, so this year
		{"октябрь", time.October, 1, 2024},  // Current month

		// Month and day (genitive case)
		{"15 июня", time.June, 15, 2025},        // June is before Oct, so next year
		{"25 декабря", time.December, 25, 2024}, // Dec is after Oct, so this year
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("godateparser.ParseDate() error = %v", err)
			}

			if result.Year() != tt.wantYear {
				t.Errorf("godateparser.ParseDate() year = %v, want %v", result.Year(), tt.wantYear)
			}
			if result.Month() != tt.wantMonth {
				t.Errorf("godateparser.ParseDate() month = %v, want %v", result.Month(), tt.wantMonth)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("godateparser.ParseDate() day = %v, want %v", result.Day(), tt.wantDay)
			}
		})
	}
}

func TestRussian_PeriodBoundaries(t *testing.T) {
	base := time.Date(2024, 10, 15, 14, 30, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"ru"},
		RelativeBase: base,
	}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
		wantYear  int
	}{
		// Beginning/end of periods
		{"начало месяца", time.October, 1, 2024},
		{"конец месяца", time.October, 31, 2024},
		{"начало года", time.January, 1, 2024},
		{"конец года", time.December, 31, 2024},
		{"начало недели", time.October, 14, 2024}, // Monday of current week

		// Next/last periods
		{"следующий месяц", time.November, 15, 2024},
		{"следующая неделя", time.October, 22, 2024},
		{"следующий год", time.October, 15, 2025},
		{"прошлый месяц", time.September, 15, 2024},
		{"прошлая неделя", time.October, 8, 2024},
		{"прошлый год", time.October, 15, 2023},

		// Beginning/end of next/last periods
		{"начало следующего месяца", time.November, 1, 2024},
		{"конец следующего месяца", time.November, 30, 2024},
		{"начало прошлого месяца", time.September, 1, 2024},
		{"конец прошлого года", time.December, 31, 2023},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("godateparser.ParseDate() error = %v", err)
			}

			if result.Year() != tt.wantYear {
				t.Errorf("godateparser.ParseDate() year = %v, want %v", result.Year(), tt.wantYear)
			}
			if result.Month() != tt.wantMonth {
				t.Errorf("godateparser.ParseDate() month = %v, want %v", result.Month(), tt.wantMonth)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("godateparser.ParseDate() day = %v, want %v", result.Day(), tt.wantDay)
			}
		})
	}
}

func TestRussian_ThisNextLast(t *testing.T) {
	base := time.Date(2024, 10, 15, 14, 30, 0, 0, time.UTC) // Tuesday
	settings := &godateparser.Settings{
		Languages:    []string{"ru"},
		RelativeBase: base,
	}

	tests := []struct {
		input       string
		wantMonth   time.Month
		wantDay     int
		wantWeekday time.Weekday
	}{
		// "этот/эта" (this) with weekdays
		{"этот понедельник", time.October, 21, time.Monday},
		{"эта среда", time.October, 16, time.Wednesday},
		{"эта пятница", time.October, 18, time.Friday},

		// "этот/эта" (this) with periods
		{"этот месяц", time.October, 1, 0},
		{"эта неделя", time.October, 14, 0}, // Monday (start of week)
		{"текущий месяц", time.October, 1, 0},
		{"текущая неделя", time.October, 14, 0},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("godateparser.ParseDate() error = %v", err)
			}

			if result.Month() != tt.wantMonth {
				t.Errorf("godateparser.ParseDate() month = %v, want %v", result.Month(), tt.wantMonth)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("godateparser.ParseDate() day = %v, want %v", result.Day(), tt.wantDay)
			}
			if tt.wantWeekday != 0 && result.Weekday() != tt.wantWeekday {
				t.Errorf("godateparser.ParseDate() weekday = %v, want %v", result.Weekday(), tt.wantWeekday)
			}
		})
	}
}

func TestRussian_MixedWithEnglish(t *testing.T) {
	// Test auto-detection when both Russian and English are enabled
	settings := &godateparser.Settings{Languages: []string{"ru", "en"}}

	tests := []struct {
		input    string
		wantLang string
	}{
		{"15 декабря 2024", "ru"},
		{"December 15 2024", "en"},
		{"вчера", "ru"},
		{"yesterday", "en"},
		{"понедельник", "ru"},
		{"Monday", "en"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			_, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("godateparser.ParseDate() error = %v", err)
			}
			// Just verify it parses successfully
			// Language detection is implicit
		})
	}
}

func TestRussian_PluralForms(t *testing.T) {
	// Test Russian plural forms (1, 2-4, 5+)
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"ru"},
		RelativeBase: base,
	}

	tests := []struct {
		input   string
		wantDay int
	}{
		// Singular (1)
		{"1 день назад", 14},
		{"1 неделя назад", 8},

		// Genitive singular (2-4)
		{"2 дня назад", 13},
		{"3 дня назад", 12},
		{"4 дня назад", 11},

		// Genitive plural (5+)
		{"5 дней назад", 10},
		{"10 дней назад", 5},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("godateparser.ParseDate() error = %v", err)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("godateparser.ParseDate(%q) day = %v, want %v", tt.input, result.Day(), tt.wantDay)
			}
		})
	}
}

func BenchmarkRussian_SimpleDate(b *testing.B) {
	settings := &godateparser.Settings{Languages: []string{"ru"}}
	for i := 0; i < b.N; i++ {
		_, _ = godateparser.ParseDate("15 декабря 2024", settings)
	}
}

func BenchmarkRussian_RelativeDate(b *testing.B) {
	settings := &godateparser.Settings{Languages: []string{"ru"}}
	for i := 0; i < b.N; i++ {
		_, _ = godateparser.ParseDate("2 дня назад", settings)
	}
}

func BenchmarkRussian_Weekday(b *testing.B) {
	settings := &godateparser.Settings{Languages: []string{"ru"}}
	for i := 0; i < b.N; i++ {
		_, _ = godateparser.ParseDate("следующий понедельник", settings)
	}
}
