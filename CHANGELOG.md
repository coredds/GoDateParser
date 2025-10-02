# Changelog

All notable changes to godateparser will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.3.3] - 2025-10-02

### Added
- **CJK Weekday Modifier Patterns**: Complex patterns combining next/last + week + specific weekday
  - Japanese: `来週月曜` (next week Monday), `先週金曜` (last week Friday), `来週月曜日` (with 日 suffix)
  - Chinese: `下周一` (next week Monday), `上周五` (last week Friday), `下星期一` (using 星期)
  - Supports both week and month modifiers with specific weekdays
  - Intelligent greedy longest-match strategy for correct tokenization
  - Handles ambiguous characters (e.g., 月 = month OR part of 月曜 Monday)

- **Enhanced Chinese Language Support**:
  - Short-form weekday names for compound expressions: `一`, `二`, `三`, `四`, `五`, `六`, `日`, `天`
  - These enable proper parsing of patterns like `下周一` where `一` alone means Monday

- **New Parser Function**:
  - `tryCJKWeekdayModifier()` - Specialized handler for CJK weekday modifiers
  - Sorts weekdays by length (longest first) to ensure correct matching
  - Calculates target dates by finding start of week and adding appropriate offsets

### Changed
- **Test Suite**:
  - 12 new test cases for CJK weekday modifiers (6 Japanese + 6 Chinese)
  - All 950+ existing tests still passing
  - `parser_relative.go` - Enhanced to support CJK-specific patterns

### Fixed
- Week boundary interpretation now correctly calculates "last week Monday" as Monday of the previous calendar week

### Summary
- CJK languages (Chinese and Japanese) now have full feature parity with Romance languages for complex weekday modifier patterns

## [1.3.2] - 2025-10-02

### Added
- **Chinese Simplified Language Support (zh-Hans)**: Comprehensive Simplified Chinese date parsing
  - Months: `一月`, `二月`, ..., `十二月` (with numeric variants: `1月`, `2月`, etc.)
  - Weekdays: `星期一`, `周一`, `礼拜一` (three variations for each weekday)
  - Simple relative: `昨天` (yesterday), `今天` (today), `明天` (tomorrow)
  - Ago patterns: `1天前` (1 day ago), `2周前` (2 weeks ago) - no-space CJK format
  - Future patterns: `1天后` (in 1 day), `2周后` (in 2 weeks) - no-space CJK format
  - Next/last: `下周` (next week), `上月` (last month)
  - CJK date format: `2024年12月31日` (YYYY年MM月DD日)
  - Time expressions: `中午` (noon), `午夜` (midnight)
  - Alternative weekday forms: `礼拜一`, `星期天`, `周天`

- **Japanese Language Support (ja-JP)**: Comprehensive Japanese date parsing
  - Months: `一月`, `二月`, ..., `十二月` (with numeric variants: `1月`, `2月`, etc.)
  - Weekdays: `月曜日`, `火曜日`, ..., `日曜日` (with short forms: `月曜`, `火曜`, etc.)
  - Simple relative: `昨日` (yesterday), `今日` (today), `明日` (tomorrow)
  - Ago patterns: `1日前` (1 day ago), `2週前` (2 weeks ago), `1ヶ月前` (1 month ago) - no-space format
  - Future patterns: `1日後` (in 1 day), `2週後` (in 2 weeks) - no-space format
  - Next/last: `来週` (next week), `先月` (last month)
  - CJK date format: `2024年12月31日` (YYYY年MM月DD日 - same as Chinese)
  - Time expressions: `正午` (noon), `真夜中` (midnight)

- **CJK-Specific Parser Enhancements**:
  - `parseCJKDate()` function for `YYYY年MM月DD日` format (shared by Chinese and Japanese)
  - Enhanced `tryParseAgoSuffixPattern()` to support no-space CJK patterns like `3日前`
  - Enhanced `tryParseInPattern()` to support no-space CJK patterns like `3日後`
  - Enhanced `tryParseNextPattern()` to support no-space CJK patterns like `来週`
  - Enhanced `tryParseLastPattern()` to support no-space CJK patterns like `先週`
  - Pattern matching adjusted for languages without spaces between units and modifiers

