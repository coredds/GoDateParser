package translations

import (
	"time"
)

// NewGermanTranslation creates the German (Germany) language translation.
func NewGermanTranslation() *Language {
	return &Language{
		Code: "de",
		Name: "German",
		Months: map[string]time.Month{
			// Full names
			"januar":  time.January,
			"februar": time.February,
			"märz":    time.March, "marz": time.March,
			"april":     time.April,
			"mai":       time.May,
			"juni":      time.June,
			"juli":      time.July,
			"august":    time.August,
			"september": time.September,
			"oktober":   time.October,
			"november":  time.November,
			"dezember":  time.December,
			// Abbreviations
			"jan": time.January,
			"feb": time.February,
			"mär": time.March, "mar": time.March,
			"apr": time.April,
			"jun": time.June,
			"jul": time.July,
			"aug": time.August,
			"sep": time.September, "sept": time.September,
			"okt": time.October,
			"nov": time.November,
			"dez": time.December,
		},
		Weekdays: map[string]time.Weekday{
			// Full names
			"montag":     time.Monday,
			"dienstag":   time.Tuesday,
			"mittwoch":   time.Wednesday,
			"donnerstag": time.Thursday,
			"freitag":    time.Friday,
			"samstag":    time.Saturday, "sonnabend": time.Saturday,
			"sonntag": time.Sunday,
			// Abbreviations
			"mo": time.Monday,
			"di": time.Tuesday,
			"mi": time.Wednesday,
			"do": time.Thursday,
			"fr": time.Friday,
			"sa": time.Saturday,
			"so": time.Sunday,
		},
		RelativeTerms: &RelativeTerms{
			Yesterday: "gestern",
			Today:     "heute",
			Tomorrow:  "morgen",
			Now:       "jetzt",
			// "vor" for past (vor 2 Tagen = 2 days ago)
			Ago: []string{"vor"},
			// "in" for future (in 3 Wochen = in 3 weeks)
			In: []string{"in"},
			// Gender variations for "next" (nächster/nächste/nächstes)
			Next: []string{"nächster", "nächste", "nächstes", "naechster", "naechste", "naechstes", "kommender", "kommende", "kommendes"},
			// Gender variations for "last" (letzter/letzte/letztes)
			Last: []string{"letzter", "letzte", "letztes", "vergangener", "vergangene", "vergangenes", "vorletzter", "vorletzte", "vorletztes"},
			// Gender variations for "this" (dieser/diese/dieses)
			This: []string{"dieser", "diese", "dieses"},
			// Time units with plural forms
			Second:    []string{"sekunde", "sekunden"},
			Minute:    []string{"minute", "minuten"},
			Hour:      []string{"stunde", "stunden"},
			Day:       []string{"tag", "tage", "tagen"},
			Week:      []string{"woche", "wochen"},
			Fortnight: []string{"vierzehn tage", "zwei wochen"},
			Month:     []string{"monat", "monate", "monaten"},
			Quarter:   []string{"quartal", "quartale"},
			Year:      []string{"jahr", "jahre", "jahren"},
			Decade:    []string{"jahrzehnt", "jahrzehnte", "dekade", "dekaden"},
			// Period boundaries
			Beginning: []string{"anfang", "beginn", "start"},
			End:       []string{"ende", "schluss"},
			Start:     []string{"anfang", "beginn", "start"},
			First:     []string{"erster", "erste", "erstes"},
		},
		TimeTerms: &TimeTerms{
			Noon:     []string{"mittag", "12 uhr mittags"},
			Midnight: []string{"mitternacht", "24 uhr", "0 uhr"},
			Quarter:  []string{"viertel"},
			Half:     []string{"halb", "halbe"},
			// "nach" for past (viertel nach 3 = quarter past 3)
			Past: []string{"nach"},
			// "vor" for to (viertel vor 3 = quarter to 3)
			To:     []string{"vor"},
			OClock: []string{"uhr"},
			AM:     []string{"uhr", "morgens", "vormittags"},
			PM:     []string{"uhr", "nachmittags", "abends", "nachts"},
		},
	}
}
