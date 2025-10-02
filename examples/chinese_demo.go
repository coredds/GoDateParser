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
	fmt.Println("=== Chinese Date Parsing Demo ===\n")

	// Set up reference time for relative dates
	refTime := time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		RelativeBase: refTime,
		Languages:    []string{"zh", "en"},
	}

	// Chinese weekdays - full forms (星期)
	fmt.Println("--- Chinese Weekdays (星期) ---")
	chineseWeekdaysFull := []string{
		"星期一", "星期二", "星期三", "星期四",
		"星期五", "星期六", "星期日",
	}
	for _, input := range chineseWeekdaysFull {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("❌ %s: %v\n", input, err)
		} else {
			fmt.Printf("✓ %-8s → %s (%s)\n", input, result.Format("2006-01-02"), result.Weekday())
		}
	}

	// Chinese weekdays - short forms (周)
	fmt.Println("\n--- Chinese Weekdays (周) ---")
	chineseWeekdaysShort := []string{
		"周一", "周二", "周三", "周四",
		"周五", "周六", "周日",
	}
	for _, input := range chineseWeekdaysShort {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("❌ %s: %v\n", input, err)
		} else {
			fmt.Printf("✓ %-6s → %s (%s)\n", input, result.Format("2006-01-02"), result.Weekday())
		}
	}

	// Chinese weekdays - alternative forms (礼拜)
	fmt.Println("\n--- Chinese Weekdays (礼拜) ---")
	chineseWeekdaysAlt := []string{
		"礼拜一", "礼拜二", "礼拜三",
		"礼拜日", "星期天",
	}
	for _, input := range chineseWeekdaysAlt {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("❌ %s: %v\n", input, err)
		} else {
			fmt.Printf("✓ %-8s → %s (%s)\n", input, result.Format("2006-01-02"), result.Weekday())
		}
	}

	// Simple relative dates
	fmt.Println("\n--- Simple Relative Dates ---")
	simpleRelative := []string{
		"昨天", "今天", "明天",
	}
	for _, input := range simpleRelative {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("❌ %s: %v\n", input, err)
		} else {
			fmt.Printf("✓ %-6s → %s\n", input, result.Format("2006-01-02"))
		}
	}

	// Chinese months
	fmt.Println("\n--- Chinese Months ---")
	chineseMonths := []string{
		"1月", "5月", "10月", "12月",
	}
	for _, input := range chineseMonths {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("❌ %s: %v\n", input, err)
		} else {
			fmt.Printf("✓ %-6s → %s\n", input, result.Format("2006-01 (January 2006)"))
		}
	}

	// Time expressions
	fmt.Println("\n--- Time Expressions ---")
	timeExpressions := []string{
		"中午", "午夜", "15:30",
	}
	for _, input := range timeExpressions {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("❌ %s: %v\n", input, err)
		} else {
			fmt.Printf("✓ %-8s → %s\n", input, result.Format("15:04"))
		}
	}

	// Mixed Chinese and English
	fmt.Println("\n--- Mixed Chinese and English ---")
	mixedExamples := []string{
		"昨天",
		"yesterday",
		"今天",
		"today",
		"December 15 2024",
	}
	for _, input := range mixedExamples {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("❌ %-20s: %v\n", input, err)
		} else {
			fmt.Printf("✓ %-20s → %s\n", input, result.Format("2006-01-02"))
		}
	}

	// Language detection
	fmt.Println("\n--- Language Detection ---")
	detectExamples := []string{
		"星期一",
		"周五",
		"礼拜三",
		"昨天",
		"Monday",
		"yesterday",
	}
	for _, input := range detectExamples {
		lang := translations.DetectLanguage(input)
		fmt.Printf("%-10s → %s\n", input, lang)
	}

	// Chinese date formats
	fmt.Println("\n--- Chinese Date Formats (YYYY年MM月DD日) ---")
	dateFormats := []string{
		"2024年12月31日",
		"2025年1月1日",
		"2024年10月15日",
	}
	for _, input := range dateFormats {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("✗ %s: %v\n", input, err)
		} else {
			fmt.Printf("✓ %-20s → %s\n", input, result.Format("2006-01-02"))
		}
	}

	// Relative patterns with numbers
	fmt.Println("\n--- Relative Patterns (N天前, N周后) ---")
	relativePatterns := []string{
		"3天前", // 3 days ago
		"2周后", // in 2 weeks
		"1个月前", // 1 month ago (note: using ヶ or 个 both work)
	}
	for _, input := range relativePatterns {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("✗ %s: %v\n", input, err)
		} else {
			fmt.Printf("✓ %-12s → %s\n", input, result.Format("2006-01-02"))
		}
	}

	// Next/Last patterns
	fmt.Println("\n--- Next/Last Patterns (下周, 上月) ---")
	nextLastPatterns := []string{
		"下周", // next week
		"上周", // last week
		"下月", // next month
		"上月", // last month
	}
	for _, input := range nextLastPatterns {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("✗ %s: %v\n", input, err)
		} else {
			fmt.Printf("✓ %-8s → %s\n", input, result.Format("2006-01-02"))
		}
	}

	// Note about Chinese date formats
	fmt.Println("\n--- Summary ---")
	fmt.Println("Chinese Simplified (zh-Hans) support is now comprehensive!")
	fmt.Println("Supported features:")
	fmt.Println("  ✓ Weekdays: 星期一-日, 周一-日, 礼拜一-日")
	fmt.Println("  ✓ Simple relative: 昨天, 今天, 明天")
	fmt.Println("  ✓ Months: 1月-12月")
	fmt.Println("  ✓ Date format: 2024年12月31日 (YYYY年MM月DD日)")
	fmt.Println("  ✓ Relative patterns: 3天前, 2周后, 1个月前")
	fmt.Println("  ✓ Next/last: 下周, 上周, 下月, 上月")
	fmt.Println("  ✓ Time terms: 中午, 午夜")
	fmt.Println("\nAdvanced patterns (planned for future):")
	fmt.Println("  • Combined weekday modifiers: 下周一 (next Monday)")
	fmt.Println("  • (requires advanced tokenization)")
}