- **New Files**:
  - `translations/chinese.go` - Chinese Simplified language implementation
  - `translations/japanese.go` - Japanese language implementation
  - `chinese_test.go` - Comprehensive Chinese test suite
  - `japanese_test.go` - Comprehensive Japanese test suite
  - `examples/chinese_demo.go` - Chinese parsing demonstration
  - `examples/japanese_demo.go` - Japanese parsing demonstration

### Changed
- **Test Suite**:
  - 950+ test cases total (up from 600+)
  - Support for 7 languages: English, Spanish, Portuguese, French, German, Chinese, Japanese
  - `translations/registry.go` - Added Chinese and Japanese to global registry
  - `parser_absolute.go` - Added CJK date format pattern
  - `parser_relative.go` - Enhanced for CJK no-space relative patterns

### Summary
- Comprehensive support for **7 languages**: English, Spanish, Portuguese, French, German, Chinese Simplified, Japanese
- Full CJK language support with culturally-appropriate date formats and patterns
- CJK languages have feature parity with Romance languages for most patterns

## [1.3.1] - 2025-10-02

### Added
- **Portuguese Language Support (pt)**: Brazilian Portuguese with comprehensive date parsing
  - Months: `janeiro`, `fevereiro`, ..., `dezembro` (full and abbreviated)
  - Weekdays: `segunda-feira`, `terça-feira`, ..., `domingo` (full and abbreviated forms)
  - Simple relative: `ontem` (yesterday), `hoje` (today), `amanhã` (tomorrow)
  - Ago patterns: `há 2 dias` (2 days ago), `há 1 semana` (1 week ago)
  - Future patterns: `em 3 dias` (in 3 days), `daqui a 2 semanas` (in 2 weeks)
  - Next/last: `próxima segunda` (next Monday), `última sexta` (last Friday)
  - Period boundaries: `início de ano`, `fim de ano`, `início de semana`
  - This/next/last: `esta segunda`, `próxima semana`, `último mês`
  - Time expressions: `meio-dia`, `meia-noite`, `3 e meia`
  - Incomplete dates: `maio`, `junho 15`, `3 de junho`
  - All date formats: `31 dezembro 2024`, `15 de junho de 2024`, `3 de junho de 2024`
  - Works with and without accents: `proximo mes`, `ultimo ano`, `ha 2 dias`

- **New Files**:
  - `translations/portuguese.go` - Portuguese language implementation
  - `portuguese_test.go` - 100+ comprehensive Portuguese test cases

### Changed
- **Test Suite**:
  - 115+ test functions (up from 103)
  - 600+ test cases including Portuguese tests
  - `translations/registry.go` - Added Portuguese to global registry

### Fixed
- **Code Quality**: Fixed all linting issues across the project
  - Fixed 24 unchecked error returns in benchmark functions
  - Removed 3 unused functions (`monthNameToNumber`, `weekdayNameToWeekday`, `getISOWeek`)
  - Removed unused imports and helper functions
  - Project now passes `golangci-lint` with zero issues

### Summary
- Comprehensive support for **English**, **Spanish**, and **Portuguese (Brazil)** languages

## [1.3.0] - 2025-10-02

### Added
- **Multi-Language Support**: Full internationalization infrastructure
  - Translation system with `Language` interface and `Registry`
  - Automatic language detection from input
  - Explicit language selection via `Settings.Languages`
  - Support for multiple languages with priority ordering
  - Localized patterns for months, weekdays, relative terms, time expressions

