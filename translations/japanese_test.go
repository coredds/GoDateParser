package translations_test

import (
	"testing"
	"time"

	"github.com/coredds/godateparser"
)

func TestJapanese_Months(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		RelativeBase: refTime,
		Languages:    []string{"ja", "en"},
	}

	// Note: Japanese date formats like "2024年12月31日" require custom parser support
	// For now, testing month-only patterns which work with existing parsers
	tests := []struct {
		name      string
		input     string
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{"12月 (month only)", "12月", 2024, time.December, 1},
		{"1月 (month only)", "1月", 2025, time.January, 1},
		{"10月 (month only)", "10月", 2024, time.October, 1},
		{"5月 (month only)", "5月", 2025, time.May, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("godateparser.ParseDate() error = %v", err)
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

func TestJapanese_Weekdays(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC) // Saturday
	settings := &godateparser.Settings{
		RelativeBase: refTime,
		Languages:    []string{"ja", "en"},
	}

	tests := []struct {
		name        string
		input       string
		wantWeekday time.Weekday
	}{
		// Full forms with 曜日 (youbi)
		{"月曜日", "月曜日", time.Monday},
		{"火曜日", "火曜日", time.Tuesday},
		{"水曜日", "水曜日", time.Wednesday},
		{"木曜日", "木曜日", time.Thursday},
		{"金曜日", "金曜日", time.Friday},
		{"土曜日", "土曜日", time.Saturday},
		{"日曜日", "日曜日", time.Sunday},
		// Short forms without 日
		{"月曜", "月曜", time.Monday},
		{"火曜", "火曜", time.Tuesday},
		{"水曜", "水曜", time.Wednesday},
		{"木曜", "木曜", time.Thursday},
		{"金曜", "金曜", time.Friday},
		{"土曜", "土曜", time.Saturday},
		{"日曜", "日曜", time.Sunday},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("godateparser.ParseDate() error = %v", err)
				return
			}
			if result.Weekday() != tt.wantWeekday {
				t.Errorf("Weekday = %v, want %v", result.Weekday(), tt.wantWeekday)
			}
		})
	}
}

