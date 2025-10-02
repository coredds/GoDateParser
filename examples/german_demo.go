//go:build examples
// +build examples

package main

import (
	"fmt"
	"time"

	"github.com/coredds/godateparser"
)

func main() {
	// Set up a reference time for relative date parsing
	refTime := time.Date(2024, 10, 2, 12, 0, 0, 0, time.UTC)

	// Settings for German language support
	settings := &godateparser.Settings{
		RelativeBase: refTime,
		Languages:    []string{"de", "en"}, // Support German and English
		DateOrder:    "DMY",                // Common in Germany (Day/Month/Year)
	}

	fmt.Println("🇩🇪 German Date Parser Demo")
	fmt.Println("============================\n")

	// Absolute dates with German month names
	fmt.Println("📅 Absolute Dates:")
	examples := []string{
		"31 Dezember 2024",
		"15 Juni 2024",
		"7 Februar 2025",
		"25 Dez 2024",
		"1 Januar 2025",
	}

	for _, input := range examples {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("  ❌ '%s' -> Error: %v\n", input, err)
		} else {
			fmt.Printf("  ✓ '%s' -> %s\n", input, result.Format("2006-01-02"))
		}
	}

	// German weekdays
	fmt.Println("\n📆 German Weekdays:")
	weekdayExamples := []string{
		"Montag",
		"Dienstag",
		"Mittwoch",
		"Donnerstag",
		"Freitag",
		"Samstag",
		"Sonntag",
	}

	for _, input := range weekdayExamples {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("  ❌ '%s' -> Error: %v\n", input, err)
		} else {
			fmt.Printf("  ✓ '%s' -> %s (%s)\n", input, result.Format("2006-01-02"), result.Weekday())
		}
	}

	// Relative dates in German
	fmt.Println("\n⏰ Relative Dates:")
	relativeExamples := []string{
		"gestern",
		"heute",
		"morgen",
		"vor 2 Tagen",
		"in 1 Woche",
		"vor 1 Monat",
	}

	for _, input := range relativeExamples {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("  ❌ '%s' -> Error: %v\n", input, err)
		} else {
			fmt.Printf("  ✓ '%s' -> %s\n", input, result.Format("2006-01-02"))
		}
	}

	// Next/Last modifiers
	fmt.Println("\n🔄 Next/Last Modifiers:")
	modifierExamples := []string{
		"nächster Montag",
		"letzter Freitag",
		"nächste Woche",
		"letzter Monat",
		"nächster Monat",
		"kommender Montag",
		"vergangener Freitag",
	}

	for _, input := range modifierExamples {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("  ❌ '%s' -> Error: %v\n", input, err)
		} else {
			fmt.Printf("  ✓ '%s' -> %s\n", input, result.Format("2006-01-02"))
		}
	}

	// Time expressions
	fmt.Println("\n🕐 Time Expressions:")
	timeExamples := []string{
		"Mittag",
		"Mitternacht",
		"15:30",
	}

	for _, input := range timeExamples {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("  ❌ '%s' -> Error: %v\n", input, err)
		} else {
			fmt.Printf("  ✓ '%s' -> %s\n", input, result.Format("15:04"))
		}
	}

	// Mixed language support
	fmt.Println("\n🌍 Mixed German/English:")
	mixedExamples := []string{
		"15 Dezember 2024",
		"December 15, 2024",
		"gestern",
		"yesterday",
		"nächster Montag",
		"next Monday",
	}

	for _, input := range mixedExamples {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("  ❌ '%s' -> Error: %v\n", input, err)
		} else {
			fmt.Printf("  ✓ '%s' -> %s\n", input, result.Format("2006-01-02"))
		}
	}

	// Numeric dates (DMY format for Germany)
	fmt.Println("\n🔢 Numeric Dates (DMY):")
	numericExamples := []string{
		"15/10/2024",
		"31/12/2024",
		"01/01/2025",
	}

	for _, input := range numericExamples {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("  ❌ '%s' -> Error: %v\n", input, err)
		} else {
			fmt.Printf("  ✓ '%s' -> %s\n", input, result.Format("2006-01-02"))
		}
	}

	// Umlaut tolerance
	fmt.Println("\n🔤 Without Umlauts (also works):")
	umlautExamples := []string{
		"Marz 15 2024",   // März
		"naechste Woche", // nächste
	}

	for _, input := range umlautExamples {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("  ❌ '%s' -> Error: %v\n", input, err)
		} else {
			fmt.Printf("  ✓ '%s' -> %s\n", input, result.Format("2006-01-02"))
		}
	}

	fmt.Println("\n✅ Demo complete!")
}