- **Spanish Language Support (es)**: Comprehensive Spanish language support
  - Months: `enero`, `febrero`, ..., `diciembre` (full and abbreviated)
  - Weekdays: `lunes`, `martes`, ..., `domingo` (with and without accents)
  - Simple relative: `ayer` (yesterday), `hoy` (today), `mañana` (tomorrow)
  - Ago patterns: `hace 2 días` (2 days ago), `hace 1 semana` (1 week ago)
  - Future patterns: `en 3 días` (in 3 days), `dentro de 1 mes` (in 1 month)
  - Next/last: `próximo viernes` (next Friday), `último martes` (last Tuesday)
  - Period boundaries: `inicio de mes`, `fin de año`, `comienzo de semana`
  - This/next/last: `este lunes`, `próxima semana`, `último mes`
  - Time expressions: `mediodía`, `medianoche`, `3 y cuarto`, `9 y media`
  - Incomplete dates: `mayo`, `junio 15`, `3 de junio`
  - All date formats: `31 diciembre 2024`, `15 de marzo de 2024`, `3 de junio 2024`
  - Works with and without accents: `ultimo ano`, `próximo año`

- **New Package**: `translations/`
  - `translations.go` - Core interfaces and types
  - `english.go` - English language implementation
  - `spanish.go` - Spanish language implementation
  - `registry.go` - Global language registry
  - `helpers.go` - Month and weekday parsing utilities

### Changed
- **Parser Updates**: All parsers now support multi-language input
  - `parser_absolute.go` - Multi-language month names
  - `parser_relative.go` - Multi-language relative terms and weekdays
  - `parser_time.go` - Multi-language time expressions
  - `parser_incomplete.go` - Multi-language incomplete dates
  - `parser_ordinal.go` - Multi-language ordinal dates
  - `parser_relative_extended.go` - Multi-language extended expressions

- **Core Changes**:
  - `godateparser.go` - Language loading from registry
  - `parserContext` now includes `languages []*Language`
  - `parser_utils.go` - Added multi-language helper functions

- **Test Suite**:
  - 103 test functions (up from 99)
  - 500+ test cases including 48 Spanish tests
  - New file: `spanish_test.go` - Comprehensive Spanish test coverage

### Performance
- Language detection adds minimal overhead (~2-5μs)
- Dynamic pattern building cached per language
- Overall parsing remains under 50μs for most operations

### Summary
- Comprehensive multi-language support for **English** and **Spanish**

## [1.2.0] - 2025-10-02

### Added
- **ISO Week Number Parsing**: Full support for ISO 8601 week dates
  - ISO 8601 format: `2024-W15`, `2024W15`
  - Natural language: `Week 15 2024`, `2024 Week 15`
  - Week only (current year): `W42`, `Week 15`
  - With specific weekday: `2024-W15-3` (1=Monday, 7=Sunday)
  - Validates week numbers (1-53) and weekdays (1-7)
  - Correctly handles year boundaries per ISO 8601 standard

- **Natural Time Expressions**: Human-friendly time phrases
  - Quarter past: `quarter past 3` → 3:15
  - Half past: `half past 9` → 9:30
  - Quarter to: `quarter to 5` → 4:45
  - With noon/midnight: `quarter past noon`, `half past midnight`
  - Synonyms supported: "past"/"after", "to"/"before"
  - Case insensitive

- **New Files**:
  - `parser_week.go` - ISO week number parsing implementation
  - `parser_week_test.go` - 48 comprehensive tests for week parsing
  - Natural time patterns added to `parser_time.go`

### Changed
- Default enabled parsers increased to 8 (added "week")
- Parser order optimized for new features
- Test suite expanded: 101 test functions, 552 test cases (up from 87/461)

### Performance
- ISO week parsing: ~4.5μs/op
- Natural time parsing: ~6.3μs/op
- All operations remain under 7μs

### Summary
- Comprehensive date parsing features including week numbers and natural time expressions

## [1.1.0] - 2025-10-02

### Added
- **PREFER_DATES_FROM Setting**: Temporal disambiguation for ambiguous dates
  - `"future"` preference: "Monday" → next Monday
  - `"past"` preference: "Monday" → last Monday
  - Works with standalone weekdays, incomplete dates, and ordinals
  - Default: "future"

- **Incomplete Date Parsing**: Dates with missing components
  - Year only: `2024` → January 1, 2024
  - Month only: `May` → May 1 (year inferred)
  - Month + day: `June 15`, `15 June` → year inferred
  - Smart year inference based on PreferDatesFrom setting

