// Package godateparser provides natural language date parsing for Go.
// It supports parsing human-readable date strings in multiple languages and formats,
// including absolute dates, relative dates, timestamps, and timezone-aware dates.
package godateparser

import (
	"errors"
	"time"

	"github.com/coredds/godateparser/translations"
)

// Settings defines customizable parsing behavior for date parsing operations.
type Settings struct {
	// DateOrder specifies the date component order preference: "YMD", "MDY", or "DMY"
	DateOrder string

	// Languages specifies preferred languages/locales for parsing (e.g., ["en", "es", "fr"])
	// If empty, all languages are considered with autodetection
	Languages []string

	// RelativeBase is the base date/time for relative date calculations
	// If zero, time.Now() is used
	RelativeBase time.Time

	// EnableParsers specifies which parsers to enable
	// Available: "timestamp", "relative", "absolute", "timezone"
	// If empty, all parsers are enabled
	EnableParsers []string

	// Strict mode returns error if parsing is ambiguous
	Strict bool

	// PreferredTimezone specifies the default timezone for dates without explicit timezone
	PreferredTimezone *time.Location

	// PreferDatesFrom specifies preference for ambiguous dates: "past", "future", or ""
	// When set to "future", ambiguous dates like "Monday" prefer next Monday
	// When set to "past", ambiguous dates like "Monday" prefer last Monday
	// When empty, defaults to "future" for forward-looking dates
	PreferDatesFrom string
}

// ParsedDate represents a date extracted from text with its position information.
type ParsedDate struct {
	// Date is the parsed date/time value
	Date time.Time

	// Position is the start index of the matched date string in the input text
	Position int

	// Length is the length of the matched date substring
	Length int

	// MatchedText is the actual text that was matched and parsed
	MatchedText string

	// Confidence is a score (0.0 to 1.0) indicating parsing confidence
	Confidence float64
}

// DefaultSettings returns a Settings struct with sensible defaults.
func DefaultSettings() *Settings {
	return &Settings{
		DateOrder:         "MDY",
		Languages:         []string{"en"},
		RelativeBase:      time.Time{},
		EnableParsers:     []string{"timestamp", "relative", "absolute", "timezone", "time", "incomplete", "ordinal", "week"},
		Strict:            false,
		PreferredTimezone: time.UTC,
		PreferDatesFrom:   "future", // Default to forward-looking dates
	}
}

// ParseDate parses a date string and returns the corresponding time.Time value.
// If opts is nil, DefaultSettings() is used.
func ParseDate(input string, opts *Settings) (time.Time, error) {
	if input == "" {
		return time.Time{}, &ErrEmptyInput{}
	}

	if opts == nil {
		opts = DefaultSettings()
	}

	// Check if we should auto-detect date order
	autoDetect := opts.DateOrder == ""

	// Normalize settings
	settings := normalizeSettings(opts)

	// Load language translations
	langs := translations.GlobalRegistry.GetMultiple(settings.Languages)

	// Create parser context
	ctx := &parserContext{
		input:               input,
		settings:            settings,
		autoDetectDateOrder: autoDetect,
		languages:           langs,
	}

	// Try each enabled parser in order
	var parseErrors []error

	// 1. Try timestamp parser
	if isParserEnabled(settings, "timestamp") {
		if result, err := parseTimestamp(ctx); err == nil {
			return result, nil
		} else {
			// Check if it's a specific error type that should be returned as-is
			if isSpecificError(err) {
				return time.Time{}, err
			}
			parseErrors = append(parseErrors, err)
		}
	}

	// 2. Try absolute date parser
	if isParserEnabled(settings, "absolute") {
		if result, err := parseAbsolute(ctx); err == nil {
			return result, nil
		} else {
			// Check if it's a specific error type that should be returned as-is
			if isSpecificError(err) {
				return time.Time{}, err
			}
			parseErrors = append(parseErrors, err)
		}
	}

	// 3. Try relative date parser
	if isParserEnabled(settings, "relative") {
		if result, err := parseRelative(ctx); err == nil {
			return result, nil
		} else {
			// Check if it's a specific error type that should be returned as-is
			if isSpecificError(err) {
				return time.Time{}, err
			}
			parseErrors = append(parseErrors, err)
		}
	}

	// 4. Try time parser (v1.0 Phase 3B)
	if isParserEnabled(settings, "time") {
		if result, err := tryParseTime(ctx); err == nil {
			return result, nil
		} else {
			// Check if it's a specific error type that should be returned as-is
			if isSpecificError(err) {
				return time.Time{}, err
			}
			parseErrors = append(parseErrors, err)
		}
	}

	// 5. Try incomplete date parser (v1.1 Phase 4)
	if isParserEnabled(settings, "incomplete") {
		if result, err := tryParseIncompleteDate(ctx); err == nil {
			return result, nil
		} else {
			// Check if it's a specific error type that should be returned as-is
			if isSpecificError(err) {
				return time.Time{}, err
			}
			parseErrors = append(parseErrors, err)
		}
	}

	// 6. Try ordinal date parser (v1.1 Phase 4)
	if isParserEnabled(settings, "ordinal") {
		if result, err := tryParseOrdinalDate(ctx); err == nil {
			return result, nil
		} else {
			// Check if it's a specific error type that should be returned as-is
			if isSpecificError(err) {
				return time.Time{}, err
			}
			parseErrors = append(parseErrors, err)
		}
	}

	// 7. Try week number parser (v1.2 Phase 5)
	if isParserEnabled(settings, "week") {
		if result, err := tryParseWeekNumber(ctx); err == nil {
			return result, nil
		} else {
			// Check if it's a specific error type that should be returned as-is
			if isSpecificError(err) {
				return time.Time{}, err
			}
			parseErrors = append(parseErrors, err)
		}
	}

	// No parser succeeded - return helpful error
	if len(parseErrors) > 0 {
		return time.Time{}, newInvalidFormatError(input)
	}

	return time.Time{}, newInvalidFormatError(input)
}

