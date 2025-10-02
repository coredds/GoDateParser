package godateparser

import (
	"testing"
	"time"
)

func TestGerman_Months(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"de", "en"},
	}

	tests := []struct {
		name      string
		input     string
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{"31 Dezember 2024", "31 Dezember 2024", 2024, time.December, 31},
		{"1 Januar 2025", "1 Januar 2025", 2025, time.January, 1},
		{"15 Juni 2024", "15 Juni 2024", 2024, time.June, 15},
		{"25 Dez 2024", "25 Dez 2024", 2024, time.December, 25},
		{"März 15 2024", "März 15 2024", 2024, time.March, 15},
		{"7 Februar 2024", "7 Februar 2024", 2024, time.February, 7},
		// Without umlaut
		{"März as Marz", "15 Marz 2024", 2024, time.March, 15},
		{"15 august 2024", "15 august 2024", 2024, time.August, 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("ParseDate() error = %v", err)
				return
			}
			if result.Year() != tt.wantYear {
				t.Errorf("Year = %v, want %v", result.Year(), tt.wantYear)
			}
			if result.Month() != tt.wantMonth {
				t.Errorf("Month = %v, want %v", result.Month(), tt.wantMonth)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("Day = %v, want %v", result.Day(), tt.wantDay)
			}
		})
	}
}

func TestGerman_Weekdays(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC) // Saturday
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"de", "en"},
	}

	tests := []struct {
		name        string
		input       string
		wantWeekday time.Weekday
	}{
		{"Montag", "Montag", time.Monday},
		{"Dienstag", "Dienstag", time.Tuesday},
		{"Mittwoch", "Mittwoch", time.Wednesday},
		{"Donnerstag", "Donnerstag", time.Thursday},
		{"Freitag", "Freitag", time.Friday},
		{"Samstag", "Samstag", time.Saturday},
		{"Sonntag", "Sonntag", time.Sunday},
		// Abbreviations
		{"Mo", "Mo", time.Monday},
		{"Di", "Di", time.Tuesday},
		{"Mi", "Mi", time.Wednesday},
		{"Do", "Do", time.Thursday},
		{"Fr", "Fr", time.Friday},
		{"Sa", "Sa", time.Saturday},
		{"So", "So", time.Sunday},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("ParseDate() error = %v", err)
				return
			}
			if result.Weekday() != tt.wantWeekday {
				t.Errorf("Weekday = %v, want %v", result.Weekday(), tt.wantWeekday)
			}
		})
	}
}

func TestGerman_RelativeSimple(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"de", "en"},
	}

	tests := []struct {
		name     string
		input    string
		wantYear int
		wantDay  int
	}{
		{"gestern", "gestern", 2024, 14},
		{"heute", "heute", 2024, 15},
		{"morgen", "morgen", 2024, 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("ParseDate() error = %v", err)
				return
			}
			if result.Year() != tt.wantYear {
				t.Errorf("Year = %v, want %v", result.Year(), tt.wantYear)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("Day = %v, want %v", result.Day(), tt.wantDay)
			}
		})
	}
}

func TestGerman_RelativeAgo(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"de", "en"},
	}

	tests := []struct {
		name     string
		input    string
		wantYear int
		wantDay  int
	}{
		{"vor 1 Tag", "vor 1 Tag", 2024, 14},
		{"vor 2 Tagen", "vor 2 Tagen", 2024, 13},
		{"vor 1 Woche", "vor 1 Woche", 2024, 8},
		{"vor 2 Wochen", "vor 2 Wochen", 2024, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("ParseDate() error = %v", err)
				return
			}
			if result.Year() != tt.wantYear {
				t.Errorf("Year = %v, want %v", result.Year(), tt.wantYear)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("Day = %v, want %v", result.Day(), tt.wantDay)
			}
		})
	}
}

