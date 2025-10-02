package main

import (
	"fmt"
	"time"

	godateparser "github.com/coredds/GoDateParser"
)

func main() {
	fmt.Println("=== GoDateParser Examples ===")
	fmt.Println()

	// Example 1: Basic absolute date parsing
	fmt.Println("1. Absolute Date Parsing:")
	examples1 := []string{
		"2024-12-31",
		"12/31/2024",
		"December 31, 2024",
		"31 Dec 2024",
	}
	for _, input := range examples1 {
		date, err := godateparser.ParseDate(input, nil)
		if err != nil {
			fmt.Printf("  ❌ '%s' -> Error: %v\n", input, err)
		} else {
			fmt.Printf("  ✓ '%s' -> %s\n", input, date.Format("2006-01-02"))
		}
	}
	fmt.Println()

	// Example 2: Relative date parsing
	fmt.Println("2. Relative Date Parsing:")
	base := time.Date(2024, 10, 2, 12, 0, 0, 0, time.UTC)
	settings := &godateparser.Settings{
		RelativeBase: base,
	}
	examples2 := []string{
		"yesterday",
		"today",
		"tomorrow",
		"2 days ago",
		"in 3 weeks",
		"last month",
		"next Monday",
	}
	fmt.Printf("  Base date: %s\n", base.Format("2006-01-02"))
	for _, input := range examples2 {
		date, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("  ❌ '%s' -> Error: %v\n", input, err)
		} else {
			fmt.Printf("  ✓ '%s' -> %s\n", input, date.Format("2006-01-02"))
		}
	}
	fmt.Println()

	// Example 3: Unix timestamp parsing
	fmt.Println("3. Unix Timestamp Parsing:")
	examples3 := []string{
		"1609459200",    // 2021-01-01 in seconds
		"1609459200000", // 2021-01-01 in milliseconds
	}
	for _, input := range examples3 {
		date, err := godateparser.ParseDate(input, nil)
		if err != nil {
			fmt.Printf("  ❌ '%s' -> Error: %v\n", input, err)
		} else {
			fmt.Printf("  ✓ '%s' -> %s\n", input, date.Format("2006-01-02 15:04:05"))
		}
	}
	fmt.Println()

	// Example 4: Date extraction from text
	fmt.Println("4. Extract Dates from Text:")
	text := "The project started on December 31, 2024 and the deadline is 2025-06-15. " +
		"We had a meeting yesterday and the next one is in 2 weeks."
	fmt.Printf("  Text: %s\n\n", text)

	results, err := godateparser.ExtractDates(text, settings)
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
	} else {
		fmt.Printf("  Found %d date(s):\n", len(results))
		for i, result := range results {
			fmt.Printf("    %d. '%s' at position %d\n", i+1, result.MatchedText, result.Position)
			fmt.Printf("       Parsed: %s (confidence: %.2f)\n",
				result.Date.Format("2006-01-02"), result.Confidence)
		}
	}
	fmt.Println()

	// Example 5: Custom settings with date order
	fmt.Println("5. Date Order Preferences:")
	dateStr := "01/02/2024"

	settingsMDY := &godateparser.Settings{DateOrder: "MDY"}
	dateMDY, _ := godateparser.ParseDate(dateStr, settingsMDY)
	fmt.Printf("  '%s' with MDY -> %s (January 2nd)\n", dateStr, dateMDY.Format("2006-01-02"))

	settingsDMY := &godateparser.Settings{DateOrder: "DMY"}
	dateDMY, _ := godateparser.ParseDate(dateStr, settingsDMY)
	fmt.Printf("  '%s' with DMY -> %s (February 1st)\n", dateStr, dateDMY.Format("2006-01-02"))
	fmt.Println()

	// Example 6: Enable/disable specific parsers
	fmt.Println("6. Selective Parser Enabling:")
	settingsOnlyAbsolute := &godateparser.Settings{
		EnableParsers: []string{"absolute"},
	}

	// This should work
	date1, err1 := godateparser.ParseDate("2024-12-31", settingsOnlyAbsolute)
	if err1 == nil {
		fmt.Printf("  ✓ '2024-12-31' with absolute parser -> %s\n", date1.Format("2006-01-02"))
	}

	// This should fail (relative parser disabled)
	_, err2 := godateparser.ParseDate("yesterday", settingsOnlyAbsolute)
	if err2 != nil {
		fmt.Printf("  ✓ 'yesterday' correctly rejected (relative parser disabled)\n")
	}
	fmt.Println()

	// Example 7: Multi-language support
	fmt.Println("7. Multi-Language Support:")

	// Portuguese
	settingsPT := &godateparser.Settings{Languages: []string{"pt"}}
	examplesPT := []string{
		"15 de junho de 2024",
		"ontem",
		"próxima segunda",
		"há 2 dias",
	}
	fmt.Println("  Portuguese:")
	for _, input := range examplesPT {
		date, err := godateparser.ParseDate(input, settingsPT)
		if err != nil {
			fmt.Printf("    ❌ '%s' -> Error: %v\n", input, err)
		} else {
			fmt.Printf("    ✓ '%s' -> %s\n", input, date.Format("2006-01-02"))
		}
	}

	// Spanish
	settingsES := &godateparser.Settings{Languages: []string{"es"}}
	examplesES := []string{
		"15 de junio de 2024",
		"ayer",
		"próximo lunes",
		"hace 2 días",
	}
	fmt.Println("  Spanish:")
	for _, input := range examplesES {
		date, err := godateparser.ParseDate(input, settingsES)
		if err != nil {
			fmt.Printf("    ❌ '%s' -> Error: %v\n", input, err)
		} else {
			fmt.Printf("    ✓ '%s' -> %s\n", input, date.Format("2006-01-02"))
		}
	}

	// French
	settingsFR := &godateparser.Settings{Languages: []string{"fr"}}
	examplesFR := []string{
		"15 juin 2024",
		"hier",
		"prochain lundi",
		"il y a 2 jours",
	}
	fmt.Println("  French:")
	for _, input := range examplesFR {
		date, err := godateparser.ParseDate(input, settingsFR)
		if err != nil {
			fmt.Printf("    ❌ '%s' -> Error: %v\n", input, err)
		} else {
			fmt.Printf("    ✓ '%s' -> %s\n", input, date.Format("2006-01-02"))
		}
	}

	// Auto-detection
	settingsMulti := &godateparser.Settings{Languages: []string{"pt", "es", "fr", "en"}}
	examplesMulti := []string{
		"amanhã",     // Portuguese
		"mañana",     // Spanish
		"demain",     // French
		"tomorrow",   // English
	}
	fmt.Println("  Auto-detection (pt, es, fr, en):")
	for _, input := range examplesMulti {
		date, err := godateparser.ParseDate(input, settingsMulti)
		if err != nil {
			fmt.Printf("    ❌ '%s' -> Error: %v\n", input, err)
		} else {
			fmt.Printf("    ✓ '%s' -> %s\n", input, date.Format("2006-01-02"))
		}
	}
	fmt.Println()

	fmt.Println("=== Examples Complete ===")
}