- **Ordinal Date Parsing**: British-style date formatting
  - Basic ordinals: `1st`, `2nd`, `3rd`, `21st`, `31st`
  - With month: `June 3rd`, `3rd June`, `3rd of June`
  - Full dates: `June 3rd 2024`, `3rd of June 2024`
  - All variations validated for day/month correctness

- **Additional Relative Terms**: Extended time units
  - Fortnight (14 days): `a fortnight ago`, `next fortnight`
  - Decade (10 years): `a decade ago`, `in a decade`
  - Quarter (3 months): `a quarter ago` (arithmetic)
  - Article support: `a/an` in relative expressions

- **Enhanced Relative Parser**: Articles support
  - `a day ago`, `an hour ago`, `in a week`, `in an hour`

- **New Files**:
  - `parser_incomplete.go` - Incomplete date parsing
  - `parser_ordinal.go` - Ordinal date parsing
  - `parser_utils.go` - Shared utility functions
  - `godateparser_phase4_test.go` - 90 comprehensive tests

### Changed
- Relative parser now supports "a/an" quantifiers
- Duration calculations enhanced with fortnight, decade, quarter
- Parser priority reordered: extended patterns before basic
- Default enabled parsers increased to 7 (added "incomplete", "ordinal")
- Settings struct includes PreferDatesFrom field

### Performance
- Incomplete date parsing: ~5.0μs/op
- Ordinal date parsing: ~5.2μs/op
- Additional relative terms: ~4.5μs/op

## [1.0.0] - 2025-07-09

### Added
- **Extended Relative Expressions** (Phase 3A): Advanced relative date patterns
  - Period boundaries: `beginning of month`, `end of year`, `start of week`
  - This/next/last disambiguation: `this Monday`, `this month`, `this quarter`
  - Complex expressions: `a week from Tuesday`, `3 days after tomorrow`
  - Chained expressions: `2 weeks before last Monday`, `1 week after next Friday`
  - Quarter support: `Q1`, `Q4 2024`, `next quarter`, `last quarter`

- **Time-Only Parsing** (Phase 3B): Comprehensive time parsing
  - 12-hour format: `3:30 PM`, `9am`, `noon`
  - 24-hour format: `14:30`, `23:59:59`, `midnight`
  - Integration with RelativeBase for date context
  - Natural language: `noon`, `midnight`

- **Date Range Parsing** (Phase 3B): Range expressions
  - From/to patterns: `from X to Y`, `between X and Y`
  - Duration ranges: `next 7 days`, `last 2 weeks`, `next 3 months`
  - Helper functions: `GetDatesInRange()`, `GetBusinessDaysInRange()`
  - Utility functions: `DaysBetween()`, `DurationBetween()`
  - Smart splitting for multi-word date expressions

- **New API Functions**:
  - `ParseDateRange(input string, opts *Settings)` - Parse date range expressions
  - Range helper functions for date manipulation

- **New Types**:
  - `DateRange` struct with Start, End, and Input fields

- **New Files**:
  - `parser_relative_extended.go` - Extended relative parsing
  - `parser_time.go` - Time-only parsing
  - `range.go` - Date range parsing and utilities
  - Comprehensive test files for all new features

### Changed
- Relative parser prioritizes extended patterns first
- Default enabled parsers increased to 5 (added "time")
- Test suite expanded: 371 tests (from 244 in v0.3)

### Performance
- Period boundaries: ~5.6μs/op
- Complex expressions: ~6.6μs/op
- Time parsing: ~4.0μs/op
- Range parsing: ~8.0μs/op

## [0.3.0] - 2025-06-12

### Added
- **Timezone Support**: Comprehensive timezone parsing and handling
  - 30+ common timezone abbreviations (EST, PST, GMT, UTC, CET, JST, etc.)
  - Timezone offset parsing: `+05:00`, `-08:00`, `+0530`, `-0800`
  - Named offsets: `UTC+5`, `GMT-8`, `UTC+05:30`
  - ISO 8601 with timezone: `2024-12-31T10:30:00Z`, `2024-12-31T10:30:00+05:00`
  - Date strings with timezone: `2024-12-31 10:30:00 EST`
  - Automatic timezone extraction from date strings
  - DST-aware via IANA timezone database
  - Ambiguous timezone detection and flagging

