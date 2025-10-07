# godateparser

[![Go Version](https://img.shields.io/github/go-mod/go-version/coredds/godateparser)](https://go.dev/)
[![Release](https://img.shields.io/github/v/release/coredds/godateparser)](https://github.com/coredds/godateparser/releases)
[![License](https://img.shields.io/github/license/coredds/godateparser)](https://github.com/coredds/godateparser/blob/master/LICENSE)
[![CI](https://github.com/coredds/godateparser/workflows/CI/badge.svg)](https://github.com/coredds/godateparser/actions)
[![CodeQL](https://github.com/coredds/godateparser/workflows/CodeQL/badge.svg)](https://github.com/coredds/godateparser/security/code-scanning)
[![Codecov](https://codecov.io/gh/coredds/godateparser/branch/master/graph/badge.svg)](https://codecov.io/gh/coredds/godateparser)
[![Go Reference](https://pkg.go.dev/badge/github.com/coredds/godateparser.svg)](https://pkg.go.dev/github.com/coredds/godateparser)

A powerful Go library for parsing human-readable date strings in multiple formats and languages.

## Features

godateparser v1.3.1 supports **English**, **Spanish**, **Portuguese (Brazil)**, **French (France)**, **German (Germany)**, **Italian (Italy)**, **Dutch (Netherlands)**, **Russian (Russia)**, **Chinese Simplified (China)**, and **Japanese (Japan)** with comprehensive date parsing capabilities.

### Core Parsing
- **Absolute Dates**: ISO 8601, numeric formats (MDY/DMY/YMD), month names, two-digit years
- **Relative Dates**: Natural language (yesterday, 2 days ago, next Monday, last week)
- **Extended Relative**: Period boundaries, complex expressions, quarter support
- **Unix Timestamps**: Seconds and milliseconds with automatic detection
- **Time Parsing**: 12/24-hour formats, natural language (noon, midnight)
- **Date Ranges**: From/to patterns, duration ranges (next 7 days, last 2 weeks)

### Advanced Features
- **Multi-Language Support**: English (en), Spanish (es), Portuguese (pt), French (fr), German (de), Italian (it), Dutch (nl), Russian (ru), Chinese Simplified (zh), and Japanese (ja) with automatic detection
- **Timezone Support**: 30+ abbreviations, offsets, DST-aware via IANA database
- **Incomplete Dates**: Year-only, month-only, month+day without year
- **Ordinal Dates**: 1st, 2nd, 3rd, 21st with full/abbreviated month names
- **Week Numbers**: ISO 8601 week dates (2024-W15, Week 42 2024)
- **Natural Time**: Quarter/half past/to expressions (quarter past 3, half to 5)
- **PREFER_DATES_FROM**: Future/past disambiguation for ambiguous dates

### Quality & Performance
- **Comprehensive**: 1200+ test cases covering all scenarios (English + Spanish + Portuguese + French + German + Italian + Dutch + Russian + Chinese + Japanese)
- **Fast**: Sub-50μs parsing for most operations
- **Robust**: Custom error types with helpful suggestions
- **Flexible**: Extensive customization options
- **Production-Ready**: Validated, documented, and battle-tested

## Installation

```bash
go get github.com/coredds/godateparser
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/coredds/godateparser"
)

func main() {
    // Parse a simple date
    date, err := godateparser.ParseDate("December 31, 2024", nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(date) // 2024-12-31 00:00:00 +0000 UTC

    // Parse relative dates
    date, _ = godateparser.ParseDate("2 days ago", nil)
    fmt.Println(date)

    // Parse with custom settings
    settings := &godateparser.Settings{
        DateOrder:    "DMY",
        Languages:    []string{"en"},
        RelativeBase: time.Now(),
    }
    date, _ = godateparser.ParseDate("31/12/2024", settings)
    fmt.Println(date) // 2024-12-31 00:00:00 +0000 UTC
}
```

## Usage Examples

### Basic Parsing

```go
// ISO 8601 format
date, err := godateparser.ParseDate("2024-12-31", nil)

// US format (Month/Day/Year)
date, err = godateparser.ParseDate("12/31/2024", nil)

// European format (Day/Month/Year)
settings := &godateparser.Settings{DateOrder: "DMY"}
date, err = godateparser.ParseDate("31/12/2024", settings)

// Month names
date, err = godateparser.ParseDate("December 31, 2024", nil)
date, err = godateparser.ParseDate("31 Dec 2024", nil)
```

### Relative Dates

```go
import "time"

base := time.Now()
settings := &godateparser.Settings{RelativeBase: base}

// Simple relative dates
date, _ := godateparser.ParseDate("yesterday", settings)
date, _ = godateparser.ParseDate("today", settings)
date, _ = godateparser.ParseDate("tomorrow", settings)

// Time units
date, _ = godateparser.ParseDate("2 days ago", settings)
date, _ = godateparser.ParseDate("in 3 weeks", settings)
date, _ = godateparser.ParseDate("5 months ago", settings)

// Periods
date, _ = godateparser.ParseDate("last week", settings)
date, _ = godateparser.ParseDate("next month", settings)

// Weekdays
date, _ = godateparser.ParseDate("next Monday", settings)
date, _ = godateparser.ParseDate("last Friday", settings)
```

### Unix Timestamps

```go
// Seconds
date, err := godateparser.ParseDate("1609459200", nil)

// Milliseconds
date, err = godateparser.ParseDate("1609459200000", nil)
```

### Extracting Dates from Text

```go
text := "The project started on December 31, 2024 and will end on 2025-06-15. " +
        "Our next meeting is tomorrow at 3 PM."

results, err := godateparser.ExtractDates(text, nil)
if err != nil {
    panic(err)
}

for _, result := range results {
    fmt.Printf("Found: %s at position %d (confidence: %.2f)\n",
        result.MatchedText, result.Position, result.Confidence)
    fmt.Printf("Parsed as: %s\n", result.Date)
}
```

### Extended Relative Expressions (v1.0)

```go
// Period boundaries
godateparser.ParseDate("beginning of month", nil)     // First day of current month
godateparser.ParseDate("end of month", nil)           // Last day of current month
godateparser.ParseDate("beginning of next year", nil) // Jan 1 of next year
godateparser.ParseDate("end of last week", nil)       // Sunday of last week

// This/next/last disambiguation
godateparser.ParseDate("this Monday", nil)    // Next occurrence of Monday
godateparser.ParseDate("this month", nil)     // First day of current month
godateparser.ParseDate("this quarter", nil)   // First day of current quarter

// Complex relative expressions
godateparser.ParseDate("a week from Tuesday", nil)         // Next Tuesday + 7 days
godateparser.ParseDate("2 days from Monday", nil)          // Next Monday + 2 days
godateparser.ParseDate("3 days after tomorrow", nil)       // Tomorrow + 3 days
godateparser.ParseDate("2 weeks before last Monday", nil)  // Last Monday - 14 days

// Quarter support
godateparser.ParseDate("Q1", nil)           // January 1 of current year
godateparser.ParseDate("Q4 2024", nil)      // October 1, 2024
godateparser.ParseDate("next quarter", nil) // First day of next quarter
godateparser.ParseDate("last quarter", nil) // First day of last quarter
```

### Advanced Date Parsing Features

```go
// PREFER_DATES_FROM: Control temporal disambiguation
settings := &godateparser.Settings{
    PreferDatesFrom: "future",  // or "past"
}

// Standalone weekdays
godateparser.ParseDate("Monday", settings)     // Next Monday (with "future")
godateparser.ParseDate("Friday", settings)     // Next Friday

settings.PreferDatesFrom = "past"
godateparser.ParseDate("Monday", settings)     // Last Monday (with "past")

// Incomplete dates (year, month, or month+day without year)
godateparser.ParseDate("2024", nil)            // January 1, 2024
godateparser.ParseDate("May", nil)             // May 1 of current/next year
godateparser.ParseDate("June 15", nil)         // June 15 of current/next year
godateparser.ParseDate("15 June", nil)         // June 15 of current/next year

// Ordinal dates
godateparser.ParseDate("1st", nil)             // 1st of current/next month
godateparser.ParseDate("June 3rd", nil)        // June 3
godateparser.ParseDate("3rd June", nil)        // June 3
godateparser.ParseDate("3rd of June", nil)     // June 3
godateparser.ParseDate("June 3rd 2024", nil)   // June 3, 2024
godateparser.ParseDate("3rd of June 2024", nil) // June 3, 2024
godateparser.ParseDate("21st March", nil)      // March 21

// Additional relative terms
godateparser.ParseDate("a fortnight ago", nil)  // 14 days ago
godateparser.ParseDate("in a fortnight", nil)   // 14 days from now
godateparser.ParseDate("next fortnight", nil)   // 14 days from now
godateparser.ParseDate("a decade ago", nil)     // 10 years ago
godateparser.ParseDate("in a decade", nil)      // 10 years from now
godateparser.ParseDate("a quarter ago", nil)    // 3 months ago
```

### Week Numbers and Natural Time Expressions

```go
// Week number parsing (ISO 8601)
godateparser.ParseDate("2024-W15", nil)         // Week 15 of 2024 (Monday April 8)
godateparser.ParseDate("2024W15", nil)          // Without dash
godateparser.ParseDate("Week 15 2024", nil)     // Natural language
godateparser.ParseDate("2024 Week 15", nil)     // Alternate format
godateparser.ParseDate("W42", nil)              // Week 42 of current year
godateparser.ParseDate("2024-W15-3", nil)       // Wednesday of week 15

// Natural time expressions
godateparser.ParseDate("quarter past 3", nil)   // 3:15
godateparser.ParseDate("half past 9", nil)      // 9:30
godateparser.ParseDate("quarter to 5", nil)     // 4:45
godateparser.ParseDate("quarter past noon", nil)     // 12:15
godateparser.ParseDate("half past midnight", nil)    // 0:30
godateparser.ParseDate("quarter to midnight", nil)   // 23:45
```

### Custom Settings

```go
settings := &godateparser.Settings{
    // Date component order: "YMD", "MDY", or "DMY"
    DateOrder: "DMY",
    
    // Preferred languages (for future multi-language support)
    Languages: []string{"en"},
    
    // Base date for relative calculations
    RelativeBase: time.Date(2024, 10, 2, 12, 0, 0, 0, time.UTC),
    
    // Enable specific parsers
    EnableParsers: []string{"timestamp", "relative", "absolute"},
    
    // Strict mode (return error on ambiguous input)
    Strict: false,
    
    // Preferred timezone
    PreferredTimezone: time.UTC,
}

date, err := godateparser.ParseDate("31/12/2024", settings)
```

## API Reference

### ParseDate

```go
func ParseDate(input string, opts *Settings) (time.Time, error)
```

Parses a date string and returns the corresponding `time.Time` value. If `opts` is `nil`, `DefaultSettings()` is used.

### ExtractDates

```go
func ExtractDates(text string, opts *Settings) ([]ParsedDate, error)
```

Scans text and extracts all recognizable dates with their positions. Returns a slice of `ParsedDate` structs.

### Settings

```go
type Settings struct {
    DateOrder         string      // "YMD", "MDY", or "DMY"
    Languages         []string    // Preferred languages/locales
    RelativeBase      time.Time   // Base date for relative parsing
    EnableParsers     []string    // List of enabled parsers
    Strict            bool        // Strict mode for ambiguous input
    PreferredTimezone *time.Location // Default timezone
    PreferDatesFrom   string      // "future", "past", or "" (v1.1.0+)
}
```

### ParsedDate

```go
type ParsedDate struct {
    Date        time.Time // Parsed date/time value
    Position    int       // Start index in input text
    Length      int       // Length of matched substring
    MatchedText string    // The actual matched text
    Confidence  float64   // Confidence score (0.0 to 1.0)
}
```

## Supported Date Formats

### Absolute Dates
- ISO 8601: `2024-12-31`, `2024-12-31T10:30:00`
- Numeric: `12/31/2024`, `31-12-2024`
- Month names: `December 31, 2024`, `31 Dec 2024`

### Relative Dates
- Simple: `yesterday`, `today`, `tomorrow`
- Time units: `2 days ago`, `in 3 weeks`, `5 months ago`
- Extended units: `a fortnight ago`, `in a decade`, `a quarter ago`
- Periods: `last week`, `next month`, `last year`, `next fortnight`, `last decade`
- Weekdays: `next Monday`, `last Friday`, `Monday` (with PreferDatesFrom)

### Incomplete Dates (v1.1.0+)
- Year only: `2024`
- Month only: `May`, `December`
- Month + Day: `June 15`, `15 June`

### Ordinal Dates (v1.1.0+)
- Basic: `1st`, `23rd`, `31st`
- With month: `June 3rd`, `3rd June`, `3rd of June`
- Full date: `June 3rd 2024`, `3rd of June 2024`, `21st March`

### Week Numbers (v1.2.0+)
- ISO 8601: `2024-W15`, `2024W15`
- Natural language: `Week 15 2024`, `2024 Week 15`
- Week only: `W42` (current year)
- With weekday: `2024-W15-3` (Wednesday)

### Natural Time Expressions (v1.2.0+)
- Quarter/half past: `quarter past 3`, `half past 9`
- Quarter/half to: `quarter to 5`, `half to 12`
- With noon/midnight: `quarter past noon`, `half past midnight`

### Timestamps
- Unix seconds: `1609459200`
- Unix milliseconds: `1609459200000`

## Testing

Run the test suite:

```bash
go test -v
```

Run benchmarks:

```bash
go test -bench=. -benchmem
```

## Performance

godateparser is optimized for performance with efficient regex patterns and minimal allocations. Benchmark results on typical hardware:

```
BenchmarkParseDate_ISO-8          500000    2500 ns/op
BenchmarkParseDate_Relative-8     300000    4000 ns/op
BenchmarkExtractDates-8           100000   15000 ns/op
```

## Roadmap

- [ ] Add support for 200+ language locales
- [ ] Implement language autodetection
- [ ] Support non-Gregorian calendar systems (Hijri, Jalali)
- [ ] Add timezone abbreviation parsing
- [ ] Improve ambiguity detection in strict mode
- [ ] Add more relative date expressions
- [ ] Performance optimizations

## Multi-Language Support

godateparser supports multiple languages with automatic detection or explicit selection.

### Supported Languages

- **English (en)**: Full support for all features
- **Spanish (es)**: Full support for all features
- **Portuguese (pt)**: Full support for all features (Brazilian Portuguese)
- **French (fr)**: Full support for all features (France)
- **German (de)**: Full support for all features (Germany)
- **Italian (it)**: Full support for all features (Italy)
- **Dutch (nl)**: Full support for all features (Netherlands)
- **Russian (ru)**: Full support for all features with Cyrillic script and grammatical cases
- **Chinese Simplified (zh)**: Full support including YYYY年MM月DD日 format, relative patterns (3天前, 2周后), next/last (下周, 上月)
- **Japanese (ja)**: Full support including YYYY年MM月DD日 format, relative patterns (3日前, 2週後), next/last (来週, 先月)

### Language Selection

#### Automatic Detection (Default)

```go
// Automatically detects language from input
godateparser.ParseDate("December 31, 2024", nil)  // English
godateparser.ParseDate("31 diciembre 2024", nil)  // Spanish
godateparser.ParseDate("31 dezembro 2024", nil)   // Portuguese
godateparser.ParseDate("31 décembre 2024", nil)   // French
godateparser.ParseDate("31 Dezember 2024", nil)   // German
godateparser.ParseDate("31 dicembre 2024", nil)   // Italian
godateparser.ParseDate("31 december 2024", nil)   // Dutch
godateparser.ParseDate("31 декабря 2024", nil)    // Russian
godateparser.ParseDate("星期一", nil)               // Chinese (Monday)
godateparser.ParseDate("月曜日", nil)               // Japanese (Monday)
```

#### Explicit Language Selection

```go
// Parse only Spanish
settings := &godateparser.Settings{
    Languages: []string{"es"},
}
godateparser.ParseDate("31 diciembre 2024", settings)

// Enable multiple languages with priority
settings = &godateparser.Settings{
    Languages: []string{"es", "en"}, // Try Spanish first, then English
}
godateparser.ParseDate("31 diciembre 2024", settings)  // Spanish
godateparser.ParseDate("December 31, 2024", settings)  // English

// Enable all languages
settings = &godateparser.Settings{
    Languages: []string{"ja", "zh", "ru", "nl", "it", "de", "fr", "pt", "es", "en"},
}
```

### Example: Spanish Language Support

Spanish demonstrates the multi-language capabilities with full support for gender variations, accent-optional parsing, and natural expressions.

```go
settings := &godateparser.Settings{
    Languages: []string{"es"},
    RelativeBase: time.Now(),
}

// Absolute dates
godateparser.ParseDate("31 diciembre 2024", settings)
godateparser.ParseDate("15 de marzo de 2024", settings)

// Simple relative terms
godateparser.ParseDate("ayer", settings)      // yesterday
godateparser.ParseDate("hoy", settings)       // today
godateparser.ParseDate("mañana", settings)    // tomorrow

// Time-based patterns
godateparser.ParseDate("hace 2 días", settings)     // 2 days ago
godateparser.ParseDate("en 3 semanas", settings)    // in 3 weeks
godateparser.ParseDate("próximo viernes", settings) // next Friday

// Period boundaries
godateparser.ParseDate("inicio de mes", settings)   // beginning of month
godateparser.ParseDate("fin de año", settings)      // end of year

// Time expressions
godateparser.ParseDate("mediodía", settings)        // 12:00 PM
godateparser.ParseDate("3 y media", settings)       // 3:30
```

**📚 For comprehensive examples in all 10 languages, see [LANGUAGE_EXAMPLES.md](LANGUAGE_EXAMPLES.md)**

## Contributing

Contributions are welcome! Please feel free to submit issues, feature requests, or pull requests.

### Adding New Languages

Interested in adding support for Korean, Arabic, Polish, or other languages? Check out the `translations/` package for the translation infrastructure. Use the existing language implementations (Spanish, Portuguese, French, German, Italian, Dutch, Russian, Chinese, or Japanese) as a reference.

## License

MIT License - see LICENSE file for details.

## Acknowledgments

Special thanks to all contributors and users of this library.

## Contact

For questions or support, please open an issue on GitHub.