// isSpecificError checks if an error is a specific typed error that should be preserved
func isSpecificError(err error) bool {
	// Check for our custom error types that should be returned as-is
	var (
		ambigErr   *ErrAmbiguousDate
		invalidErr *ErrInvalidDate
	)
	return err != nil && (errors.As(err, &ambigErr) || errors.As(err, &invalidErr))
}

// ExtractDates scans text and extracts all recognizable dates with their positions.
// If opts is nil, DefaultSettings() is used.
func ExtractDates(text string, opts *Settings) ([]ParsedDate, error) {
	if text == "" {
		return nil, &ErrEmptyInput{}
	}

	if opts == nil {
		opts = DefaultSettings()
	}

	// Check if we should auto-detect date order
	autoDetect := opts.DateOrder == ""

	settings := normalizeSettings(opts)

	// Load language translations
	langs := translations.GlobalRegistry.GetMultiple(settings.Languages)

	// Create parser context
	ctx := &parserContext{
		input:               text,
		settings:            settings,
		autoDetectDateOrder: autoDetect,
		languages:           langs,
	}

	return extractAllDates(ctx)
}

// parserContext holds the state during parsing operations.
type parserContext struct {
	input               string
	settings            *Settings
	autoDetectDateOrder bool                     // true if DateOrder should be auto-detected
	languages           []*translations.Language // loaded language translations
}

// normalizeSettings ensures settings have valid values.
func normalizeSettings(opts *Settings) *Settings {
	settings := &Settings{
		DateOrder:         opts.DateOrder,
		Languages:         opts.Languages,
		RelativeBase:      opts.RelativeBase,
		EnableParsers:     opts.EnableParsers,
		Strict:            opts.Strict,
		PreferredTimezone: opts.PreferredTimezone,
		PreferDatesFrom:   opts.PreferDatesFrom,
	}

	// Set defaults for empty values
	if settings.DateOrder == "" {
		settings.DateOrder = "MDY"
	}

	if len(settings.Languages) == 0 {
		settings.Languages = []string{"en"}
	}

	if settings.RelativeBase.IsZero() {
		settings.RelativeBase = time.Now()
	}

	if len(settings.EnableParsers) == 0 {
		settings.EnableParsers = []string{"timestamp", "relative", "absolute", "timezone", "time", "incomplete", "ordinal", "week"}
	}

	if settings.PreferredTimezone == nil {
		settings.PreferredTimezone = time.UTC
	}

	if settings.PreferDatesFrom == "" {
		settings.PreferDatesFrom = "future"
	}

	return settings
}

// isParserEnabled checks if a specific parser is enabled in settings.
func isParserEnabled(settings *Settings, parserName string) bool {
	for _, enabled := range settings.EnableParsers {
		if enabled == parserName {
			return true
		}
	}
	return false
}
