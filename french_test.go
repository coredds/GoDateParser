package godateparser

import (
	"testing"
	"time"
)

func TestFrench_Months(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"fr", "en"},
	}

	tests := []struct {
		name      string
		input     string
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{"31 décembre 2024", "31 décembre 2024", 2024, time.December, 31},
		{"1 janvier 2025", "1 janvier 2025", 2025, time.January, 1},
		{"15 juin 2024", "15 juin 2024", 2024, time.June, 15},
		{"25 déc 2024", "25 déc 2024", 2024, time.December, 25},
		{"mars 15 2024", "mars 15 2024", 2024, time.March, 15},
		{"15 de mars de 2024", "15 de mars de 2024", 2024, time.March, 15},
		{"10 de mai de 2024", "10 de mai de 2024", 2024, time.May, 10},
		{"7 février 2024", "7 février 2024", 2024, time.February, 7},
		// Without accents
		{"31 decembre 2024", "31 decembre 2024", 2024, time.December, 31},
		{"7 fevrier 2024", "7 fevrier 2024", 2024, time.February, 7},
		{"15 aout 2024", "15 aout 2024", 2024, time.August, 15},
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

func TestFrench_Weekdays(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC) // Saturday
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"fr", "en"},
	}

	tests := []struct {
		name        string
		input       string
		wantWeekday time.Weekday
	}{
		{"lundi", "lundi", time.Monday},
		{"mardi", "mardi", time.Tuesday},
		{"mercredi", "mercredi", time.Wednesday},
		{"jeudi", "jeudi", time.Thursday},
		{"vendredi", "vendredi", time.Friday},
		{"samedi", "samedi", time.Saturday},
		{"dimanche", "dimanche", time.Sunday},
		// Abbreviations
		{"lun", "lun", time.Monday},
		{"mar", "mar", time.Tuesday},
		{"mer", "mer", time.Wednesday},
		{"jeu", "jeu", time.Thursday},
		{"ven", "ven", time.Friday},
		{"sam", "sam", time.Saturday},
		{"dim", "dim", time.Sunday},
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

func TestFrench_RelativeSimple(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"fr", "en"},
	}

	tests := []struct {
		name     string
		input    string
		wantYear int
		wantDay  int
	}{
		{"hier", "hier", 2024, 14},
		{"aujourd'hui", "aujourd'hui", 2024, 15},
		{"demain", "demain", 2024, 16},
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

func TestFrench_RelativeAgo(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"fr", "en"},
	}

	tests := []struct {
		name     string
		input    string
		wantYear int
		wantDay  int
	}{
		{"il y a 1 jour", "il y a 1 jour", 2024, 14},
		{"il y a 2 jours", "il y a 2 jours", 2024, 13},
		{"il y a 1 semaine", "il y a 1 semaine", 2024, 8},
		{"il y a 2 semaines", "il y a 2 semaines", 2024, 1},
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

func TestFrench_RelativeIn(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"fr", "en"},
	}

	tests := []struct {
		name     string
		input    string
		wantYear int
		wantDay  int
	}{
		{"dans 1 jour", "dans 1 jour", 2024, 16},
		{"dans 2 jours", "dans 2 jours", 2024, 17},
		{"dans 1 semaine", "dans 1 semaine", 2024, 22},
		{"dans 2 semaines", "dans 2 semaines", 2024, 29},
		{"en 1 jour", "en 1 jour", 2024, 16},
		{"en 2 jours", "en 2 jours", 2024, 17},
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

func TestFrench_RelativeNextLast(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"fr", "en"},
	}

	tests := []struct {
		name      string
		input     string
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{"prochaine semaine", "prochaine semaine", 2024, time.June, 22},
		{"dernière semaine", "dernière semaine", 2024, time.June, 8},
		{"prochain mois", "prochain mois", 2024, time.July, 15},
		{"dernier mois", "dernier mois", 2024, time.May, 15},
		// Without accents
		{"derniere semaine", "derniere semaine", 2024, time.June, 8},
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

func TestFrench_WeekdaysWithModifiers(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC) // Saturday
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"fr", "en"},
	}

	tests := []struct {
		name        string
		input       string
		wantWeekday time.Weekday
		wantDay     int
	}{
		{"prochain lundi", "prochain lundi", time.Monday, 17},
		{"dernier lundi", "dernier lundi", time.Monday, 10},
		{"prochain vendredi", "prochain vendredi", time.Friday, 21},
		{"dernier vendredi", "dernier vendredi", time.Friday, 14},
		{"prochaine lundi", "prochaine lundi", time.Monday, 17},
		{"dernière vendredi", "dernière vendredi", time.Friday, 14},
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

func TestFrench_TimeExpressions(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"fr", "en"},
	}

	tests := []struct {
		name     string
		input    string
		wantHour int
		wantMin  int
	}{
		{"midi", "midi", 12, 0},
		{"minuit", "minuit", 0, 0},
		{"15h30", "15h30", 15, 30},
		{"15h", "15h", 15, 0},
		{"3 heures 30", "3 heures 30", 3, 30},
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

func TestFrench_IncompleteDates(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"fr", "en"},
	}

	tests := []struct {
		name      string
		input     string
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{"mai", "mai", 2025, time.May, 1}, // May has passed (ref: June 15, 2024), so next May is 2025
		{"décembre", "décembre", 2024, time.December, 1},
		{"octobre", "octobre", 2024, time.October, 1},
		{"juin 15", "juin 15", 2024, time.June, 15},
		{"15 juin", "15 juin", 2024, time.June, 15},
		{"décembre 25", "décembre 25", 2024, time.December, 25},
		{"25 décembre", "25 décembre", 2024, time.December, 25},
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

func TestFrench_OrdinalDates(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"fr", "en"},
	}

	tests := []struct {
		name      string
		input     string
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{"juin 3", "juin 3", 2025, time.June, 3}, // June 3 has passed (ref: June 15, 2024), so next is 2025
		{"3 juin", "3 juin", 2025, time.June, 3},
		{"3 de juin", "3 de juin", 2025, time.June, 3},
		{"décembre 31", "décembre 31", 2024, time.December, 31},
		{"31 décembre", "31 décembre", 2024, time.December, 31},
		{"31 de décembre", "31 de décembre", 2024, time.December, 31},
		{"juin 3 2024", "juin 3 2024", 2024, time.June, 3},
		{"3 juin 2024", "3 juin 2024", 2024, time.June, 3},
		{"3 de juin de 2024", "3 de juin de 2024", 2024, time.June, 3},
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

func TestFrench_ThisNextLast(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC) // Saturday
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"fr", "en"},
	}

	tests := []struct {
		name        string
		input       string
		wantWeekday time.Weekday
	}{
		{"ce lundi", "ce lundi", time.Monday},
		{"cet mercredi", "cet mercredi", time.Wednesday},
		{"cette vendredi", "cette vendredi", time.Friday},
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

func TestFrench_ThisWeek(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC) // Saturday
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"fr", "en"},
	}

	result, err := ParseDate("cette semaine", settings)
	if err != nil {
		t.Errorf("ParseDate() error = %v", err)
		return
	}

	// "this week" should return the start of the current week (Monday)
	expectedDay := 10 // Monday, June 10, 2024
	if result.Day() != expectedDay {
		t.Errorf("Day = %v, want %v", result.Day(), expectedDay)
	}
	if result.Weekday() != time.Monday {
		t.Errorf("Weekday = %v, want Monday", result.Weekday())
	}
}

func TestFrench_FrenchSpecific(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"fr", "en"},
		DateOrder:    "DMY",
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

func TestFrench_MixedWithEnglish(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &Settings{
		RelativeBase: refTime,
		Languages:    []string{"fr", "en"},
	}

	tests := []struct {
		name      string
		input     string
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{"15 décembre 2024", "15 décembre 2024", 2024, time.December, 15},
		{"December 15 2024", "December 15 2024", 2024, time.December, 15},
		{"hier", "hier", 2024, time.June, 14},
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