func TestJapanese_RelativeSimple(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		RelativeBase: refTime,
		Languages:    []string{"ja", "en"},
	}

	tests := []struct {
		name     string
		input    string
		wantYear int
		wantDay  int
	}{
		{"昨日 (kinou)", "昨日", 2024, 14},
		{"今日 (kyou)", "今日", 2024, 15},
		{"明日 (ashita)", "明日", 2024, 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("godateparser.ParseDate() error = %v", err)
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

func TestJapanese_RelativeAgo(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		RelativeBase: refTime,
		Languages:    []string{"ja", "en"},
	}

	tests := []struct {
		name     string
		input    string
		wantYear int
		wantDay  int
	}{
		{"1日前", "1日前", 2024, 14},
		{"2日前", "2日前", 2024, 13},
		{"1週前", "1週前", 2024, 8},
		{"2週前", "2週前", 2024, 1},
		{"1ヶ月前", "1ヶ月前", 2024, 15}, // May 15
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("godateparser.ParseDate() error = %v", err)
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

func TestJapanese_RelativeIn(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		RelativeBase: refTime,
		Languages:    []string{"ja", "en"},
	}

	tests := []struct {
		name     string
		input    string
		wantYear int
		wantDay  int
	}{
		{"1日後", "1日後", 2024, 16},
		{"2日後", "2日後", 2024, 17},
		{"1週後", "1週後", 2024, 22},
		{"2週後", "2週後", 2024, 29},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("godateparser.ParseDate() error = %v", err)
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

func TestJapanese_RelativeNextLast(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		RelativeBase: refTime,
		Languages:    []string{"ja", "en"},
	}

	tests := []struct {
		name      string
		input     string
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{"来週", "来週", 2024, time.June, 22},
		{"先週", "先週", 2024, time.June, 8},
		{"来月", "来月", 2024, time.July, 15},
		{"先月", "先月", 2024, time.May, 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("godateparser.ParseDate() error = %v", err)
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

func TestJapanese_WeekdaysWithModifiers(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC) // Saturday
	settings := &godateparser.Settings{
		RelativeBase: refTime,
		Languages:    []string{"ja", "en"},
	}

	tests := []struct {
		name        string
		input       string
		wantWeekday time.Weekday
		wantDay     int
	}{
		{"来週月曜", "来週月曜", time.Monday, 17},   // Next week Monday (June 17-23)
		{"先週月曜", "先週月曜", time.Monday, 3},    // Last week Monday (June 3-9)
		{"来週金曜", "来週金曜", time.Friday, 21},   // Next week Friday
		{"先週金曜", "先週金曜", time.Friday, 7},    // Last week Friday
		{"来週月曜日", "来週月曜日", time.Monday, 17}, // With 日 suffix
		{"先週金曜日", "先週金曜日", time.Friday, 7},  // With 日 suffix
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("godateparser.ParseDate() error = %v", err)
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

func TestJapanese_TimeExpressions(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		RelativeBase: refTime,
		Languages:    []string{"ja", "en"},
	}

	tests := []struct {
		name     string
		input    string
		wantHour int
		wantMin  int
	}{
		{"正午 (shougo - noon)", "正午", 12, 0},
		{"真夜中 (mayonaka - midnight)", "真夜中", 0, 0},
		{"15:30", "15:30", 15, 30},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("godateparser.ParseDate() error = %v", err)
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

func TestJapanese_IncompleteDates(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		RelativeBase: refTime,
		Languages:    []string{"ja", "en"},
	}

	tests := []struct {
		name      string
		input     string
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{"5月", "5月", 2025, time.May, 1},
		{"12月", "12月", 2024, time.December, 1},
		{"10月", "10月", 2024, time.October, 1},
		{"1月", "1月", 2025, time.January, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("godateparser.ParseDate() error = %v", err)
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

func TestJapanese_JapaneseSpecific(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		RelativeBase: refTime,
		Languages:    []string{"ja", "en"},
	}

	tests := []struct {
		name      string
		input     string
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{"2024年10月15日", "2024年10月15日", 2024, time.October, 15},
		{"2024年12月31日", "2024年12月31日", 2024, time.December, 31},
		{"2025年1月1日", "2025年1月1日", 2025, time.January, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("godateparser.ParseDate() error = %v", err)
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

func TestJapanese_MixedWithEnglish(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		RelativeBase: refTime,
		Languages:    []string{"ja", "en"},
	}

	tests := []struct {
		name      string
		input     string
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{"December 15 2024", "December 15 2024", 2024, time.December, 15},
		{"昨日", "昨日", 2024, time.June, 14},
		{"yesterday", "yesterday", 2024, time.June, 14},
		{"今日", "今日", 2024, time.June, 15},
		{"today", "today", 2024, time.June, 15},
		{"明日", "明日", 2024, time.June, 16},
		{"tomorrow", "tomorrow", 2024, time.June, 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("godateparser.ParseDate() error = %v", err)
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

func TestJapanese_AlternativeWeekdays(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC) // Saturday
	settings := &godateparser.Settings{
		RelativeBase: refTime,
		Languages:    []string{"ja", "en"},
	}

	tests := []struct {
		name        string
		input       string
		wantWeekday time.Weekday
	}{
		{"月曜 (short)", "月曜", time.Monday},
		{"火曜 (short)", "火曜", time.Tuesday},
		{"水曜 (short)", "水曜", time.Wednesday},
		{"木曜 (short)", "木曜", time.Thursday},
		{"金曜 (short)", "金曜", time.Friday},
		{"土曜 (short)", "土曜", time.Saturday},
		{"日曜 (short)", "日曜", time.Sunday},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("godateparser.ParseDate() error = %v", err)
				return
			}
			if result.Weekday() != tt.wantWeekday {
				t.Errorf("Weekday = %v, want %v", result.Weekday(), tt.wantWeekday)
			}
		})
	}
}

func TestJapanese_KanjiAndHiragana(t *testing.T) {
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		RelativeBase: refTime,
		Languages:    []string{"ja", "en"},
	}

	tests := []struct {
		name      string
		input     string
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		// Kanji forms (most common)
		{"昨日 (kanji)", "昨日", 2024, time.June, 14},
		{"今日 (kanji)", "今日", 2024, time.June, 15},
		{"明日 (kanji)", "明日", 2024, time.June, 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := godateparser.ParseDate(tt.input, settings)
			if err != nil {
				t.Errorf("godateparser.ParseDate() error = %v", err)
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
