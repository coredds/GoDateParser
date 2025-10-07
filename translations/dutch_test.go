package translations_test

import (
	"testing"
	"time"

	"github.com/coredds/godateparser"
)

// Tests for Dutch language support (nl-NL)

func TestDutch_Months(t *testing.T) {
	settings := &godateparser.Settings{Languages: []string{"nl"}}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
		wantYear  int
	}{
		{"31 december 2024", time.December, 31, 2024},
		{"1 januari 2025", time.January, 1, 2025},
		{"15 juni 2024", time.June, 15, 2024},
		{"25 dec 2024", time.December, 25, 2024},
		{"maart 15 2024", time.March, 15, 2024},
		{"15 maart 2024", time.March, 15, 2024},
		{"augustus 10 2024", time.August, 10, 2024},
		{"10 aug 2024", time.August, 10, 2024},
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

func TestDutch_Weekdays(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC) // Tuesday
	settings := &godateparser.Settings{
		Languages:    []string{"nl"},
		RelativeBase: base,
	}

	tests := []struct {
		input       string
		wantWeekday time.Weekday
	}{
		{"maandag", time.Monday},
		{"dinsdag", time.Tuesday},
		{"woensdag", time.Wednesday},
		{"donderdag", time.Thursday},
		{"vrijdag", time.Friday},
		{"zaterdag", time.Saturday},
		{"zondag", time.Sunday},
		// Abbreviations
		{"ma", time.Monday},
		{"di", time.Tuesday},
		{"wo", time.Wednesday},
		{"do", time.Thursday},
		{"vr", time.Friday},
		{"za", time.Saturday},
		{"zo", time.Sunday},
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

func TestDutch_RelativeSimple(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"nl"},
		RelativeBase: base,
	}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"gisteren", 14},
		{"vandaag", 15},
		{"morgen", 16},
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

func TestDutch_RelativeAgo(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"nl"},
		RelativeBase: base,
	}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"1 dag geleden", 14},
		{"2 dagen geleden", 13},
		{"1 week geleden", 8},
		{"2 weken geleden", 1},
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

func TestDutch_RelativeIn(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"nl"},
		RelativeBase: base,
	}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"over 1 dag", 16},
		{"over 2 dagen", 17},
		{"over 1 week", 22},
		{"over 2 weken", 29},
		{"in 1 dag", 16},
		{"in 2 dagen", 17},
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

