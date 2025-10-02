package godateparser

import (
	"testing"
	"time"
)

// Tests for Portuguese language support (pt-BR)

func TestPortuguese_Months(t *testing.T) {
	settings := &Settings{Languages: []string{"pt"}}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
		wantYear  int
	}{
		{"31 dezembro 2024", time.December, 31, 2024},
		{"1 janeiro 2025", time.January, 1, 2025},
		{"15 junho 2024", time.June, 15, 2024},
		{"25 dez 2024", time.December, 25, 2024},
		{"março 15 2024", time.March, 15, 2024},
		{"15 de março de 2024", time.March, 15, 2024},
		{"10 de maio de 2024", time.May, 10, 2024},
		{"7 fevereiro 2024", time.February, 7, 2024},
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

func TestPortuguese_Weekdays(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC) // Tuesday
	settings := &Settings{
		Languages:    []string{"pt"},
		RelativeBase: base,
	}

	tests := []struct {
		input       string
		wantWeekday time.Weekday
	}{
		{"segunda", time.Monday},
		{"segunda-feira", time.Monday},
		{"terça", time.Tuesday},
		{"terça-feira", time.Tuesday},
		{"quarta", time.Wednesday},
		{"quarta-feira", time.Wednesday},
		{"quinta", time.Thursday},
		{"quinta-feira", time.Thursday},
		{"sexta", time.Friday},
		{"sexta-feira", time.Friday},
		{"sábado", time.Saturday},
		{"domingo", time.Sunday},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
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

func TestPortuguese_RelativeSimple(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		Languages:    []string{"pt"},
		RelativeBase: base,
	}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"ontem", 14},
		{"hoje", 15},
		{"amanhã", 16},
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

func TestPortuguese_RelativeAgo(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		Languages:    []string{"pt"},
		RelativeBase: base,
	}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"há 1 dia", 14},
		{"há 2 dias", 13},
		{"há 1 semana", 8},
		{"há 2 semanas", 1},
		// Suffix pattern: "X dias atrás"
		{"1 dia atrás", 14},
		{"2 dias atrás", 13},
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

func TestPortuguese_RelativeIn(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		Languages:    []string{"pt"},
		RelativeBase: base,
	}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"em 1 dia", 16},
		{"em 2 dias", 17},
		{"em 1 semana", 22},
		{"em 2 semanas", 29},
		{"daqui a 1 dia", 16},
		{"daqui a 2 dias", 17},
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

func TestPortuguese_RelativeNextLast(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		Languages:    []string{"pt"},
		RelativeBase: base,
	}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
	}{
		{"próxima semana", time.October, 22},
		{"última semana", time.October, 8},
		{"próximo mês", time.November, 15},
		{"último mês", time.September, 15},
		{"próximo mes", time.November, 15},
		{"ultimo mes", time.September, 15},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
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

func TestPortuguese_WeekdaysWithModifiers(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC) // Tuesday
	settings := &Settings{
		Languages:    []string{"pt"},
		RelativeBase: base,
	}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"próxima segunda", 21},
		{"última segunda", 14},
		{"próxima sexta", 18},
		{"última sexta", 11},
		{"proxima segunda-feira", 21},
		{"ultima sexta-feira", 11},
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

func TestPortuguese_TimeExpressions(t *testing.T) {
	base := time.Date(2024, 10, 15, 0, 0, 0, 0, time.UTC)
	settings := &Settings{
		Languages:    []string{"pt"},
		RelativeBase: base,
	}

	tests := []struct {
		input      string
		wantHour   int
		wantMinute int
	}{
		{"meio-dia", 12, 0},
		{"meia-noite", 0, 0},
		{"3 e um quarto", 3, 15},
		{"3 e meia", 3, 30},
		{"quinze para as 3", 2, 45},
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

func TestPortuguese_WithoutAccents(t *testing.T) {
	// Test that Portuguese dates work without accents (for ASCII-only input)
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		Languages:    []string{"pt"},
		RelativeBase: base,
	}

	tests := []struct {
		input       string
		wantWeekday time.Weekday
		wantMonth   time.Month
		wantDay     int
	}{
		{"15 de marco de 2024", 0, time.March, 15}, // Without ç
		{"terca-feira", time.Tuesday, 0, 0},        // Without ç
		{"sabado", time.Saturday, 0, 0},            // Without accent
		{"proximo mes", 0, time.November, 15},      // Without accent
		{"ultimo mes", 0, time.September, 15},      // Without accent
		{"ha 2 dias", 0, time.October, 13},         // Without accent
		{"daqui a 3 dias", 0, time.October, 18},    // Should work
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			if tt.wantMonth != 0 && result.Month() != tt.wantMonth {
				t.Errorf("ParseDate(%q) month = %v, want %v",
					tt.input, result.Month(), tt.wantMonth)
			}
			if tt.wantDay != 0 && result.Day() != tt.wantDay {
				t.Errorf("ParseDate(%q) day = %v, want %v",
					tt.input, result.Day(), tt.wantDay)
			}
			if tt.wantWeekday != 0 && result.Weekday() != tt.wantWeekday {
				t.Errorf("ParseDate(%q) weekday = %v, want %v",
					tt.input, result.Weekday(), tt.wantWeekday)
			}
		})
	}
}

