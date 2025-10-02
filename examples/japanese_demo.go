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

	// Note about Japanese date formats
	fmt.Println("\n--- Note ---")
	fmt.Println("Japanese support is now available!")
	fmt.Println("Current support includes:")
	fmt.Println("  Success Weekdays: 月曜日-日曜日 (full), 月曜-日曜 (short)")
	fmt.Println("  Success Simple relative: 昨日 (kinou), 今日 (kyou), 明日 (ashita)")
	fmt.Println("  Success Months: 1月-12月 (ichigatsu-juunigatsu)")
	fmt.Println("  Success Time terms: 正午 (shougo), 真夜中 (mayonaka)")
	fmt.Println("\nPatterns requiring custom parser (coming soon):")
	fmt.Println("  Japanese date format: 2024年12月31日")
	fmt.Println("  Relative patterns: 3日前 (3 days ago), 2週後 (in 2 weeks)")
	fmt.Println("  Next/last: 来週 (next week), 先月 (last month)")
	fmt.Println("  Modified weekdays: 来週月曜日 (next Monday)")
}

