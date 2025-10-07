package godateparser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/coredds/godateparser/translations"
)

// Time parsing patterns for v1.0.0 Phase 3B
// Supports: 12-hour format, 24-hour format, with/without seconds

// timePattern represents a time-only parsing pattern
type timePattern struct {
	regex  *regexp.Regexp
	parser func(*parserContext, []string) (time.Time, error)
}

// Time patterns (ordered by specificity)
var timePatterns = []*timePattern{
	// 12-hour format with AM/PM
	// Examples: 3:30 PM, 9:15 AM, 11:45:30 PM
	{
		regex: regexp.MustCompile(`(?i)^(\d{1,2}):(\d{2})(?::(\d{2}))?\s*(AM|PM)$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			hour, _ := strconv.Atoi(matches[1])
			minute, _ := strconv.Atoi(matches[2])
			second := 0
			if matches[3] != "" {
				second, _ = strconv.Atoi(matches[3])
			}
			period := strings.ToUpper(matches[4])

			// Convert to 24-hour format
			if period == "PM" && hour != 12 {
				hour += 12
			} else if period == "AM" && hour == 12 {
				hour = 0
			}

			// Validate time components
			if err := validateTime(hour, minute, second); err != nil {
				return time.Time{}, err
			}

			// Use base date from settings
			base := ctx.settings.RelativeBase
			return time.Date(base.Year(), base.Month(), base.Day(), hour, minute, second, 0, base.Location()), nil
		},
	},
	// 12-hour format without colon seconds (9:15AM, 3:30PM)
	{
		regex: regexp.MustCompile(`(?i)^(\d{1,2}):(\d{2})(AM|PM)$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			hour, _ := strconv.Atoi(matches[1])
			minute, _ := strconv.Atoi(matches[2])
			period := strings.ToUpper(matches[3])

			// Convert to 24-hour format
			if period == "PM" && hour != 12 {
				hour += 12
			} else if period == "AM" && hour == 12 {
				hour = 0
			}

			// Validate time components
			if err := validateTime(hour, minute, 0); err != nil {
				return time.Time{}, err
			}

			// Use base date from settings
			base := ctx.settings.RelativeBase
			return time.Date(base.Year(), base.Month(), base.Day(), hour, minute, 0, 0, base.Location()), nil
		},
	},
	// Short format (9am, 3pm, 12pm)
	{
		regex: regexp.MustCompile(`(?i)^(\d{1,2})(am|pm)$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			hour, _ := strconv.Atoi(matches[1])
			period := strings.ToLower(matches[2])

			// Convert to 24-hour format
			if period == "pm" && hour != 12 {
				hour += 12
			} else if period == "am" && hour == 12 {
				hour = 0
			}

			// Validate time components
			if err := validateTime(hour, 0, 0); err != nil {
				return time.Time{}, err
			}

			// Use base date from settings
			base := ctx.settings.RelativeBase
			return time.Date(base.Year(), base.Month(), base.Day(), hour, 0, 0, 0, base.Location()), nil
		},
	},
	// 24-hour format with seconds (14:30:00, 09:15:45)
	{
		regex: regexp.MustCompile(`^(\d{1,2}):(\d{2}):(\d{2})$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			hour, _ := strconv.Atoi(matches[1])
			minute, _ := strconv.Atoi(matches[2])
			second, _ := strconv.Atoi(matches[3])

			// Validate time components
			if err := validateTime(hour, minute, second); err != nil {
				return time.Time{}, err
			}

			// Use base date from settings
			base := ctx.settings.RelativeBase
			return time.Date(base.Year(), base.Month(), base.Day(), hour, minute, second, 0, base.Location()), nil
		},
	},
	// 24-hour format without seconds (14:30, 09:15, 23:59)
	{
		regex: regexp.MustCompile(`^(\d{1,2}):(\d{2})$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			hour, _ := strconv.Atoi(matches[1])
			minute, _ := strconv.Atoi(matches[2])

			// Validate time components
			if err := validateTime(hour, minute, 0); err != nil {
				return time.Time{}, err
			}

			// Use base date from settings
			base := ctx.settings.RelativeBase
			return time.Date(base.Year(), base.Month(), base.Day(), hour, minute, 0, 0, base.Location()), nil
		},
	},
	// Natural language time expressions (v1.2 Phase 5)
	// "quarter past 3", "half past 9", "quarter to 5"
	{
		regex: regexp.MustCompile(`(?i)^(quarter|half)\s+(past|to|before|after)\s+(\d{1,2})$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			fraction := strings.ToLower(matches[1])
			direction := strings.ToLower(matches[2])
			hour, _ := strconv.Atoi(matches[3])

			var minute int
			if fraction == "quarter" {
				minute = 15
			} else { // half
				minute = 30
			}

			// Adjust for "to" or "before"
			if direction == "to" || direction == "before" {
				minute = 60 - minute
				hour--
				if hour < 0 {
					hour = 23
				}
			}

			// Validate and use 24-hour format
			if hour > 23 {
				return time.Time{}, &ErrInvalidDate{
					Year:   0,
					Month:  0,
					Day:    0,
					Reason: fmt.Sprintf("hour %d out of range (0-23)", hour),
				}
			}

			base := ctx.settings.RelativeBase
			return time.Date(base.Year(), base.Month(), base.Day(), hour, minute, 0, 0, base.Location()), nil
		},
	},
	// "quarter past noon", "half past midnight", "quarter to noon"
	{
		regex: regexp.MustCompile(`(?i)^(quarter|half)\s+(past|to|before|after)\s+(noon|midnight)$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			fraction := strings.ToLower(matches[1])
			direction := strings.ToLower(matches[2])
			reference := strings.ToLower(matches[3])

			var hour, minute int
			if reference == "noon" {
				hour = 12
			} else { // midnight
				hour = 0
			}

			if fraction == "quarter" {
				minute = 15
			} else { // half
				minute = 30
			}

			// Adjust for "to" or "before"
			if direction == "to" || direction == "before" {
				minute = 60 - minute
				hour--
				if hour < 0 {
					hour = 23
				}
			}

			base := ctx.settings.RelativeBase
			return time.Date(base.Year(), base.Month(), base.Day(), hour, minute, 0, 0, base.Location()), nil
		},
	},
	// Basic "noon" and "midnight"
	{
		regex: regexp.MustCompile(`(?i)^(noon|midnight)$`),
		parser: func(ctx *parserContext, matches []string) (time.Time, error) {
			word := strings.ToLower(matches[1])
			base := ctx.settings.RelativeBase

			switch word {
			case "noon":
				return time.Date(base.Year(), base.Month(), base.Day(), 12, 0, 0, 0, base.Location()), nil
			case "midnight":
				return time.Date(base.Year(), base.Month(), base.Day(), 0, 0, 0, 0, base.Location()), nil
			}

			return time.Time{}, fmt.Errorf("unrecognized time word: %s", word)
		},
	},
}

// validateTime validates time components
func validateTime(hour, minute, second int) error {
	if hour < 0 || hour > 23 {
		return &ErrInvalidDate{Year: 0, Month: 0, Day: 0, Reason: fmt.Sprintf("hour %d out of range (0-23)", hour)}
	}
	if minute < 0 || minute > 59 {
		return &ErrInvalidDate{Year: 0, Month: 0, Day: 0, Reason: fmt.Sprintf("minute %d out of range (0-59)", minute)}
	}
	if second < 0 || second > 59 {
		return &ErrInvalidDate{Year: 0, Month: 0, Day: 0, Reason: fmt.Sprintf("second %d out of range (0-59)", second)}
	}
	return nil
}

// tryParseTime attempts to parse time-only inputs
func tryParseTime(ctx *parserContext) (time.Time, error) {
	input := strings.TrimSpace(ctx.input)

	// Try multi-language time expressions first
	if result, err := tryParseMultiLangTime(ctx, input); err == nil {
		return result, nil
	}

	for _, pattern := range timePatterns {
		matches := pattern.regex.FindStringSubmatch(input)
		if matches != nil {
			return pattern.parser(ctx, matches)
		}
	}

	return time.Time{}, fmt.Errorf("no time pattern matched")
}

// ParseTime is a convenience function for parsing time-only strings with a base date
func ParseTime(timeStr string, baseDate time.Time) (time.Time, error) {
	settings := &Settings{
		RelativeBase: baseDate,
	}

	ctx := &parserContext{
		input:    timeStr,
		settings: settings,
	}

	return tryParseTime(ctx)
}

// tryParseMultiLangTime attempts to parse time expressions in multiple languages.
func tryParseMultiLangTime(ctx *parserContext, input string) (time.Time, error) {
	input = strings.ToLower(strings.TrimSpace(input))
	base := ctx.settings.RelativeBase

	// Try each language's time terms
	for _, lang := range ctx.languages {
		if lang.TimeTerms == nil {
			continue
		}

		// Try special terms: noon, midnight
		for _, noonTerm := range lang.TimeTerms.Noon {
			if strings.EqualFold(input, noonTerm) {
				return time.Date(base.Year(), base.Month(), base.Day(), 12, 0, 0, 0, base.Location()), nil
			}
		}

		for _, midnightTerm := range lang.TimeTerms.Midnight {
			if strings.EqualFold(input, midnightTerm) {
				return time.Date(base.Year(), base.Month(), base.Day(), 0, 0, 0, 0, base.Location()), nil
			}
		}

		// Try "X y cuarto" (X and quarter) patterns - Spanish "quarter past"
		if result, err := tryParseYCuarto(ctx, input, lang); err == nil {
			return result, nil
		}

		// Try "X y media" (X and half) patterns - Spanish "half past"
		if result, err := tryParseYMedia(ctx, input, lang); err == nil {
			return result, nil
		}

		// Try "menos cuarto las X" patterns - Spanish "quarter to"
		if result, err := tryParseMenosCuarto(ctx, input, lang); err == nil {
			return result, nil
		}

		// Try "quinze para as X" patterns - Portuguese "quarter to" (reversed order)
		if result, err := tryParseQuarterTo(ctx, input, lang); err == nil {
			return result, nil
		}

		// Try "15h30" or "15h" patterns - French time format
		if result, err := tryParseFrenchHFormat(ctx, input, lang); err == nil {
			return result, nil
		}

		// Try "3 heures 30" patterns - French "hours minutes"
		if result, err := tryParseFrenchHeures(ctx, input, lang); err == nil {
			return result, nil
		}

		// Try "kwart over 3" patterns - Dutch "quarter past"
		if result, err := tryParseDutchKwartOver(ctx, input, lang); err == nil {
			return result, nil
		}

		// Try "half 4" patterns - Dutch "half past" (special: means 3:30, not 4:30)
		if result, err := tryParseDutchHalf(ctx, input, lang); err == nil {
			return result, nil
		}

		// Try "3 e un quarto" patterns - Italian "quarter past"
		if result, err := tryParseItalianQuarto(ctx, input, lang); err == nil {
			return result, nil
		}

		// Try "meno un quarto le 3" patterns - Italian "quarter to"
		if result, err := tryParseItalianMenoQuarto(ctx, input, lang); err == nil {
			return result, nil
		}

		// Try "3 часа дня" patterns - Russian "3 hours of day/night"
		if result, err := tryParseRussianHoursAMPM(ctx, input, lang); err == nil {
			return result, nil
		}
	}

	return time.Time{}, fmt.Errorf("no multi-language time pattern matched")
}

// tryParseYCuarto parses Spanish "X y cuarto" (quarter past X)
func tryParseYCuarto(ctx *parserContext, input string, lang *translations.Language) (time.Time, error) {
	if lang.TimeTerms == nil {
		return time.Time{}, fmt.Errorf("no time terms")
	}

	// Try with each "past" term
	for _, pastTerm := range lang.TimeTerms.Past {
		if pastTerm == "" {
			continue
		}

		// Try with each "quarter" term
		for _, quarterTerm := range lang.TimeTerms.Quarter {
			if quarterTerm == "" {
				continue
			}

			// Pattern: "3 y cuarto" or similar
			pattern := fmt.Sprintf(`^(\d{1,2})\s+%s\s+%s$`,
				regexp.QuoteMeta(strings.ToLower(pastTerm)),
				regexp.QuoteMeta(strings.ToLower(quarterTerm)))
			re := regexp.MustCompile(pattern)

			if matches := re.FindStringSubmatch(input); matches != nil {
				hour, _ := strconv.Atoi(matches[1])
				minute := 15

				if hour > 23 {
					return time.Time{}, &ErrInvalidDate{
						Year:   0,
						Month:  0,
						Day:    0,
						Reason: fmt.Sprintf("hour %d out of range (0-23)", hour),
					}
				}

				base := ctx.settings.RelativeBase
				return time.Date(base.Year(), base.Month(), base.Day(), hour, minute, 0, 0, base.Location()), nil
			}
		}
	}

	return time.Time{}, fmt.Errorf("no match")
}

// tryParseYMedia parses Spanish "X y media" (half past X)
func tryParseYMedia(ctx *parserContext, input string, lang *translations.Language) (time.Time, error) {
	if lang.TimeTerms == nil {
		return time.Time{}, fmt.Errorf("no time terms")
	}

	// Try with each "past" term
	for _, pastTerm := range lang.TimeTerms.Past {
		if pastTerm == "" {
			continue
		}

		// Try with each "half" term
		for _, halfTerm := range lang.TimeTerms.Half {
			if halfTerm == "" {
				continue
			}

			// Pattern: "3 y media" or similar
			pattern := fmt.Sprintf(`^(\d{1,2})\s+%s\s+%s$`,
				regexp.QuoteMeta(strings.ToLower(pastTerm)),
				regexp.QuoteMeta(strings.ToLower(halfTerm)))
			re := regexp.MustCompile(pattern)

			if matches := re.FindStringSubmatch(input); matches != nil {
				hour, _ := strconv.Atoi(matches[1])
				minute := 30

				if hour > 23 {
					return time.Time{}, &ErrInvalidDate{
						Year:   0,
						Month:  0,
						Day:    0,
						Reason: fmt.Sprintf("hour %d out of range (0-23)", hour),
					}
				}

				base := ctx.settings.RelativeBase
				return time.Date(base.Year(), base.Month(), base.Day(), hour, minute, 0, 0, base.Location()), nil
			}
		}
	}

	return time.Time{}, fmt.Errorf("no match")
}

// tryParseMenosCuarto parses Spanish "menos cuarto las X" (quarter to X)
func tryParseMenosCuarto(ctx *parserContext, input string, lang *translations.Language) (time.Time, error) {
	if lang.TimeTerms == nil {
		return time.Time{}, fmt.Errorf("no time terms")
	}

	// Try with each "to" term
	for _, toTerm := range lang.TimeTerms.To {
		if toTerm == "" {
			continue
		}

		// Try with each "quarter" term
		for _, quarterTerm := range lang.TimeTerms.Quarter {
			if quarterTerm == "" {
				continue
			}

			// Pattern: "menos cuarto las 3" or "quinze para as 3" - with optional article
			pattern := fmt.Sprintf(`^%s\s+%s\s+(?:las\s+|as\s+)?(\d{1,2})$`,
				regexp.QuoteMeta(strings.ToLower(toTerm)),
				regexp.QuoteMeta(strings.ToLower(quarterTerm)))
			re := regexp.MustCompile(pattern)

			if matches := re.FindStringSubmatch(input); matches != nil {
				hour, _ := strconv.Atoi(matches[1])
				minute := 45 // quarter to = 45 minutes of previous hour
				hour--       // Go back one hour

				if hour < 0 {
					hour = 23
				}

				if hour > 23 {
					return time.Time{}, &ErrInvalidDate{
						Year:   0,
						Month:  0,
						Day:    0,
						Reason: fmt.Sprintf("hour %d out of range (0-23)", hour),
					}
				}

				base := ctx.settings.RelativeBase
				return time.Date(base.Year(), base.Month(), base.Day(), hour, minute, 0, 0, base.Location()), nil
			}
		}
	}

	return time.Time{}, fmt.Errorf("no match")
}

// tryParseFrenchHFormat parses French "15h30" or "15h" format
func tryParseFrenchHFormat(ctx *parserContext, input string, _ *translations.Language) (time.Time, error) {
	// Pattern: "15h30" or "15h" (h is the separator)
	pattern := `^(\d{1,2})h(\d{2})?$`
	re := regexp.MustCompile(pattern)

	if matches := re.FindStringSubmatch(input); matches != nil {
		hour, _ := strconv.Atoi(matches[1])
		minute := 0
		if matches[2] != "" {
			minute, _ = strconv.Atoi(matches[2])
		}

		// Validate time components
		if err := validateTime(hour, minute, 0); err != nil {
			return time.Time{}, err
		}

		base := ctx.settings.RelativeBase
		return time.Date(base.Year(), base.Month(), base.Day(), hour, minute, 0, 0, base.Location()), nil
	}

	return time.Time{}, fmt.Errorf("no match")
}

// tryParseFrenchHeures parses French "3 heures 30" format
func tryParseFrenchHeures(ctx *parserContext, input string, lang *translations.Language) (time.Time, error) {
	if lang.TimeTerms == nil {
		return time.Time{}, fmt.Errorf("no time terms")
	}

	// Try with "heure" or "heures" from OClock terms
	for _, oclockTerm := range lang.TimeTerms.OClock {
		if oclockTerm == "" {
			continue
		}

		// Pattern: "3 heures 30" or "3 heure 30"
		pattern := fmt.Sprintf(`^(\d{1,2})\s+%s\s+(\d{1,2})$`,
			regexp.QuoteMeta(strings.ToLower(oclockTerm)))
		re := regexp.MustCompile(pattern)

		if matches := re.FindStringSubmatch(input); matches != nil {
			hour, _ := strconv.Atoi(matches[1])
			minute, _ := strconv.Atoi(matches[2])

			// Validate time components
			if err := validateTime(hour, minute, 0); err != nil {
				return time.Time{}, err
			}

			base := ctx.settings.RelativeBase
			return time.Date(base.Year(), base.Month(), base.Day(), hour, minute, 0, 0, base.Location()), nil
		}
	}

	return time.Time{}, fmt.Errorf("no match")
}

// tryParseQuarterTo parses Portuguese "quinze para as 3" (quarter to 3) - quarter FIRST
func tryParseQuarterTo(ctx *parserContext, input string, lang *translations.Language) (time.Time, error) {
	if lang.TimeTerms == nil {
		return time.Time{}, fmt.Errorf("no time terms")
	}

	// Try with each "quarter" term
	for _, quarterTerm := range lang.TimeTerms.Quarter {
		if quarterTerm == "" {
			continue
		}

		// Try with each "to" term
		for _, toTerm := range lang.TimeTerms.To {
			if toTerm == "" {
				continue
			}

			// Pattern: "quinze para as 3" - QUARTER first, TO second, HOUR last
			pattern := fmt.Sprintf(`^%s\s+%s\s+(?:as\s+|o\s+)?(\d{1,2})$`,
				regexp.QuoteMeta(strings.ToLower(quarterTerm)),
				regexp.QuoteMeta(strings.ToLower(toTerm)))
			re := regexp.MustCompile(pattern)

			if matches := re.FindStringSubmatch(input); matches != nil {
				hour, _ := strconv.Atoi(matches[1])
				minute := 45 // quarter to = 45 minutes of previous hour
				hour--       // Go back one hour

				if hour < 0 {
					hour = 23
				}

				if hour > 23 {
					return time.Time{}, &ErrInvalidDate{
						Year:   0,
						Month:  0,
						Day:    0,
						Reason: fmt.Sprintf("hour %d out of range (0-23)", hour),
					}
				}

				base := ctx.settings.RelativeBase
				return time.Date(base.Year(), base.Month(), base.Day(), hour, minute, 0, 0, base.Location()), nil
			}
		}
	}

	return time.Time{}, fmt.Errorf("no match")
}

// tryParseDutchKwartOver parses Dutch "kwart over 3" (quarter past 3)
func tryParseDutchKwartOver(ctx *parserContext, input string, lang *translations.Language) (time.Time, error) {
	if lang.TimeTerms == nil {
		return time.Time{}, fmt.Errorf("no time terms")
	}

	// Try with each "quarter" term and "past" term
	for _, quarterTerm := range lang.TimeTerms.Quarter {
		if quarterTerm == "" {
			continue
		}
		for _, pastTerm := range lang.TimeTerms.Past {
			if pastTerm == "" {
				continue
			}

			// Pattern: "kwart over 3" - quarter + past + hour
			pattern := fmt.Sprintf(`^%s\s+%s\s+(\d{1,2})$`,
				regexp.QuoteMeta(strings.ToLower(quarterTerm)),
				regexp.QuoteMeta(strings.ToLower(pastTerm)))
			re := regexp.MustCompile(pattern)

			if matches := re.FindStringSubmatch(input); matches != nil {
				hour, _ := strconv.Atoi(matches[1])
				minute := 15

				if hour > 23 {
					return time.Time{}, &ErrInvalidDate{
						Year:   0,
						Month:  0,
						Day:    0,
						Reason: fmt.Sprintf("hour %d out of range (0-23)", hour),
					}
				}

				base := ctx.settings.RelativeBase
				return time.Date(base.Year(), base.Month(), base.Day(), hour, minute, 0, 0, base.Location()), nil
			}
		}
	}

	return time.Time{}, fmt.Errorf("no match")
}

// tryParseDutchHalf parses Dutch "half 4" (means 3:30, not 4:30!)
func tryParseDutchHalf(ctx *parserContext, input string, lang *translations.Language) (time.Time, error) {
	if lang.TimeTerms == nil {
		return time.Time{}, fmt.Errorf("no time terms")
	}

	// Try with each "half" term
	for _, halfTerm := range lang.TimeTerms.Half {
		if halfTerm == "" {
			continue
		}

		// Pattern: "half 4" - half + hour
		pattern := fmt.Sprintf(`^%s\s+(\d{1,2})$`,
			regexp.QuoteMeta(strings.ToLower(halfTerm)))
		re := regexp.MustCompile(pattern)

		if matches := re.FindStringSubmatch(input); matches != nil {
			hour, _ := strconv.Atoi(matches[1])
			// Dutch "half 4" means "half to 4" = 3:30
			hour--
			minute := 30

			if hour < 0 {
				hour = 23
			}

			if hour > 23 {
				return time.Time{}, &ErrInvalidDate{
					Year:   0,
					Month:  0,
					Day:    0,
					Reason: fmt.Sprintf("hour %d out of range (0-23)", hour),
				}
			}

			base := ctx.settings.RelativeBase
			return time.Date(base.Year(), base.Month(), base.Day(), hour, minute, 0, 0, base.Location()), nil
		}
	}

	return time.Time{}, fmt.Errorf("no match")
}

// tryParseItalianQuarto parses Italian "3 e un quarto" (3 and a quarter = 3:15)
func tryParseItalianQuarto(ctx *parserContext, input string, lang *translations.Language) (time.Time, error) {
	if lang.TimeTerms == nil {
		return time.Time{}, fmt.Errorf("no time terms")
	}

	// Try with each "quarter" term
	for _, quarterTerm := range lang.TimeTerms.Quarter {
		if quarterTerm == "" {
			continue
		}

		// Pattern: "3 e un quarto" - hour + e + un + quarter
		pattern := fmt.Sprintf(`^(\d{1,2})\s+e\s+un\s+%s$`,
			regexp.QuoteMeta(strings.ToLower(quarterTerm)))
		re := regexp.MustCompile(pattern)

		if matches := re.FindStringSubmatch(input); matches != nil {
			hour, _ := strconv.Atoi(matches[1])
			minute := 15

			if hour > 23 {
				return time.Time{}, &ErrInvalidDate{
					Year:   0,
					Month:  0,
					Day:    0,
					Reason: fmt.Sprintf("hour %d out of range (0-23)", hour),
				}
			}

			base := ctx.settings.RelativeBase
			return time.Date(base.Year(), base.Month(), base.Day(), hour, minute, 0, 0, base.Location()), nil
		}
	}

	return time.Time{}, fmt.Errorf("no match")
}

// tryParseItalianMenoQuarto parses Italian "meno un quarto le 3" (less a quarter to 3 = 2:45)
func tryParseItalianMenoQuarto(ctx *parserContext, input string, lang *translations.Language) (time.Time, error) {
	if lang.TimeTerms == nil {
		return time.Time{}, fmt.Errorf("no time terms")
	}

	// Try with each "quarter" term
	for _, quarterTerm := range lang.TimeTerms.Quarter {
		if quarterTerm == "" {
			continue
		}

		// Pattern: "meno un quarto le 3" - meno + un + quarter + le + hour
		pattern := fmt.Sprintf(`^meno\s+un\s+%s\s+le\s+(\d{1,2})$`,
			regexp.QuoteMeta(strings.ToLower(quarterTerm)))
		re := regexp.MustCompile(pattern)

		if matches := re.FindStringSubmatch(input); matches != nil {
			hour, _ := strconv.Atoi(matches[1])
			minute := 45
			hour-- // Go back one hour

			if hour < 0 {
				hour = 23
			}

			if hour > 23 {
				return time.Time{}, &ErrInvalidDate{
					Year:   0,
					Month:  0,
					Day:    0,
					Reason: fmt.Sprintf("hour %d out of range (0-23)", hour),
				}
			}

			base := ctx.settings.RelativeBase
			return time.Date(base.Year(), base.Month(), base.Day(), hour, minute, 0, 0, base.Location()), nil
		}
	}

	return time.Time{}, fmt.Errorf("no match")
}

// tryParseRussianHoursAMPM parses Russian "3 часа дня" (3 hours of day = 3 PM)
func tryParseRussianHoursAMPM(ctx *parserContext, input string, lang *translations.Language) (time.Time, error) {
	if lang.TimeTerms == nil {
		return time.Time{}, fmt.Errorf("no time terms")
	}

	// Try with each "o'clock" term (часов, час, часа)
	for _, oclockTerm := range lang.TimeTerms.OClock {
		if oclockTerm == "" {
			continue
		}

		// Try with AM terms (утра, ночи)
		for _, amTerm := range lang.TimeTerms.AM {
			if amTerm == "" {
				continue
			}

			// Pattern: "3 часа утра" or "9 часов утра"
			pattern := fmt.Sprintf(`^(\d{1,2})\s+%s\s+%s$`,
				regexp.QuoteMeta(strings.ToLower(oclockTerm)),
				regexp.QuoteMeta(strings.ToLower(amTerm)))
			re := regexp.MustCompile(pattern)

			if matches := re.FindStringSubmatch(input); matches != nil {
				hour, _ := strconv.Atoi(matches[1])

				// "ночи" (night) is for 12 AM - 5 AM, "утра" (morning) is for 6 AM - 11 AM
				// If hour is 12 with "ночи", it means midnight (0:00)
				if amTerm == "ночи" && hour == 12 {
					hour = 0
				}

				if hour > 23 {
					return time.Time{}, &ErrInvalidDate{
						Year:   0,
						Month:  0,
						Day:    0,
						Reason: fmt.Sprintf("hour %d out of range (0-23)", hour),
					}
				}

				base := ctx.settings.RelativeBase
				return time.Date(base.Year(), base.Month(), base.Day(), hour, 0, 0, 0, base.Location()), nil
			}
		}

		// Try with PM terms (дня, вечера)
		for _, pmTerm := range lang.TimeTerms.PM {
			if pmTerm == "" {
				continue
			}

			// Pattern: "3 часа дня" or "7 часов вечера"
			pattern := fmt.Sprintf(`^(\d{1,2})\s+%s\s+%s$`,
				regexp.QuoteMeta(strings.ToLower(oclockTerm)),
				regexp.QuoteMeta(strings.ToLower(pmTerm)))
			re := regexp.MustCompile(pattern)

			if matches := re.FindStringSubmatch(input); matches != nil {
				hour, _ := strconv.Atoi(matches[1])

				// "дня" (day) is for 12 PM - 5 PM, "вечера" (evening) is for 6 PM - 11 PM
				// If hour is less than 12, add 12 for PM
				if hour < 12 {
					hour += 12
				}

				if hour > 23 {
					return time.Time{}, &ErrInvalidDate{
						Year:   0,
						Month:  0,
						Day:    0,
						Reason: fmt.Sprintf("hour %d out of range (0-23)", hour),
					}
				}

				base := ctx.settings.RelativeBase
				return time.Date(base.Year(), base.Month(), base.Day(), hour, 0, 0, 0, base.Location()), nil
			}
		}
	}

	return time.Time{}, fmt.Errorf("no match")
}
