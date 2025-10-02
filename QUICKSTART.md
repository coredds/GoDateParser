# GoDateParser - Quick Start Guide

## Installation

```bash
go get github.com/coredds/GoDateParser
```

## 5-Minute Tutorial

### 1. Parse a Simple Date

```go
package main

import (
    "fmt"
    "github.com/coredds/GoDateParser"
)

func main() {
    date, err := godateparser.ParseDate("December 31, 2024", nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(date) // 2024-12-31 00:00:00 +0000 UTC
}
```

### 2. Parse Relative Dates

```go
// Parse relative to current time
date, _ := godateparser.ParseDate("2 days ago", nil)
date, _ = godateparser.ParseDate("next Monday", nil)
date, _ = godateparser.ParseDate("in 3 weeks", nil)
```

### 3. Extract Dates from Text

```go
text := "The meeting is on December 31, 2024 and the deadline is 2025-06-15."
results, _ := godateparser.ExtractDates(text, nil)

for _, result := range results {
    fmt.Printf("Found '%s' at position %d\n", result.MatchedText, result.Position)
    fmt.Printf("Date: %s\n", result.Date)
}
```

### 4. Customize with Settings

```go
import "time"

settings := &godateparser.Settings{
    DateOrder:         "DMY",  // European format
    RelativeBase:      time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
    PreferredTimezone: time.UTC,
}

date, _ := godateparser.ParseDate("31/12/2024", settings)
```

### 5. Handle Multiple Format Types

```go
inputs := []string{
    "2024-12-31",           // ISO 8601
    "12/31/2024",           // US format
    "December 31, 2024",    // Month name
    "yesterday",            // Relative
    "1609459200",           // Unix timestamp
}

for _, input := range inputs {
    date, err := godateparser.ParseDate(input, nil)
    if err == nil {
        fmt.Printf("%s -> %s\n", input, date.Format("2006-01-02"))
    }
}
```

## Supported Formats

### Absolute Dates
- **ISO 8601**: `2024-12-31`, `2024-12-31T10:30:00`
- **Numeric**: `12/31/2024`, `31-12-2024`
- **Month Names**: `December 31, 2024`, `31 Dec 2024`

### Relative Dates
- **Simple**: `yesterday`, `today`, `tomorrow`
- **Units**: `2 days ago`, `in 3 weeks`
- **Periods**: `last week`, `next month`
- **Weekdays**: `next Monday`, `last Friday`

### Timestamps
- **Seconds**: `1609459200`
- **Milliseconds**: `1609459200000`

## Common Use Cases

### Web Scraping
```go
// Extract all dates from scraped content
html := scrapeWebsite()
dates, _ := godateparser.ExtractDates(html, nil)
```

### Log Parsing
```go
// Parse dates from log files
logLine := "[2024-12-31 10:30:00] Error occurred"
date, _ := godateparser.ParseDate("2024-12-31 10:30:00", nil)
```

### User Input Processing
```go
// Handle natural language from users
userInput := "remind me in 2 days"
reminderDate, _ := godateparser.ParseDate("in 2 days", nil)
```

### Data Normalization
```go
// Convert various formats to standard format
dates := []string{"12/31/2024", "31 Dec 2024", "2024-12-31"}
for _, d := range dates {
    parsed, _ := godateparser.ParseDate(d, nil)
    normalized := parsed.Format("2006-01-02")
    fmt.Println(normalized) // All output: 2024-12-31
}
```

## Error Handling

```go
date, err := godateparser.ParseDate("invalid date", nil)
if err != nil {
    fmt.Printf("Failed to parse: %v\n", err)
    // Handle error appropriately
}
```

## Performance Tips

1. **Reuse Settings**: Create a `Settings` object once and reuse it
2. **Enable Specific Parsers**: Disable unnecessary parsers for better performance
3. **Batch Processing**: Process multiple dates in a loop efficiently

```go
// Good: Reuse settings
settings := godateparser.DefaultSettings()
for _, input := range manyInputs {
    date, _ := godateparser.ParseDate(input, settings)
    // Process date...
}

// Better: Enable only needed parsers
settings.EnableParsers = []string{"absolute"} // Only absolute dates
```

## Running Tests

```bash
# Run all tests
go test

# Run with coverage
go test -cover

# Run benchmarks
go test -bench=. -run=XXX
```

## Running Examples

```bash
# Run the example application
go run examples/main.go

# Or build and run
go build -o example examples/main.go
./example
```

## Multi-Language Support

GoDateParser supports English and Spanish with automatic detection.

### Automatic Language Detection

```go
// English - automatically detected
date, _ := godateparser.ParseDate("December 31, 2024", nil)

// Spanish - automatically detected
date, _ = godateparser.ParseDate("31 diciembre 2024", nil)
```

### Explicit Language Selection

```go
// Spanish only
settings := &godateparser.Settings{
    Languages: []string{"es"},
}
date, _ := godateparser.ParseDate("hace 2 días", settings)        // 2 days ago
date, _ = godateparser.ParseDate("próximo lunes", settings)       // next Monday
date, _ = godateparser.ParseDate("inicio de mes", settings)       // beginning of month
date, _ = godateparser.ParseDate("3 de junio 2024", settings)     // June 3, 2024
date, _ = godateparser.ParseDate("mediodía", settings)            // noon

// Multiple languages (Spanish with English fallback)
settings = &godateparser.Settings{
    Languages: []string{"es", "en"},
}
date, _ = godateparser.ParseDate("próximo viernes", settings)     // next Friday (Spanish)
date, _ = godateparser.ParseDate("next Monday", settings)         // next Monday (English)
```

### Spanish Examples

```go
settings := &godateparser.Settings{
    Languages: []string{"es"},
    RelativeBase: time.Now(),
}

// Dates
godateparser.ParseDate("31 diciembre 2024", settings)
godateparser.ParseDate("15 de marzo de 2024", settings)

// Relative
godateparser.ParseDate("ayer", settings)                // yesterday
godateparser.ParseDate("mañana", settings)              // tomorrow
godateparser.ParseDate("hace 1 semana", settings)       // 1 week ago
godateparser.ParseDate("en 3 días", settings)           // in 3 days

// Extended
godateparser.ParseDate("este lunes", settings)          // this Monday
godateparser.ParseDate("próxima semana", settings)      // next week
godateparser.ParseDate("fin de mes", settings)          // end of month

// Time
godateparser.ParseDate("mediodía", settings)            // noon
godateparser.ParseDate("3 y cuarto", settings)          // 3:15
```

## Next Steps

- Read the full [README.md](README.md) for complete documentation and Spanish examples
- Check the `translations/` package to see how language support is implemented
- Review [CHANGELOG.md](CHANGELOG.md) for v1.3.0 multi-language features

## Getting Help

- Check the test files for more examples
- Review the source code comments
- Open an issue on GitHub for bugs or feature requests

## License

MIT License - see [LICENSE](LICENSE) file for details


