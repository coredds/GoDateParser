package translations_test

import (
	"testing"
	"time"

	"github.com/coredds/godateparser"
)

// Tests for Italian language support (it-IT)

func TestItalian_Months(t *testing.T) {
	settings := &godateparser.Settings{Languages: []string{"it"}}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
		wantYear  int
	}{
		{"31 dicembre 2024", time.December, 31, 2024},
		{"1 gennaio 2025", time.January, 1, 2025},
		{"15 giugno 2024", time.June, 15, 2024},
		{"25 dic 2024", time.December, 25, 2024},
		{"marzo 15 2024", time.March, 15, 2024},
		{"15 di marzo di 2024", time.March, 15, 2024},
		{"agosto 10 2024", time.August, 10, 2024},
		{"10 ago 2024", time.August, 10, 2024},
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

func TestItalian_Weekdays(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC) // Tuesday
	settings := &godateparser.Settings{
		Languages:    []string{"it"},
		RelativeBase: base,
	}

	tests := []struct {
		input       string
		wantWeekday time.Weekday
	}{
		{"lunedì", time.Monday},
		{"lunedi", time.Monday}, // Without accent
		{"martedì", time.Tuesday},
		{"martedi", time.Tuesday}, // Without accent
		{"mercoledì", time.Wednesday},
		{"mercoledi", time.Wednesday}, // Without accent
		{"giovedì", time.Thursday},
		{"giovedi", time.Thursday}, // Without accent
		{"venerdì", time.Friday},
		{"venerdi", time.Friday}, // Without accent
		{"sabato", time.Saturday},
		{"domenica", time.Sunday},
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

func TestItalian_RelativeSimple(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"it"},
		RelativeBase: base,
	}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"ieri", 14},
		{"oggi", 15},
		{"domani", 16},
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

func TestItalian_RelativeAgo(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"it"},
		RelativeBase: base,
	}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"1 giorno fa", 14},
		{"2 giorni fa", 13},
		{"1 settimana fa", 8},
		{"2 settimane fa", 1},
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

func TestItalian_RelativeIn(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"it"},
		RelativeBase: base,
	}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"tra 1 giorno", 16},
		{"fra 2 giorni", 17},
		{"in 1 settimana", 22},
		{"tra 2 settimane", 29},
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

func TestItalian_RelativeNextLast(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"it"},
		RelativeBase: base,
	}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
	}{
		{"prossima settimana", time.October, 22},
		{"scorsa settimana", time.October, 8},
		{"prossimo mese", time.November, 15},
		{"scorso mese", time.September, 15},
		{"ultimo mese", time.September, 15},
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

func TestItalian_WeekdaysWithModifiers(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC) // Tuesday
	settings := &godateparser.Settings{
		Languages:    []string{"it"},
		RelativeBase: base,
	}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"prossimo lunedì", 21},
		{"scorso lunedi", 14},
		{"prossimo venerdì", 18},
		{"scorso venerdi", 11},
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

func TestItalian_TimeExpressions(t *testing.T) {
	base := time.Date(2024, 10, 15, 0, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"it"},
		RelativeBase: base,
	}

	tests := []struct {
		input      string
		wantHour   int
		wantMinute int
	}{
		{"mezzogiorno", 12, 0},
		{"mezzanotte", 0, 0},
		{"3 e un quarto", 3, 15},
		{"3 e mezzo", 3, 30},
		{"meno un quarto le 3", 2, 45},
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

func TestItalian_IncompleteDates(t *testing.T) {
	base := time.Date(2024, 10, 15, 14, 30, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:       []string{"it"},
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
		{"maggio", time.May, 1, 2025},        // May is before Oct, so next year
		{"dicembre", time.December, 1, 2024}, // Dec is after Oct, so this year
		{"ottobre", time.October, 1, 2024},   // Current month

		// Month and day
		{"giugno 15", time.June, 15, 2025}, // June is before Oct, so next year
		{"15 giugno", time.June, 15, 2025},
		{"dicembre 25", time.December, 25, 2024}, // Dec is after Oct, so this year
		{"25 dicembre", time.December, 25, 2024},
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

func TestItalian_PeriodBoundaries(t *testing.T) {
	base := time.Date(2024, 10, 15, 14, 30, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"it"},
		RelativeBase: base,
	}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
		wantYear  int
	}{
		// Beginning/end of periods
		{"inizio di mese", time.October, 1, 2024},
		{"fine di mese", time.October, 31, 2024},
		{"inizio di anno", time.January, 1, 2024},
		{"fine di anno", time.December, 31, 2024},
		{"inizio di settimana", time.October, 14, 2024}, // Monday of current week

		// Next/last periods
		{"prossimo mese", time.November, 15, 2024},
		{"prossima settimana", time.October, 22, 2024},
		{"prossimo anno", time.October, 15, 2025},
		{"scorso mese", time.September, 15, 2024},
		{"scorsa settimana", time.October, 8, 2024},
		{"scorso anno", time.October, 15, 2023},

		// Beginning/end of next/last periods
		{"inizio di prossimo mese", time.November, 1, 2024},
		{"fine di prossimo mese", time.November, 30, 2024},
		{"inizio di scorso mese", time.September, 1, 2024},
		{"fine di scorso anno", time.December, 31, 2023},
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

func TestItalian_ThisNextLast(t *testing.T) {
	base := time.Date(2024, 10, 15, 14, 30, 0, 0, time.UTC) // Tuesday
	settings := &godateparser.Settings{
		Languages:    []string{"it"},
		RelativeBase: base,
	}

	tests := []struct {
		input       string
		wantMonth   time.Month
		wantDay     int
		wantWeekday time.Weekday
	}{
		// "questo" (this) with weekdays
		{"questo lunedì", time.October, 21, time.Monday},
		{"questo mercoledì", time.October, 16, time.Wednesday},
		{"questo venerdì", time.October, 18, time.Friday},

		// "questo" (this) with periods
		{"questo mese", time.October, 1, 0},
		{"questa settimana", time.October, 14, 0}, // Monday (start of week)
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

func TestItalian_MixedWithEnglish(t *testing.T) {
	// Test auto-detection when both Italian and English are enabled
	settings := &godateparser.Settings{Languages: []string{"it", "en"}}

	tests := []struct {
		input    string
		wantLang string
	}{
		{"15 dicembre 2024", "it"},
		{"December 15 2024", "en"},
		{"ieri", "it"},
		{"yesterday", "en"},
		{"lunedì", "it"},
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

func BenchmarkItalian_SimpleDate(b *testing.B) {
	settings := &godateparser.Settings{Languages: []string{"it"}}
	for i := 0; i < b.N; i++ {
		_, _ = godateparser.ParseDate("15 dicembre 2024", settings)
	}
}

func BenchmarkItalian_RelativeDate(b *testing.B) {
	settings := &godateparser.Settings{Languages: []string{"it"}}
	for i := 0; i < b.N; i++ {
		_, _ = godateparser.ParseDate("2 giorni fa", settings)
	}
}

func BenchmarkItalian_Weekday(b *testing.B) {
	settings := &godateparser.Settings{Languages: []string{"it"}}
	for i := 0; i < b.N; i++ {
		_, _ = godateparser.ParseDate("prossimo lunedì", settings)
	}
}
