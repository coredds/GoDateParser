//go:build examples
// +build examples

package main

import (
	"fmt"
	"time"

	"github.com/coredds/godateparser"
	"github.com/coredds/godateparser/translations"
)

func main() {
	fmt.Println("=== Japanese Date Parsing Demo ===\n")

	// Set up reference time for relative dates
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		RelativeBase: refTime,
		Languages:    []string{"ja", "en"},
	}

	// Japanese weekdays - full forms (曜日)
	fmt.Println("--- Japanese Weekdays (youbi) ---")
	japaneseWeekdaysFull := []string{
		"月曜日", "火曜日", "水曜日", "木曜日",
		"金曜日", "土曜日", "日曜日",
	}
	for _, input := range japaneseWeekdaysFull {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("Error %s: %v\n", input, err)
		} else {
			fmt.Printf("Success %-8s -> %s (%s)\n", input, result.Format("2006-01-02"), result.Weekday())
		}
	}

	// Japanese weekdays - short forms (曜)
	fmt.Println("\n--- Japanese Weekdays (short) ---")
	japaneseWeekdaysShort := []string{
		"月曜", "火曜", "水曜", "木曜",
		"金曜", "土曜", "日曜",
	}
	for _, input := range japaneseWeekdaysShort {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("Error %s: %v\n", input, err)
		} else {
			fmt.Printf("Success %-6s -> %s (%s)\n", input, result.Format("2006-01-02"), result.Weekday())
		}
	}

	// Simple relative dates
	fmt.Println("\n--- Simple Relative Dates ---")
	simpleRelative := []string{
		"昨日", // kinou (yesterday)
		"今日", // kyou (today)
		"明日", // ashita (tomorrow)
	}
	for _, input := range simpleRelative {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("Error %s: %v\n", input, err)
		} else {
			fmt.Printf("Success %-6s -> %s\n", input, result.Format("2006-01-02"))
		}
	}

	// Japanese months
	fmt.Println("\n--- Japanese Months ---")
	japaneseMonths := []string{
		"1月", "5月", "10月", "12月",
	}
	for _, input := range japaneseMonths {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("Error %s: %v\n", input, err)
		} else {
			fmt.Printf("Success %-6s -> %s\n", input, result.Format("2006-01 (January 2006)"))
		}
	}

	// Time expressions
	fmt.Println("\n--- Time Expressions ---")
	timeExpressions := []string{
		"正午",   // shougo (noon)
		"真夜中", // mayonaka (midnight)
		"15:30",
	}
	for _, input := range timeExpressions {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("Error %s: %v\n", input, err)
		} else {
			fmt.Printf("Success %-8s -> %s\n", input, result.Format("15:04"))
		}
	}

	// Mixed Japanese and English
	fmt.Println("\n--- Mixed Japanese and English ---")
	mixedExamples := []string{
		"昨日",
		"yesterday",
		"今日",
		"today",
		"明日",
		"tomorrow",
		"December 15 2024",
	}
	for _, input := range mixedExamples {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("Error %-20s: %v\n", input, err)
		} else {
			fmt.Printf("Success %-20s -> %s\n", input, result.Format("2006-01-02"))
		}
	}

	// Language detection
	fmt.Println("\n--- Language Detection ---")
	detectExamples := []string{
		"月曜日",
		"火曜",
		"昨日",
		"今日",
		"Monday",
		"yesterday",
	}
	for _, input := range detectExamples {
		lang := translations.DetectLanguage(input)
		fmt.Printf("%-10s -> %s\n", input, lang)
	}

	// Weekday comparison
	fmt.Println("\n--- Weekday Forms Comparison ---")
	fmt.Println("Full form:  月曜日 (getsuyoubi)")
	fmt.Println("Short form: 月曜   (getsuyo)")
	fmt.Println("Both forms work the same way!")

	// Japanese date formats
	fmt.Println("\n--- Japanese Date Formats (YYYY年MM月DD日) ---")
	dateFormats := []string{
		"2024年12月31日",
		"2025年1月1日",
		"2024年10月15日",
	}
	for _, input := range dateFormats {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("Error %s: %v\n", input, err)
		} else {
			fmt.Printf("Success %-20s -> %s\n", input, result.Format("2006-01-02"))
		}
	}

	// Relative patterns with numbers
	fmt.Println("\n--- Relative Patterns (N日前, N週後) ---")
	relativePatterns := []string{
		"3日前",   // 3 days ago
		"2週後",   // in 2 weeks
		"1ヶ月前", // 1 month ago
	}
	for _, input := range relativePatterns {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("Error %s: %v\n", input, err)
		} else {
			fmt.Printf("Success %-12s -> %s\n", input, result.Format("2006-01-02"))
		}
	}

	// Next/Last patterns
	fmt.Println("\n--- Next/Last Patterns (来週, 先月) ---")
	nextLastPatterns := []string{
		"来週", // next week
		"先週", // last week
		"来月", // next month
		"先月", // last month
	}
	for _, input := range nextLastPatterns {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("Error %s: %v\n", input, err)
		} else {
			fmt.Printf("Success %-8s -> %s\n", input, result.Format("2006-01-02"))
		}
	}

	// Note about Japanese date formats
	fmt.Println("\n--- Summary ---")
	fmt.Println("Japanese support is now comprehensive!")
	fmt.Println("Supported features:")
	fmt.Println("  Success Weekdays: 月曜日-日曜日 (full), 月曜-日曜 (short)")
	fmt.Println("  Success Simple relative: 昨日 (kinou), 今日 (kyou), 明日 (ashita)")
	fmt.Println("  Success Months: 1月-12月 (ichigatsu-juunigatsu)")
	fmt.Println("  Success Date format: 2024年12月31日 (YYYY年MM月DD日)")
	fmt.Println("  Success Relative patterns: 3日前, 2週後, 1ヶ月前")
	fmt.Println("  Success Next/last: 来週, 先週, 来月, 先月")
	fmt.Println("  Success Time terms: 正午 (shougo), 真夜中 (mayonaka)")
	fmt.Println("\nAdvanced patterns (planned for future):")
	fmt.Println("  Combined weekday modifiers: 来週月曜日 (next Monday)")
	fmt.Println("  (requires advanced tokenization)")
}

