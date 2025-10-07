# Language Examples

This document provides comprehensive examples for all supported languages in godateparser. For basic usage and API documentation, see the main [README.md](README.md).

## Table of Contents

- [English (en)](#english-en)
- [Spanish (es)](#spanish-es)
- [Portuguese (pt)](#portuguese-pt)
- [French (fr)](#french-fr)
- [German (de)](#german-de)
- [Italian (it)](#italian-it)
- [Dutch (nl)](#dutch-nl)
- [Russian (ru)](#russian-ru)
- [Chinese Simplified (zh)](#chinese-simplified-zh)
- [Japanese (ja)](#japanese-ja)
- [Mixed Language Usage](#mixed-language-usage)

---

## English (en)

English is the default language with full support for all features.

```go
settings := &godateparser.Settings{
    Languages: []string{"en"},
}

// Absolute dates
godateparser.ParseDate("December 31, 2024", settings)
godateparser.ParseDate("31 Dec 2024", settings)
godateparser.ParseDate("12/31/2024", settings)

// Relative dates
godateparser.ParseDate("yesterday", settings)
godateparser.ParseDate("today", settings)
godateparser.ParseDate("tomorrow", settings)
godateparser.ParseDate("2 days ago", settings)
godateparser.ParseDate("in 3 weeks", settings)

// Weekdays
godateparser.ParseDate("Monday", settings)
godateparser.ParseDate("next Friday", settings)
godateparser.ParseDate("last Tuesday", settings)

// Period boundaries
godateparser.ParseDate("beginning of month", settings)
godateparser.ParseDate("end of year", settings)
godateparser.ParseDate("start of next week", settings)

// Time expressions
godateparser.ParseDate("noon", settings)
godateparser.ParseDate("midnight", settings)
godateparser.ParseDate("quarter past 3", settings)
godateparser.ParseDate("half past 9", settings)
```

---

## Spanish (es)

Spanish has full support with gender variations and accent-optional parsing.

### Absolute Dates

```go
settings := &godateparser.Settings{Languages: []string{"es"}}

// Month names
godateparser.ParseDate("31 diciembre 2024", settings)
godateparser.ParseDate("15 de marzo de 2024", settings)
godateparser.ParseDate("marzo 15 2024", settings)

// Abbreviated months
godateparser.ParseDate("25 dic 2024", settings)
godateparser.ParseDate("15 mar 2024", settings)

// Without accents (also supported)
godateparser.ParseDate("31 de diciembre de 2024", settings)
```

### Relative Dates

```go
settings := &godateparser.Settings{
    Languages: []string{"es"},
    RelativeBase: time.Now(),
}

// Simple terms
godateparser.ParseDate("ayer", settings)      // yesterday
godateparser.ParseDate("hoy", settings)       // today
godateparser.ParseDate("mañana", settings)    // tomorrow

// Ago patterns
godateparser.ParseDate("hace 2 días", settings)     // 2 days ago
godateparser.ParseDate("hace 1 semana", settings)   // 1 week ago
godateparser.ParseDate("hace 3 meses", settings)    // 3 months ago

// Future patterns
godateparser.ParseDate("en 3 días", settings)       // in 3 days
godateparser.ParseDate("en 2 semanas", settings)    // in 2 weeks
godateparser.ParseDate("dentro de 1 mes", settings) // in 1 month

// Weekdays
godateparser.ParseDate("lunes", settings)           // Monday
godateparser.ParseDate("próximo viernes", settings) // next Friday
godateparser.ParseDate("último martes", settings)   // last Tuesday
```

### Extended Relative Dates

```go
settings := &godateparser.Settings{
    Languages: []string{"es"},
    RelativeBase: time.Date(2024, 10, 15, 0, 0, 0, 0, time.UTC),
}

// Period boundaries
godateparser.ParseDate("inicio de mes", settings)          // October 1, 2024
godateparser.ParseDate("fin de mes", settings)             // October 31, 2024
godateparser.ParseDate("comienzo de año", settings)        // January 1, 2024
godateparser.ParseDate("fin de año", settings)             // December 31, 2024

// This/next/last
godateparser.ParseDate("este lunes", settings)             // this Monday
godateparser.ParseDate("esta semana", settings)            // this week (Monday)
godateparser.ParseDate("próxima semana", settings)         // next week
godateparser.ParseDate("último mes", settings)             // last month

// Combined patterns
godateparser.ParseDate("inicio de próximo mes", settings)  // November 1, 2024
godateparser.ParseDate("fin de último año", settings)      // December 31, 2023
```

### Time Expressions

```go
settings := &godateparser.Settings{Languages: []string{"es"}}

// Special times
godateparser.ParseDate("mediodía", settings)       // 12:00 PM
godateparser.ParseDate("medianoche", settings)     // 12:00 AM

// Quarter/half past/to
godateparser.ParseDate("3 y cuarto", settings)     // 3:15
godateparser.ParseDate("9 y media", settings)      // 9:30
godateparser.ParseDate("menos cuarto las 5", settings) // 4:45
```

---

## Portuguese (pt)

Portuguese (Brazilian) with full support including accent-optional parsing.

```go
settings := &godateparser.Settings{
    Languages: []string{"pt"},
}

// Months and weekdays
godateparser.ParseDate("15 de junho de 2024", settings)  // June 15, 2024
godateparser.ParseDate("25 dezembro 2024", settings)     // December 25, 2024
godateparser.ParseDate("segunda-feira", settings)         // Next Monday
godateparser.ParseDate("sexta", settings)                 // Next Friday

// Simple relative dates
godateparser.ParseDate("ontem", settings)      // Yesterday
godateparser.ParseDate("hoje", settings)       // Today  
godateparser.ParseDate("amanhã", settings)     // Tomorrow

// Relative dates with time units
godateparser.ParseDate("há 2 dias", settings)        // 2 days ago
godateparser.ParseDate("há 1 semana", settings)      // 1 week ago
godateparser.ParseDate("em 3 dias", settings)        // in 3 days
godateparser.ParseDate("daqui a 2 semanas", settings) // in 2 weeks

// Next/Last patterns
godateparser.ParseDate("próxima segunda", settings)   // Next Monday
godateparser.ParseDate("última sexta", settings)      // Last Friday
godateparser.ParseDate("próximo mês", settings)       // Next month
godateparser.ParseDate("último ano", settings)        // Last year

// Time expressions
godateparser.ParseDate("meio-dia", settings)      // 12:00
godateparser.ParseDate("meia-noite", settings)    // 00:00
godateparser.ParseDate("3 e meia", settings)      // 3:30

// Works with or without accents
godateparser.ParseDate("proximo mes", settings)   // Next month (no accent)
godateparser.ParseDate("ultimo ano", settings)    // Last year (no accent)
godateparser.ParseDate("ha 2 dias", settings)     // 2 days ago (no accent)
```

---

## French (fr)

French with full support including accent-optional parsing.

```go
settings := &godateparser.Settings{
    Languages: []string{"fr"},
}

// Months and weekdays
godateparser.ParseDate("31 décembre 2024", settings)  // December 31, 2024
godateparser.ParseDate("15 juin 2024", settings)      // June 15, 2024
godateparser.ParseDate("7 février 2025", settings)    // February 7, 2025
godateparser.ParseDate("lundi", settings)             // Next Monday
godateparser.ParseDate("vendredi", settings)          // Next Friday

// Simple relative dates
godateparser.ParseDate("hier", settings)          // Yesterday
godateparser.ParseDate("aujourd'hui", settings)   // Today
godateparser.ParseDate("demain", settings)        // Tomorrow

// Relative dates with time units
godateparser.ParseDate("il y a 2 jours", settings)    // 2 days ago
godateparser.ParseDate("il y a 1 semaine", settings)  // 1 week ago
godateparser.ParseDate("dans 3 jours", settings)      // in 3 days
godateparser.ParseDate("dans 2 semaines", settings)   // in 2 weeks
godateparser.ParseDate("en 1 mois", settings)         // in 1 month

// Next/Last patterns
godateparser.ParseDate("prochain lundi", settings)    // Next Monday
godateparser.ParseDate("dernier vendredi", settings)  // Last Friday
godateparser.ParseDate("prochaine semaine", settings) // Next week
godateparser.ParseDate("prochain mois", settings)     // Next month
godateparser.ParseDate("dernier mois", settings)      // Last month

// Time expressions
godateparser.ParseDate("midi", settings)      // 12:00
godateparser.ParseDate("minuit", settings)    // 00:00
godateparser.ParseDate("15h30", settings)     // 3:30 PM
godateparser.ParseDate("15h", settings)       // 3:00 PM

// Works with or without accents
godateparser.ParseDate("decembre", settings)      // December (no accent)
godateparser.ParseDate("fevrier", settings)       // February (no accent)
godateparser.ParseDate("derniere semaine", settings)  // Last week (no accent)
```

---

## German (de)

German with full support including umlaut-optional parsing and gender variations.

```go
settings := &godateparser.Settings{
    Languages: []string{"de"},
}

// Months and weekdays
godateparser.ParseDate("31 Dezember 2024", settings)  // December 31, 2024
godateparser.ParseDate("15 Juni 2024", settings)      // June 15, 2024
godateparser.ParseDate("7 Februar 2025", settings)    // February 7, 2025
godateparser.ParseDate("Montag", settings)            // Next Monday
godateparser.ParseDate("Freitag", settings)           // Next Friday

// Simple relative dates
godateparser.ParseDate("gestern", settings)    // Yesterday
godateparser.ParseDate("heute", settings)      // Today
godateparser.ParseDate("morgen", settings)     // Tomorrow

// Relative dates with time units
godateparser.ParseDate("vor 2 Tagen", settings)        // 2 days ago
godateparser.ParseDate("vor 1 Woche", settings)        // 1 week ago
godateparser.ParseDate("in 3 Tagen", settings)         // in 3 days
godateparser.ParseDate("in 2 Wochen", settings)        // in 2 weeks
godateparser.ParseDate("in 1 Monat", settings)         // in 1 month

// Next/Last patterns
godateparser.ParseDate("nächster Montag", settings)    // Next Monday
godateparser.ParseDate("letzter Freitag", settings)    // Last Friday
godateparser.ParseDate("nächste Woche", settings)      // Next week
godateparser.ParseDate("nächster Monat", settings)     // Next month
godateparser.ParseDate("letzter Monat", settings)      // Last month
godateparser.ParseDate("kommender Montag", settings)   // Coming Monday
godateparser.ParseDate("vergangener Freitag", settings) // Past Friday

// Time expressions
godateparser.ParseDate("Mittag", settings)      // 12:00
godateparser.ParseDate("Mitternacht", settings) // 00:00
godateparser.ParseDate("15:30", settings)       // 3:30 PM

// Works with or without umlauts
godateparser.ParseDate("Marz", settings)           // März (March, no umlaut)
godateparser.ParseDate("naechste Woche", settings) // nächste Woche (no umlaut)
```

---

## Italian (it)

Italian with full support including accent-optional parsing and gender variations.

```go
settings := &godateparser.Settings{
    Languages: []string{"it"},
}

// Months and weekdays
godateparser.ParseDate("31 dicembre 2024", settings)  // December 31, 2024
godateparser.ParseDate("15 giugno 2024", settings)    // June 15, 2024
godateparser.ParseDate("7 febbraio 2025", settings)   // February 7, 2025
godateparser.ParseDate("lunedì", settings)            // Next Monday
godateparser.ParseDate("venerdì", settings)           // Next Friday

// Simple relative dates
godateparser.ParseDate("ieri", settings)      // Yesterday
godateparser.ParseDate("oggi", settings)      // Today
godateparser.ParseDate("domani", settings)    // Tomorrow

// Relative dates with time units
godateparser.ParseDate("2 giorni fa", settings)       // 2 days ago
godateparser.ParseDate("1 settimana fa", settings)    // 1 week ago
godateparser.ParseDate("tra 3 giorni", settings)      // in 3 days
godateparser.ParseDate("fra 2 settimane", settings)   // in 2 weeks
godateparser.ParseDate("in 1 mese", settings)         // in 1 month

// Next/Last patterns
godateparser.ParseDate("prossimo lunedì", settings)   // Next Monday
godateparser.ParseDate("scorso venerdì", settings)    // Last Friday
godateparser.ParseDate("prossima settimana", settings) // Next week
godateparser.ParseDate("prossimo mese", settings)     // Next month
godateparser.ParseDate("scorso mese", settings)       // Last month

// Time expressions
godateparser.ParseDate("mezzogiorno", settings)   // 12:00
godateparser.ParseDate("mezzanotte", settings)    // 00:00
godateparser.ParseDate("3 e un quarto", settings) // 3:15
godateparser.ParseDate("9 e mezzo", settings)     // 9:30

// Period boundaries
godateparser.ParseDate("inizio di mese", settings)  // Beginning of month
godateparser.ParseDate("fine di anno", settings)    // End of year

// Works with or without accents
godateparser.ParseDate("lunedi", settings)        // lunedì (no accent)
godateparser.ParseDate("venerdi", settings)       // venerdì (no accent)
```

---

## Dutch (nl)

Dutch with full support including multiple variations for next/last patterns.

```go
settings := &godateparser.Settings{
    Languages: []string{"nl"},
}

// Months and weekdays
godateparser.ParseDate("31 december 2024", settings)  // December 31, 2024
godateparser.ParseDate("15 juni 2024", settings)      // June 15, 2024
godateparser.ParseDate("7 februari 2025", settings)   // February 7, 2025
godateparser.ParseDate("maandag", settings)           // Next Monday
godateparser.ParseDate("vrijdag", settings)           // Next Friday

// Simple relative dates
godateparser.ParseDate("gisteren", settings)  // Yesterday
godateparser.ParseDate("vandaag", settings)   // Today
godateparser.ParseDate("morgen", settings)    // Tomorrow

// Relative dates with time units
godateparser.ParseDate("2 dagen geleden", settings)   // 2 days ago
godateparser.ParseDate("1 week geleden", settings)    // 1 week ago
godateparser.ParseDate("over 3 dagen", settings)      // in 3 days
godateparser.ParseDate("over 2 weken", settings)      // in 2 weeks
godateparser.ParseDate("in 1 maand", settings)        // in 1 month

// Next/Last patterns (multiple variations)
godateparser.ParseDate("volgende maandag", settings)  // Next Monday
godateparser.ParseDate("vorige vrijdag", settings)    // Last Friday
godateparser.ParseDate("volgende week", settings)     // Next week
godateparser.ParseDate("komende maand", settings)     // Coming month
godateparser.ParseDate("afgelopen week", settings)    // Last week (literally: past week)
godateparser.ParseDate("aanstaande vrijdag", settings) // Upcoming Friday

// Time expressions
godateparser.ParseDate("middag", settings)        // 12:00
godateparser.ParseDate("middernacht", settings)   // 00:00
godateparser.ParseDate("kwart over 3", settings)  // 3:15
godateparser.ParseDate("half 4", settings)        // 3:30 (Note: Dutch "half 4" means 3:30!)
godateparser.ParseDate("kwart voor 3", settings)  // 2:45

// Period boundaries
godateparser.ParseDate("begin van maand", settings)  // Beginning of month
godateparser.ParseDate("einde van jaar", settings)   // End of year
```

---

## Russian (ru)

Russian with full Cyrillic support, grammatical cases, and gender agreement.

```go
settings := &godateparser.Settings{
    Languages: []string{"ru"},
}

// Months and weekdays (with grammatical cases)
godateparser.ParseDate("31 декабря 2024", settings)   // December 31, 2024 (genitive case)
godateparser.ParseDate("15 июня 2024", settings)      // June 15, 2024
godateparser.ParseDate("7 февраля 2025", settings)    // February 7, 2025
godateparser.ParseDate("понедельник", settings)       // Next Monday (nominative)
godateparser.ParseDate("пятница", settings)           // Next Friday
godateparser.ParseDate("в понедельнике", settings)    // On Monday (prepositional case)

// Simple relative dates
godateparser.ParseDate("вчера", settings)     // Yesterday
godateparser.ParseDate("сегодня", settings)   // Today
godateparser.ParseDate("завтра", settings)    // Tomorrow

// Relative dates with time units (with Russian plural forms)
godateparser.ParseDate("1 день назад", settings)      // 1 day ago (nominative)
godateparser.ParseDate("2 дня назад", settings)       // 2 days ago (genitive singular: 2-4)
godateparser.ParseDate("5 дней назад", settings)      // 5 days ago (genitive plural: 5+)
godateparser.ParseDate("1 неделя назад", settings)    // 1 week ago
godateparser.ParseDate("через 3 дня", settings)       // in 3 days
godateparser.ParseDate("через 2 недели", settings)    // in 2 weeks
godateparser.ParseDate("спустя 1 месяц", settings)    // in 1 month (alternative)

// Next/Last patterns (with gender agreement)
godateparser.ParseDate("следующий понедельник", settings)  // Next Monday (masculine)
godateparser.ParseDate("прошлая пятница", settings)        // Last Friday (feminine)
godateparser.ParseDate("следующая неделя", settings)       // Next week (feminine)
godateparser.ParseDate("будущий месяц", settings)          // Future month (masculine)
godateparser.ParseDate("предыдущая неделя", settings)      // Previous week (feminine)

// Time expressions
godateparser.ParseDate("полдень", settings)       // 12:00 (noon)
godateparser.ParseDate("полночь", settings)       // 00:00 (midnight)
godateparser.ParseDate("3 часа дня", settings)    // 3:00 PM
godateparser.ParseDate("9 часов утра", settings)  // 9:00 AM
godateparser.ParseDate("7 часов вечера", settings) // 7:00 PM

// Period boundaries
godateparser.ParseDate("начало месяца", settings)  // Beginning of month
godateparser.ParseDate("конец года", settings)     // End of year
godateparser.ParseDate("начало недели", settings)  // Beginning of week
```

---

## Chinese Simplified (zh)

Chinese with full support including multiple weekday forms and CJK date formats.

```go
settings := &godateparser.Settings{
    Languages: []string{"zh"},
}

// Chinese weekdays - multiple forms
godateparser.ParseDate("星期一", settings)   // Monday (formal: xīngqī)
godateparser.ParseDate("周一", settings)     // Monday (common: zhōu)
godateparser.ParseDate("礼拜一", settings)   // Monday (colloquial: lǐbài)
godateparser.ParseDate("星期天", settings)   // Sunday alternative

// Simple relative dates
godateparser.ParseDate("昨天", settings)     // yesterday (zuótiān)
godateparser.ParseDate("今天", settings)     // today (jīntiān)
godateparser.ParseDate("明天", settings)     // tomorrow (míngtiān)

// Chinese months
godateparser.ParseDate("1月", settings)      // January
godateparser.ParseDate("5月", settings)      // May
godateparser.ParseDate("12月", settings)     // December

// Chinese date formats
godateparser.ParseDate("2024年12月31日", settings)  // 2024-12-31
godateparser.ParseDate("2025年1月1日", settings)    // 2025-01-01

// Relative patterns with numbers
godateparser.ParseDate("3天前", settings)           // 3 days ago
godateparser.ParseDate("2周后", settings)           // in 2 weeks
godateparser.ParseDate("1个月前", settings)         // 1 month ago

// Next/Last patterns
godateparser.ParseDate("下周", settings)             // next week
godateparser.ParseDate("上月", settings)             // last month
godateparser.ParseDate("明年", settings)             // next year

// Weekday modifiers - next/last week with specific weekday
godateparser.ParseDate("下周一", settings)           // next week Monday (using 周)
godateparser.ParseDate("上周五", settings)           // last week Friday (using 周)
godateparser.ParseDate("下星期一", settings)         // next week Monday (using 星期)

// Time expressions
godateparser.ParseDate("中午", settings)     // noon (zhōngwǔ)
godateparser.ParseDate("午夜", settings)     // midnight (wǔyè)
```

---

## Japanese (ja)

Japanese with full support including multiple weekday forms and CJK date formats.

```go
settings := &godateparser.Settings{
    Languages: []string{"ja"},
}

// Japanese weekdays - multiple forms
godateparser.ParseDate("月曜日", settings)   // Monday (getsuyoubi - full)
godateparser.ParseDate("月曜", settings)     // Monday (getsuyo - short)
godateparser.ParseDate("火曜日", settings)   // Tuesday
godateparser.ParseDate("水曜", settings)     // Wednesday (short)

// Simple relative dates
godateparser.ParseDate("昨日", settings)     // yesterday (kinou)
godateparser.ParseDate("今日", settings)     // today (kyou)
godateparser.ParseDate("明日", settings)     // tomorrow (ashita)

// Japanese months
godateparser.ParseDate("1月", settings)      // January (ichigatsu)
godateparser.ParseDate("5月", settings)      // May (gogatsu)
godateparser.ParseDate("12月", settings)     // December (juunigatsu)

// Japanese date formats
godateparser.ParseDate("2024年12月31日", settings)  // 2024-12-31
godateparser.ParseDate("2025年1月1日", settings)    // 2025-01-01

// Relative patterns with numbers
godateparser.ParseDate("3日前", settings)           // 3 days ago (mikka mae)
godateparser.ParseDate("2週後", settings)           // in 2 weeks (ni shuu go)
godateparser.ParseDate("1ヶ月前", settings)         // 1 month ago

// Next/Last patterns
godateparser.ParseDate("来週", settings)             // next week (raishuu)
godateparser.ParseDate("先週", settings)             // last week (senshuu)
godateparser.ParseDate("来月", settings)             // next month (raigetsu)
godateparser.ParseDate("先月", settings)             // last month (sengetsu)

// Weekday modifiers - next/last week with specific weekday
godateparser.ParseDate("来週月曜", settings)         // next week Monday (raishuu getsuyou)
godateparser.ParseDate("先週金曜", settings)         // last week Friday (senshuu kinyou)
godateparser.ParseDate("来週月曜日", settings)       // next week Monday (with 日 suffix)

// Time expressions
godateparser.ParseDate("正午", settings)     // noon (shougo)
godateparser.ParseDate("真夜中", settings)   // midnight (mayonaka)
```

---

## Mixed Language Usage

You can enable multiple languages simultaneously for maximum flexibility.

```go
// Enable both English and Spanish
settings := &godateparser.Settings{
    Languages: []string{"es", "en"},
}

// Both languages work
godateparser.ParseDate("31 diciembre 2024", settings)  // Spanish
godateparser.ParseDate("December 31, 2024", settings)  // English

// Can even mix within same application
godateparser.ParseDate("próximo lunes", settings)      // Spanish: next Monday
godateparser.ParseDate("next Friday", settings)        // English: next Friday

// Enable all ten languages
settings = &godateparser.Settings{
    Languages: []string{"ja", "zh", "ru", "nl", "it", "de", "fr", "pt", "es", "en"},
}

// All languages work simultaneously
godateparser.ParseDate("明日", settings)               // Japanese: tomorrow
godateparser.ParseDate("明天", settings)               // Chinese: tomorrow
godateparser.ParseDate("завтра", settings)            // Russian: tomorrow
godateparser.ParseDate("morgen", settings)            // Dutch/German: tomorrow
godateparser.ParseDate("domani", settings)            // Italian: tomorrow
godateparser.ParseDate("demain", settings)            // French: tomorrow
godateparser.ParseDate("amanhã", settings)            // Portuguese: tomorrow
godateparser.ParseDate("mañana", settings)            // Spanish: tomorrow  
godateparser.ParseDate("tomorrow", settings)          // English: tomorrow
```

---

## Notes

- **Accent-optional parsing**: Most Romance languages (Spanish, Portuguese, French, Italian) support parsing with or without accents for better compatibility with ASCII-only input.
- **Umlaut-optional parsing**: German supports parsing with or without umlauts (ä, ö, ü).
- **Grammatical cases**: Russian properly handles different grammatical cases (nominative, genitive, prepositional) for months and weekdays.
- **Gender agreement**: Languages with grammatical gender (Spanish, Portuguese, French, German, Italian, Russian) support appropriate gender variations for modifiers like "next" and "last".
- **Multiple forms**: Chinese and Japanese support multiple forms for weekdays and time expressions.
- **Plural forms**: Russian implements proper plural forms (1, 2-4, 5+) for time units.

For more examples and API documentation, see the main [README.md](README.md).

