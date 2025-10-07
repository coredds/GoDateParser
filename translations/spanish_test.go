package translations_test

import (
	"testing"
	"time"

	"github.com/coredds/godateparser"
)

// Tests for Spanish language support (es-ES)

func TestSpanish_Months(t *testing.T) {
	settings := &godateparser.Settings{Languages: []string{"es"}}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
		wantYear  int
	}{
		{"31 diciembre 2024", time.December, 31, 2024},
		{"1 enero 2025", time.January, 1, 2025},
		{"15 junio 2024", time.June, 15, 2024},
		{"25 dic 2024", time.December, 25, 2024},
		{"marzo 15 2024", time.March, 15, 2024},
		{"15 de marzo de 2024", time.March, 15, 2024},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
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

func TestSpanish_Weekdays(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC) // Tuesday
	settings := &godateparser.Settings{
		Languages:    []string{"es"},
		RelativeBase: base,
	}

	tests := []struct {
		input       string
		wantWeekday time.Weekday
	}{
		{"lunes", time.Monday},
		{"martes", time.Tuesday},
		{"miércoles", time.Wednesday},
		{"jueves", time.Thursday},
		{"viernes", time.Friday},
		{"sábado", time.Saturday},
		{"domingo", time.Sunday},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if result.Weekday() != tt.wantWeekday {
				t.Errorf("ParseDate(%q) weekday = %v, want %v",
					tt.input, result.Weekday(), tt.wantWeekday)
			}
		})
	}
}

func TestSpanish_RelativeSimple(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"es"},
		RelativeBase: base,
	}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"ayer", 14},
		{"hoy", 15},
		{"mañana", 16},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("ParseDate(%q) day = %v, want %v", tt.input, result.Day(), tt.wantDay)
			}
		})
	}
}

func TestSpanish_RelativeAgo(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"es"},
		RelativeBase: base,
	}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"hace 1 día", 14},
		{"hace 2 días", 13},
		{"hace 1 semana", 8},
		{"hace 2 semanas", 1},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("ParseDate(%q) day = %v, want %v", tt.input, result.Day(), tt.wantDay)
			}
		})
	}
}

func TestSpanish_RelativeIn(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"es"},
		RelativeBase: base,
	}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"en 1 día", 16},
		{"en 2 días", 17},
		{"en 1 semana", 22},
		{"en 2 semanas", 29},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("ParseDate(%q) day = %v, want %v", tt.input, result.Day(), tt.wantDay)
			}
		})
	}
}

func TestSpanish_RelativeNextLast(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"es"},
		RelativeBase: base,
	}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
	}{
		{"próxima semana", time.October, 22},
		{"última semana", time.October, 8},
		{"próximo mes", time.November, 15},
		{"último mes", time.September, 15},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if result.Month() != tt.wantMonth || result.Day() != tt.wantDay {
				t.Errorf("ParseDate(%q) = %v-%v, want %v-%v",
					tt.input, result.Month(), result.Day(), tt.wantMonth, tt.wantDay)
			}
		})
	}
}

func TestSpanish_WeekdaysWithModifiers(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC) // Tuesday
	settings := &godateparser.Settings{
		Languages:    []string{"es"},
		RelativeBase: base,
	}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"próximo lunes", 21},
		{"último lunes", 14},
		{"próximo viernes", 18},
		{"último viernes", 11},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("ParseDate(%q) day = %v, want %v", tt.input, result.Day(), tt.wantDay)
			}
		})
	}
}

func TestSpanish_TimeExpressions(t *testing.T) {
	base := time.Date(2024, 10, 15, 0, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"es"},
		RelativeBase: base,
	}

	tests := []struct {
		input      string
		wantHour   int
		wantMinute int
	}{
		{"mediodía", 12, 0},
		{"medianoche", 0, 0},
		{"3 y cuarto", 3, 15},
		{"3 y media", 3, 30},
		{"menos cuarto las 3", 2, 45},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
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

func TestSpanish_WithoutAccents(t *testing.T) {
	// Test that Spanish dates work without accents (for ASCII-only input)
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"es"},
		RelativeBase: base,
	}

	tests := []struct {
		input       string
		wantWeekday time.Weekday
		wantMonth   time.Month
	}{
		{"15 de marzo de 2024", 0, time.March},
		{"miercoles", time.Wednesday, 0}, // Without accent
		{"sabado", time.Saturday, 0},     // Without accent
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if tt.wantMonth != 0 && result.Month() != tt.wantMonth {
				t.Errorf("ParseDate(%q) month = %v, want %v",
					tt.input, result.Month(), tt.wantMonth)
			}
			if tt.wantWeekday != 0 && result.Weekday() != tt.wantWeekday {
				t.Errorf("ParseDate(%q) weekday = %v, want %v",
					tt.input, result.Weekday(), tt.wantWeekday)
			}
		})
	}
}

func TestSpanish_MixedWithEnglish(t *testing.T) {
	// Test auto-detection when both Spanish and English are enabled
	settings := &godateparser.Settings{Languages: []string{"es", "en"}}

	tests := []struct {
		input    string
		wantLang string
	}{
		{"15 diciembre 2024", "es"},
		{"December 15 2024", "en"},
		{"ayer", "es"},
		{"yesterday", "en"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			_, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			// Just verify it parses successfully
			// Language detection is implicit
		})
	}
}

func BenchmarkSpanish_SimpleDate(b *testing.B) {
	settings := &godateparser.Settings{Languages: []string{"es"}}
	for i := 0; i < b.N; i++ {
		_, _ = godateparser.ParseDate("15 diciembre 2024", settings)
	}
}