- **New API Functions**:
  - `ParseTimezone(tz string)` - Parse timezone strings
  - `ExtractTimezone(input string)` - Extract timezone from dates
  - `ApplyTimezone(t time.Time, tzInfo *TimezoneInfo)` - Apply timezone

- **New Types**:
  - `TimezoneInfo` struct with location, offset, name, ambiguity flag

- **New Files**:
  - `timezone.go` - Timezone parsing implementation
  - `timezone_test.go` - 48 comprehensive timezone tests

### Changed
- Absolute date parser automatically extracts and applies timezones
- ISO 8601 parser enhanced for Z suffix and offsets

### Fixed
- Prevented false positive timezone matches in numeric dates

### Performance
- Timezone abbreviation parsing: ~64μs/op
- Timezone offset parsing: ~317ns/op
- Date with timezone: ~1μs/op

## [0.2.0] - 2025-06-05

### Added
- **Custom Error Types**: Structured error handling
  - `ErrEmptyInput` - Empty input strings
  - `ErrInvalidFormat` - Unrecognized formats with suggestions
  - `ErrInvalidDate` - Invalid date components
  - `ErrAmbiguousDate` - Ambiguous dates in strict mode
  - `ErrParseFailure` - Generic parse errors with context

- **Two-Digit Year Support**: Automatic 2-digit year interpretation
  - Years 00-69 → 2000-2069
  - Years 70-99 → 1970-1999
  - Works with all date formats

- **Enhanced Date Validation**: Comprehensive validation
  - Invalid months (< 1 or > 12)
  - Invalid days (< 1 or > 31)
  - Month/day combinations (Feb 31, Apr 31)
  - Leap year validation
  - Time component validation

- **Strict Mode**: Ambiguity detection
  - Detects ambiguous numeric dates (01/02/2024)
  - Returns ErrAmbiguousDate when ambiguous
  - Respects explicit DateOrder settings

- **Auto-Detection**: Intelligent date order detection
  - Auto-detects DMY when first number > 12
  - Auto-detects MDY when second number > 12
  - Falls back to configured DateOrder

- **Improved Error Messages**: Contextual errors with suggestions

- **New Files**:
  - `errors.go` - Custom error types
  - `validation.go` - Date validation functions
  - `godateparser_v02_test.go` - 64 comprehensive tests

### Changed
- ParseDate returns specific error types
- More rigorous date validation
- Improved settings normalization

### Fixed
- Invalid dates (Feb 31) now rejected
- Month values > 12 now rejected
- Time components validated
- Ambiguous dates detected in strict mode

## [0.1.0] - 2025-06-05

### Added
- Initial release
- Core date parsing functionality
  - Absolute dates: ISO 8601, numeric formats, month names
  - Relative dates: yesterday, 2 days ago, next Monday
  - Unix timestamps: seconds and milliseconds
- Date extraction from text
- Customizable settings: DateOrder, RelativeBase, Languages
- Comprehensive test suite: 128 tests
- Full documentation: README, QUICKSTART, examples
- MIT License

### Performance
- ISO date parsing: ~600 ns/op
- Relative date parsing: ~1.5 μs/op
- Text extraction: ~65 μs/op

## Summary

- **v0.1.0**: Initial release with core parsing (128 tests)
- **v0.2.0**: Production hardening with validation and error handling (192 tests)
- **v0.3.0**: Timezone support (244 tests)
- **v1.0.0**: Extended relative expressions, time parsing, date ranges (371 tests)
- **v1.1.0**: Incomplete dates, ordinals, PREFER_DATES_FROM (455 tests)
- **v1.2.0**: Week numbers, natural time expressions (552 tests)

**Total growth: 431% increase in test coverage from v0.1.0 to v1.2.0**
