package main

import (
	"fmt"
	"time"

	"github.com/coredds/godateparser"
)

func main() {
	// Set up a reference time for relative date parsing
	refTime := time.Date(2024, 10, 2, 12, 0, 0, 0, time.UTC)

	// Settings for French language support
	settings := &godateparser.Settings{
		RelativeBase: refTime,
		Languages:    []string{"fr", "en"}, // Support French and English
		DateOrder:    "DMY",                // Common in France (Day/Month/Year)
	}

	fmt.Println("🇫🇷 French Date Parser Demo")
	fmt.Println("============================\n")

	// Absolute dates with French month names
	fmt.Println("📅 Absolute Dates:")
	examples := []string{
		"31 décembre 2024",
		"15 juin 2024",
		"7 février 2025",
		"25 déc 2024",
		"1 janvier 2025",
	}

	for _, input := range examples {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("  ❌ '%s' -> Error: %v\n", input, err)
		} else {
			fmt.Printf("  ✓ '%s' -> %s\n", input, result.Format("2006-01-02"))
		}
	}

	// French weekdays
	fmt.Println("\n📆 French Weekdays:")
	weekdayExamples := []string{
		"lundi",
		"mardi",
		"mercredi",
		"jeudi",
		"vendredi",
		"samedi",
		"dimanche",
	}

	for _, input := range weekdayExamples {
		result, err := godateparser.ParseDate(input, settings)
		if err != nil {
			fmt.Printf("  ❌ '%s' -> Error: %v\n", input, err)
		} else {
			fmt.Printf("  ✓ '%s' -> %s (%s)\n", input, result.Format("2006-01-02"), result.Weekday())
		}
	}

	// Relative dates in French
	fmt.Println("\n⏰ Relative Dates:")
	relativeExamples := []string{
		"hier",
		"aujourd'hui",
		"demain",
		"il y a 2 jours",
		"dans 1 semaine",
		"il y a 1 mois",
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
		"prochain lundi",
		"dernier vendredi",
		"prochaine semaine",
		"dernier mois",
		"prochain mois",
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
		"midi",
		"minuit",
		"15h30",
		"15h",
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
	fmt.Println("\n🌍 Mixed French/English:")
	mixedExamples := []string{
		"15 décembre 2024",
		"December 15, 2024",
		"hier",
		"yesterday",
		"prochain lundi",
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

	// Numeric dates (DMY format for France)
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

	fmt.Println("\n✅ Demo complete!")
}