func BenchmarkSpanish_RelativeDate(b *testing.B) {
	settings := &godateparser.Settings{Languages: []string{"es"}}
	for i := 0; i < b.N; i++ {
		_, _ = godateparser.ParseDate("hace 2 días", settings)
	}
}

func BenchmarkSpanish_Weekday(b *testing.B) {
	settings := &godateparser.Settings{Languages: []string{"es"}}
	for i := 0; i < b.N; i++ {
		_, _ = godateparser.ParseDate("próximo lunes", settings)
	}
}

func TestSpanish_IncompleteDates(t *testing.T) {
	base := time.Date(2024, 10, 15, 14, 30, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:       []string{"es"},
		RelativeBase:    base,
		PreferDatesFrom: "future",
	}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
		wantYear  int
	}{
		// Month only
		{"mayo", time.May, 1, 2025},           // May is before Oct, so next year
		{"diciembre", time.December, 1, 2024}, // Dec is after Oct, so this year
		{"octubre", time.October, 1, 2024},    // Current month

		// Month and day
		{"junio 15", time.June, 15, 2025}, // June is before Oct, so next year
		{"15 junio", time.June, 15, 2025},
		{"diciembre 25", time.December, 25, 2024}, // Dec is after Oct, so this year
		{"25 diciembre", time.December, 25, 2024},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}

			if result.Year() != tt.wantYear {
				t.Errorf("ParseDate() year = %v, want %v", result.Year(), tt.wantYear)
			}
			if result.Month() != tt.wantMonth {
				t.Errorf("ParseDate() month = %v, want %v", result.Month(), tt.wantMonth)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("ParseDate() day = %v, want %v", result.Day(), tt.wantDay)
			}
		})
	}
}

func TestSpanish_OrdinalDates(t *testing.T) {
	base := time.Date(2024, 10, 15, 14, 30, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"es"},
		RelativeBase: base,
	}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
		wantYear  int
	}{
		// Month and day (Spanish doesn't use ordinal suffixes like English)
		{"junio 3", time.June, 3, 2025},
		{"3 junio", time.June, 3, 2025},
		{"3 de junio", time.June, 3, 2025},
		{"diciembre 31", time.December, 31, 2024},
		{"31 diciembre", time.December, 31, 2024},
		{"31 de diciembre", time.December, 31, 2024},

		// With year
		{"junio 3 2024", time.June, 3, 2024},
		{"3 junio 2024", time.June, 3, 2024},
		{"3 de junio 2024", time.June, 3, 2024},
		{"3 de junio de 2024", time.June, 3, 2024},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}

			if result.Year() != tt.wantYear {
				t.Errorf("ParseDate() year = %v, want %v", result.Year(), tt.wantYear)
			}
			if result.Month() != tt.wantMonth {
				t.Errorf("ParseDate() month = %v, want %v", result.Month(), tt.wantMonth)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("ParseDate() day = %v, want %v", result.Day(), tt.wantDay)
			}
		})
	}
}

func TestSpanish_PeriodBoundaries(t *testing.T) {
	base := time.Date(2024, 10, 15, 14, 30, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"es"},
		RelativeBase: base,
	}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
		wantYear  int
	}{
		// Beginning/end of periods
		{"inicio de mes", time.October, 1, 2024},
		{"comienzo de mes", time.October, 1, 2024},
		{"fin de mes", time.October, 31, 2024},
		{"inicio de año", time.January, 1, 2024},
		{"fin de año", time.December, 31, 2024},
		{"inicio de semana", time.October, 14, 2024}, // Monday of current week

		// Next/last periods
		{"próximo mes", time.November, 15, 2024},
		{"próxima semana", time.October, 22, 2024},
		{"próximo año", time.October, 15, 2025},
		{"último mes", time.September, 15, 2024},
		{"última semana", time.October, 8, 2024},
		{"ultimo ano", time.October, 15, 2023},

		// Beginning/end of next/last periods
		{"inicio de próximo mes", time.November, 1, 2024},
		{"fin de próximo mes", time.November, 30, 2024},
		{"inicio de último mes", time.September, 1, 2024},
		{"fin de ultimo ano", time.December, 31, 2023},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}

			if result.Year() != tt.wantYear {
				t.Errorf("ParseDate() year = %v, want %v", result.Year(), tt.wantYear)
			}
			if result.Month() != tt.wantMonth {
				t.Errorf("ParseDate() month = %v, want %v", result.Month(), tt.wantMonth)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("ParseDate() day = %v, want %v", result.Day(), tt.wantDay)
			}
		})
	}
}

func TestSpanish_ThisNextLast(t *testing.T) {
	base := time.Date(2024, 10, 15, 14, 30, 0, 0, time.UTC) // Tuesday
	settings := &godateparser.Settings{
		Languages:    []string{"es"},
		RelativeBase: base,
	}

	tests := []struct {
		input       string
		wantMonth   time.Month
		wantDay     int
		wantWeekday time.Weekday
	}{
		// "este" (this) with weekdays
		{"este lunes", time.October, 21, time.Monday},
		{"este miércoles", time.October, 16, time.Wednesday},
		{"este viernes", time.October, 18, time.Friday},

		// "este" (this) with periods
		{"este mes", time.October, 1, 0},
		{"esta semana", time.October, 14, 0}, // Monday (start of week)
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}

			if result.Month() != tt.wantMonth {
				t.Errorf("ParseDate() month = %v, want %v", result.Month(), tt.wantMonth)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("ParseDate() day = %v, want %v", result.Day(), tt.wantDay)
			}
			if tt.wantWeekday != 0 && result.Weekday() != tt.wantWeekday {
				t.Errorf("ParseDate() weekday = %v, want %v", result.Weekday(), tt.wantWeekday)
			}
		})
	}
}