func TestGerman_RelativeIn(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"de", "en"},
	}

	tests := []struct {
		name     string
		input    string
		wantYear int
		wantDay  int
	}{
		{"in 1 Tag", "in 1 Tag", 2024, 16},
		{"in 2 Tagen", "in 2 Tagen", 2024, 17},
		{"in 1 Woche", "in 1 Woche", 2024, 22},
		{"in 2 Wochen", "in 2 Wochen", 2024, 29},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("ParseDate() error = %v", err)
				return
			}
			if result.Year() != tt.wantYear {
				t.Errorf("Year = %v, want %v", result.Year(), tt.wantYear)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("Day = %v, want %v", result.Day(), tt.wantDay)
			}
		})
	}
}

func TestGerman_RelativeNextLast(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"de", "en"},
	}

	tests := []struct {
		name      string
		input     string
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{"nächste Woche", "nächste Woche", 2024, time.June, 22},
		{"letzte Woche", "letzte Woche", 2024, time.June, 8},
		{"nächster Monat", "nächster Monat", 2024, time.July, 15},
		{"letzter Monat", "letzter Monat", 2024, time.May, 15},
		// Without umlaut
		{"naechste Woche", "naechste Woche", 2024, time.June, 22},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("ParseDate() error = %v", err)
				return
			}
			if result.Year() != tt.wantYear {
				t.Errorf("Year = %v, want %v", result.Year(), tt.wantYear)
			}
			if result.Month() != tt.wantMonth {
				t.Errorf("Month = %v, want %v", result.Month(), tt.wantMonth)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("Day = %v, want %v", result.Day(), tt.wantDay)
			}
		})
	}
}

func TestGerman_WeekdaysWithModifiers(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC) // Saturday
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"de", "en"},
	}

	tests := []struct {
		name        string
		input       string
		wantWeekday time.Weekday
		wantDay     int
	}{
		{"nächster Montag", "nächster Montag", time.Monday, 17},
		{"letzter Montag", "letzter Montag", time.Monday, 10},
		{"nächster Freitag", "nächster Freitag", time.Friday, 21},
		{"letzter Freitag", "letzter Freitag", time.Friday, 14},
		{"kommender Montag", "kommender Montag", time.Monday, 17},
		{"vergangener Freitag", "vergangener Freitag", time.Friday, 14},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("ParseDate() error = %v", err)
				return
			}
			if result.Weekday() != tt.wantWeekday {
				t.Errorf("Weekday = %v, want %v", result.Weekday(), tt.wantWeekday)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("Day = %v, want %v", result.Day(), tt.wantDay)
			}
		})
	}
}

func TestGerman_TimeExpressions(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"de", "en"},
	}

	tests := []struct {
		name     string
		input    string
		wantHour int
		wantMin  int
	}{
		{"Mittag", "Mittag", 12, 0},
		{"Mitternacht", "Mitternacht", 0, 0},
		{"15:30", "15:30", 15, 30},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("ParseDate() error = %v", err)
				return
			}
			if result.Hour() != tt.wantHour {
				t.Errorf("Hour = %v, want %v", result.Hour(), tt.wantHour)
			}
			if result.Minute() != tt.wantMin {
				t.Errorf("Minute = %v, want %v", result.Minute(), tt.wantMin)
			}
		})
	}
}

func TestGerman_IncompleteDates(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"de", "en"},
	}

	tests := []struct {
		name      string
		input     string
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{"Mai", "Mai", 2025, time.May, 1},
		{"Dezember", "Dezember", 2024, time.December, 1},
		{"Oktober", "Oktober", 2024, time.October, 1},
		{"Juni 15", "Juni 15", 2024, time.June, 15},
		{"15 Juni", "15 Juni", 2024, time.June, 15},
		{"Dezember 25", "Dezember 25", 2024, time.December, 25},
		{"25 Dezember", "25 Dezember", 2024, time.December, 25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("ParseDate() error = %v", err)
				return
			}
			if result.Year() != tt.wantYear {
				t.Errorf("Year = %v, want %v", result.Year(), tt.wantYear)
			}
			if result.Month() != tt.wantMonth {
				t.Errorf("Month = %v, want %v", result.Month(), tt.wantMonth)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("Day = %v, want %v", result.Day(), tt.wantDay)
			}
		})
	}
}

