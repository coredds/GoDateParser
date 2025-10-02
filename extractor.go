package godateparser

import (
	"regexp"
	"strings"
)

// Date extraction patterns for scanning text
var extractionPatterns = []*regexp.Regexp{
	// ISO dates
	regexp.MustCompile(`\b\d{4}-\d{1,2}-\d{1,2}(?:[T\s]\d{1,2}:\d{1,2}(?::\d{1,2})?)?\b`),
	// Numeric dates: 12/31/2024, 31-12-2024
	regexp.MustCompile(`\b\d{1,2}[/-]\d{1,2}[/-]\d{4}\b`),
	// Month name dates: "December 31, 2024", "31 Dec 2024"
	regexp.MustCompile(`(?i)\b\d{1,2}\s+(?:Jan(?:uary)?|Feb(?:ruary)?|Mar(?:ch)?|Apr(?:il)?|May|Jun(?:e)?|Jul(?:y)?|Aug(?:ust)?|Sep(?:tember)?|Oct(?:ober)?|Nov(?:ember)?|Dec(?:ember)?)[,\s]+\d{4}\b`),
	regexp.MustCompile(`(?i)\b(?:Jan(?:uary)?|Feb(?:ruary)?|Mar(?:ch)?|Apr(?:il)?|May|Jun(?:e)?|Jul(?:y)?|Aug(?:ust)?|Sep(?:tember)?|Oct(?:ober)?|Nov(?:ember)?|Dec(?:ember)?)\s+\d{1,2}[,\s]+\d{4}\b`),
	// Relative dates
	regexp.MustCompile(`(?i)\b\d+\s+(?:second|minute|hour|day|week|month|year)s?\s+ago\b`),
	regexp.MustCompile(`(?i)\bin\s+\d+\s+(?:second|minute|hour|day|week|month|year)s?\b`),
	regexp.MustCompile(`(?i)\b(?:yesterday|today|tomorrow)\b`),
	regexp.MustCompile(`(?i)\b(?:last|next)\s+(?:week|month|year)\b`),
	regexp.MustCompile(`(?i)\b(?:next|last)\s+(?:monday|tuesday|wednesday|thursday|friday|saturday|sunday)\b`),
	// Timestamps
	regexp.MustCompile(`\b\d{10,13}\b`),
}

// extractAllDates scans text and extracts all date occurrences.
func extractAllDates(ctx *parserContext) ([]ParsedDate, error) {
	var results []ParsedDate
	text := ctx.input

	// Track processed positions to avoid duplicates
	processed := make(map[int]bool)

	for _, pattern := range extractionPatterns {
		matches := pattern.FindAllStringIndex(text, -1)
		for _, match := range matches {
			start := match[0]
			end := match[1]

			// Skip if already processed
			if processed[start] {
				continue
			}

			matchedText := text[start:end]

			// Try to parse the matched text
			parsedDate, err := ParseDate(matchedText, ctx.settings)
			if err == nil {
				results = append(results, ParsedDate{
					Date:        parsedDate,
					Position:    start,
					Length:      end - start,
					MatchedText: matchedText,
					Confidence:  calculateConfidence(matchedText),
				})
				processed[start] = true
			}
		}
	}

	return results, nil
}

// calculateConfidence estimates the confidence of a date match.
func calculateConfidence(text string) float64 {
	text = strings.TrimSpace(text)

	// ISO format gets highest confidence
	if regexp.MustCompile(`^\d{4}-\d{2}-\d{2}`).MatchString(text) {
		return 0.95
	}

	// Month names get high confidence
	if regexp.MustCompile(`(?i)(january|february|march|april|may|june|july|august|september|october|november|december)`).MatchString(text) {
		return 0.90
	}

	// Relative dates with numbers
	if regexp.MustCompile(`\d+\s+(day|week|month|year)s?\s+ago`).MatchString(text) {
		return 0.85
	}

	// Common relative terms
	if regexp.MustCompile(`(?i)(yesterday|today|tomorrow|next|last)`).MatchString(text) {
		return 0.80
	}

	// Numeric dates (more ambiguous)
	if regexp.MustCompile(`^\d{1,2}[/-]\d{1,2}[/-]\d{4}$`).MatchString(text) {
		return 0.75
	}

	// Timestamps
	if regexp.MustCompile(`^\d{10,13}$`).MatchString(text) {
		return 0.70
	}

	return 0.60
}