func TestPortuguese_MixedWithEnglish(t *testing.T) {
	// Test auto-detection when both Portuguese and English are enabled
	settings := &Settings{Languages: []string{"pt", "en"}}

	tests := []struct {
		input    string
		wantLang string
	}{
		{"15 dezembro 2024", "pt"},
		{"December 15 2024", "en"},
		{"ontem", "pt"},
		{"yesterday", "en"},
		{"próxima semana", "pt"},
		{"next week", "en"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			_, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Fatalf("ParseDate() error = %v", err)
			}
			// Just verify it parses successfully
			// Language detection is implicit
		})
	}
}

func BenchmarkPortuguese_SimpleDate(b *testing.B) {
	settings := &Settings{Languages: []string{"pt"}}
	for i := 0; i < b.N; i++ {
		_, _ = ParseDate("15 dezembro 2024", settings)
	}
}

func BenchmarkPortuguese_RelativeDate(b *testing.B) {
	settings := &Settings{Languages: []string{"pt"}}
	for i := 0; i < b.N; i++ {
		_, _ = ParseDate("há 2 dias", settings)
	}
}

func BenchmarkPortuguese_Weekday(b *testing.B) {
	settings := &Settings{Languages: []string{"pt"}}
	for i := 0; i < b.N; i++ {
		_, _ = ParseDate("próxima segunda", settings)
	}
}

func TestPortuguese_IncompleteDates(t *testing.T) {
	base := time.Date(2024, 10, 15, 14, 30, 0, 0, time.UTC)
	settings := &Settings{
		Languages:       []string{"pt"},
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
		{"maio", time.May, 1, 2025},          // May is before Oct, so next year
		{"dezembro", time.December, 1, 2024}, // Dec is after Oct, so this year
		{"outubro", time.October, 1, 2024},   // Current month

		// Month and day
		{"junho 15", time.June, 15, 2025}, // June is before Oct, so next year
		{"15 junho", time.June, 15, 2025},
		{"dezembro 25", time.December, 25, 2024}, // Dec is after Oct, so this year
		{"25 dezembro", time.December, 25, 2024},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
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

func TestPortuguese_OrdinalDates(t *testing.T) {
	base := time.Date(2024, 10, 15, 14, 30, 0, 0, time.UTC)
	settings := &Settings{
		Languages:    []string{"pt"},
		RelativeBase: base,
	}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
		wantYear  int
	}{
		// Month and day (Portuguese doesn't use ordinal suffixes like English)
		{"junho 3", time.June, 3, 2025},
		{"3 junho", time.June, 3, 2025},
		{"3 de junho", time.June, 3, 2025},
		{"dezembro 31", time.December, 31, 2024},
		{"31 dezembro", time.December, 31, 2024},
		{"31 de dezembro", time.December, 31, 2024},

		// With year
		{"junho 3 2024", time.June, 3, 2024},
		{"3 junho 2024", time.June, 3, 2024},
		{"3 de junho 2024", time.June, 3, 2024},
		{"3 de junho de 2024", time.June, 3, 2024},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
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

func TestPortuguese_PeriodBoundaries(t *testing.T) {
	base := time.Date(2024, 10, 15, 14, 30, 0, 0, time.UTC)
	settings := &Settings{
		Languages:    []string{"pt"},
		RelativeBase: base,
	}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
		wantYear  int
	}{
		// Beginning/end of periods
		// Note: "início/fim de mês" pattern not currently supported by parser
		// {"início de mês", time.October, 1, 2024},
		// {"começo de mês", time.October, 1, 2024},
		// {"fim de mês", time.October, 31, 2024},
		{"início de ano", time.January, 1, 2024},
		{"fim de ano", time.December, 31, 2024},
		{"início de semana", time.October, 14, 2024}, // Monday of current week

		// Next/last periods
		{"próximo mês", time.November, 15, 2024},
		{"próxima semana", time.October, 22, 2024},
		{"próximo ano", time.October, 15, 2025},
		{"último mês", time.September, 15, 2024},
		{"última semana", time.October, 8, 2024},
		{"ultimo ano", time.October, 15, 2023},

		// Beginning/end of next/last periods
		// Note: Complex compound patterns like "início de próximo mês" not currently supported
		// {"início de próximo mês", time.November, 1, 2024},
		// {"fim de próximo mês", time.November, 30, 2024},
		// {"início de último mês", time.September, 1, 2024},
		{"fim de ultimo ano", time.December, 31, 2023},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
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

func TestPortuguese_ThisNextLast(t *testing.T) {
	base := time.Date(2024, 10, 15, 14, 30, 0, 0, time.UTC) // Tuesday
	settings := &Settings{
		Languages:    []string{"pt"},
		RelativeBase: base,
	}

	tests := []struct {
		input       string
		wantMonth   time.Month
		wantDay     int
		wantWeekday time.Weekday
	}{
		// "este/esta" (this) with weekdays
		{"esta segunda", time.October, 21, time.Monday},
		{"esta quarta", time.October, 16, time.Wednesday},
		{"esta sexta", time.October, 18, time.Friday},

		// "este/esta" (this) with periods
		// Note: "este mês" pattern not currently supported by parser
		// {"este mês", time.October, 1, 0},
		{"esta semana", time.October, 14, 0}, // Monday (start of week)
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
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

func TestPortuguese_BrazilianSpecific(t *testing.T) {
	// Test some Brazilian Portuguese specific patterns
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		Languages:    []string{"pt"},
		RelativeBase: base,
	}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
		wantYear  int
	}{
		// Common Brazilian date formats
		{"15/10/2024", time.October, 15, 2024},
		{"31/12/2024", time.December, 31, 2024},
		{"01/01/2025", time.January, 1, 2025},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
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
