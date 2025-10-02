package godateparser

import (
	"fmt"
	"time"
)

// Error types for specific parsing failures

// ErrInvalidFormat indicates the input string doesn't match any known date format.
type ErrInvalidFormat struct {
	Input      string
	Suggestion string
}

func (e *ErrInvalidFormat) Error() string {
	if e.Suggestion != "" {
		return fmt.Sprintf("invalid date format: %q (suggestion: %s)", e.Input, e.Suggestion)
	}
	return fmt.Sprintf("invalid date format: %q", e.Input)
}

// ErrAmbiguousDate indicates the input could be interpreted in multiple ways.
type ErrAmbiguousDate struct {
	Input      string
	Candidates []time.Time
	Reason     string
}

func (e *ErrAmbiguousDate) Error() string {
	if e.Reason != "" {
		return fmt.Sprintf("ambiguous date: %q (%s)", e.Input, e.Reason)
	}
	return fmt.Sprintf("ambiguous date: %q (use strict mode settings to resolve)", e.Input)
}

// ErrInvalidDate indicates the date components are invalid (e.g., Feb 31).
type ErrInvalidDate struct {
	Input  string
	Year   int
	Month  int
	Day    int
	Reason string
}

func (e *ErrInvalidDate) Error() string {
	if e.Reason != "" {
		return fmt.Sprintf("invalid date: %q - %s", e.Input, e.Reason)
	}
	return fmt.Sprintf("invalid date: %q (year=%d, month=%d, day=%d)", e.Input, e.Year, e.Month, e.Day)
}

// ErrEmptyInput indicates an empty input string was provided.
type ErrEmptyInput struct{}

func (e *ErrEmptyInput) Error() string {
	return "input string is empty"
}

// ErrParseFailure is a generic parse error with context.
type ErrParseFailure struct {
	Input  string
	Parser string // "timestamp", "absolute", "relative"
	Reason error
}

func (e *ErrParseFailure) Error() string {
	if e.Parser != "" {
		return fmt.Sprintf("failed to parse %q with %s parser: %v", e.Input, e.Parser, e.Reason)
	}
	return fmt.Sprintf("failed to parse %q: %v", e.Input, e.Reason)
}

func (e *ErrParseFailure) Unwrap() error {
	return e.Reason
}

// Helper functions to create errors with suggestions

func newInvalidFormatError(input string) error {
	suggestion := suggestFormat(input)
	return &ErrInvalidFormat{
		Input:      input,
		Suggestion: suggestion,
	}
}

func suggestFormat(input string) string {
	// Provide helpful suggestions based on common mistakes
	if input == "" {
		return ""
	}

	// Check for common patterns
	if len(input) == 8 && isAllDigits(input) {
		return "try ISO format: YYYY-MM-DD (e.g., 2024-12-31)"
	}

	if len(input) >= 10 && len(input) <= 13 && isAllDigits(input) {
		return "looks like a timestamp (Unix seconds or milliseconds)"
	}

	// Check for slash/dash confusion
	if containsChar(input, '/') && !containsChar(input, '-') {
		return "numeric dates with slashes (use DateOrder setting: MDY or DMY)"
	}

	return "supported formats: ISO (YYYY-MM-DD), numeric (MM/DD/YYYY), month names (Dec 31 2024), relative (2 days ago)"
}

func isAllDigits(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func containsChar(s string, ch rune) bool {
	for _, c := range s {
		if c == ch {
			return true
		}
	}
	return false
}