func TestGerman_OrdinalDates(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"de", "en"},
	}

	tests := []struct {
		name      string
		input     string
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{"Juni 3", "Juni 3", 2025, time.June, 3},
		{"3 Juni", "3 Juni", 2025, time.June, 3},
		{"Dezember 31", "Dezember 31", 2024, time.December, 31},
		{"31 Dezember", "31 Dezember", 2024, time.December, 31},
		{"Juni 3 2024", "Juni 3 2024", 2024, time.June, 3},
		{"3 Juni 2024", "3 Juni 2024", 2024, time.June, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("ParseDate() error = %v", err)
				return
			}
			if result.Year() != tt.wantYear {
				t.Errorf("Year = %v, want %v", result.Year(), tt.wantYear)
			}
			if result.Month() != tt.wantMonth {
				t.Errorf("Month = %v, want %v", result.Month(), tt.wantMonth)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("Day = %v, want %v", result.Day(), tt.wantDay)
			}
		})
	}
}

func TestGerman_ThisNextLast(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC) // Saturday
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"de", "en"},
	}

	tests := []struct {
		name        string
		input       string
		wantWeekday time.Weekday
	}{
		{"dieser Montag", "dieser Montag", time.Monday},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("ParseDate() error = %v", err)
				return
			}
			if result.Weekday() != tt.wantWeekday {
				t.Errorf("Weekday = %v, want %v", result.Weekday(), tt.wantWeekday)
			}
		})
	}
}

func TestGerman_WithoutUmlaut(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"de", "en"},
	}

	tests := []struct {
		name      string
		input     string
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{"Marz instead of März", "15 Marz 2024", 2024, time.March, 15},
		{"naechste instead of nächste", "naechste Woche", 2024, time.June, 22},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("ParseDate() error = %v", err)
				return
			}
			if result.Year() != tt.wantYear {
				t.Errorf("Year = %v, want %v", result.Year(), tt.wantYear)
			}
			if result.Month() != tt.wantMonth {
				t.Errorf("Month = %v, want %v", result.Month(), tt.wantMonth)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("Day = %v, want %v", result.Day(), tt.wantDay)
			}
		})
	}
}

func TestGerman_GermanSpecific(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"de", "en"},
		DateOrder:    "DMY", // Common in Germany
	}

	tests := []struct {
		name      string
		input     string
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{"15/10/2024", "15/10/2024", 2024, time.October, 15},
		{"31/12/2024", "31/12/2024", 2024, time.December, 31},
		{"01/01/2025", "01/01/2025", 2025, time.January, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("ParseDate() error = %v", err)
				return
			}
			if result.Year() != tt.wantYear {
				t.Errorf("Year = %v, want %v", result.Year(), tt.wantYear)
			}
			if result.Month() != tt.wantMonth {
				t.Errorf("Month = %v, want %v", result.Month(), tt.wantMonth)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("Day = %v, want %v", result.Day(), tt.wantDay)
			}
		})
	}
}

func TestGerman_MixedWithEnglish(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"de", "en"},
	}

	tests := []struct {
		name      string
		input     string
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{"15 Dezember 2024", "15 Dezember 2024", 2024, time.December, 15},
		{"December 15 2024", "December 15 2024", 2024, time.December, 15},
		{"gestern", "gestern", 2024, time.June, 14},
		{"yesterday", "yesterday", 2024, time.June, 14},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("ParseDate() error = %v", err)
				return
			}
			if result.Year() != tt.wantYear {
				t.Errorf("Year = %v, want %v", result.Year(), tt.wantYear)
			}
			if result.Month() != tt.wantMonth {
				t.Errorf("Month = %v, want %v", result.Month(), tt.wantMonth)
			}
			if result.Day() != tt.wantDay {
				t.Errorf("Day = %v, want %v", result.Day(), tt.wantDay)
			}
		})
	}
}