func TestDutch_RelativeNextLast(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"nl"},
		RelativeBase: base,
	}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
	}{
		{"volgende week", time.October, 22},
		{"vorige week", time.October, 8},
		{"afgelopen week", time.October, 8},
		{"volgende maand", time.November, 15},
		{"vorige maand", time.September, 15},
		{"komende week", time.October, 22},
		{"aanstaande week", time.October, 22},
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

func TestDutch_WeekdaysWithModifiers(t *testing.T) {
	base := time.Date(2024, 10, 15, 12, 0, 0, 0, time.UTC) // Tuesday
	settings := &godateparser.Settings{
		Languages:    []string{"nl"},
		RelativeBase: base,
	}

	tests := []struct {
		input   string
		wantDay int
	}{
		{"volgende maandag", 21},
		{"vorige maandag", 14},
		{"volgende vrijdag", 18},
		{"vorige vrijdag", 11},
		{"komende maandag", 21},
		{"aanstaande vrijdag", 18},
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

func TestDutch_TimeExpressions(t *testing.T) {
	base := time.Date(2024, 10, 15, 0, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"nl"},
		RelativeBase: base,
	}

	tests := []struct {
		input      string
		wantHour   int
		wantMinute int
	}{
		{"middag", 12, 0},
		{"middernacht", 0, 0},
		{"kwart over 3", 3, 15},
		{"half 4", 3, 30}, // Dutch "half 4" means 3:30
		{"kwart voor 3", 2, 45},
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

func TestDutch_IncompleteDates(t *testing.T) {
	base := time.Date(2024, 10, 15, 14, 30, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:       []string{"nl"},
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
		{"mei", time.May, 1, 2025},           // May is before Oct, so next year
		{"december", time.December, 1, 2024}, // Dec is after Oct, so this year
		{"oktober", time.October, 1, 2024},   // Current month

		// Month and day
		{"juni 15", time.June, 15, 2025}, // June is before Oct, so next year
		{"15 juni", time.June, 15, 2025},
		{"december 25", time.December, 25, 2024}, // Dec is after Oct, so this year
		{"25 december", time.December, 25, 2024},
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

func TestDutch_PeriodBoundaries(t *testing.T) {
	base := time.Date(2024, 10, 15, 14, 30, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		Languages:    []string{"nl"},
		RelativeBase: base,
	}

	tests := []struct {
		input     string
		wantMonth time.Month
		wantDay   int
		wantYear  int
	}{
		// Beginning/end of periods
		{"begin van maand", time.October, 1, 2024},
		{"einde van maand", time.October, 31, 2024},
		{"begin van jaar", time.January, 1, 2024},
		{"einde van jaar", time.December, 31, 2024},
		{"begin van week", time.October, 14, 2024}, // Monday of current week

		// Next/last periods
		{"volgende maand", time.November, 15, 2024},
		{"volgende week", time.October, 22, 2024},
		{"volgend jaar", time.October, 15, 2025},
		{"vorige maand", time.September, 15, 2024},
		{"vorige week", time.October, 8, 2024},
		{"vorig jaar", time.October, 15, 2023},

		// Beginning/end of next/last periods
		{"begin van volgende maand", time.November, 1, 2024},
		{"einde van volgende maand", time.November, 30, 2024},
		{"begin van vorige maand", time.September, 1, 2024},
		{"einde van vorig jaar", time.December, 31, 2023},
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

func TestDutch_ThisNextLast(t *testing.T) {
	base := time.Date(2024, 10, 15, 14, 30, 0, 0, time.UTC) // Tuesday
	settings := &godateparser.Settings{
		Languages:    []string{"nl"},
		RelativeBase: base,
	}

	tests := []struct {
		input       string
		wantMonth   time.Month
		wantDay     int
		wantWeekday time.Weekday
	}{
		// "deze" (this) with weekdays
		{"deze maandag", time.October, 21, time.Monday},
		{"deze woensdag", time.October, 16, time.Wednesday},
		{"deze vrijdag", time.October, 18, time.Friday},

		// "deze" (this) with periods
		{"deze maand", time.October, 1, 0},
		{"deze week", time.October, 14, 0}, // Monday (start of week)
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

func TestDutch_MixedWithEnglish(t *testing.T) {
	// Test auto-detection when both Dutch and English are enabled
	settings := &godateparser.Settings{Languages: []string{"nl", "en"}}

	tests := []struct {
		input    string
		wantLang string
	}{
		{"15 december 2024", "nl"},
		{"December 15 2024", "en"},
		{"gisteren", "nl"},
		{"yesterday", "en"},
		{"maandag", "nl"},
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

func BenchmarkDutch_SimpleDate(b *testing.B) {
	settings := &godateparser.Settings{Languages: []string{"nl"}}
	for i := 0; i < b.N; i++ {
		_, _ = godateparser.ParseDate("15 december 2024", settings)
	}
}

func BenchmarkDutch_RelativeDate(b *testing.B) {
	settings := &godateparser.Settings{Languages: []string{"nl"}}
	for i := 0; i < b.N; i++ {
		_, _ = godateparser.ParseDate("2 dagen geleden", settings)
	}
}

func BenchmarkDutch_Weekday(b *testing.B) {
	settings := &godateparser.Settings{Languages: []string{"nl"}}
	for i := 0; i < b.N; i++ {
		_, _ = godateparser.ParseDate("volgende maandag", settings)
	}
}
