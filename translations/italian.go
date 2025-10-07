package translations

import (
	"time"
)

// NewItalianTranslation creates the Italian (Italy) language translation.
func NewItalianTranslation() *Language {
	return &Language{
		Code: "it",
		Name: "Italian",
		Months: map[string]time.Month{
			// Full names
			"gennaio":   time.January,
			"febbraio":  time.February,
			"marzo":     time.March,
			"aprile":    time.April,
			"maggio":    time.May,
			"giugno":    time.June,
			"luglio":    time.July,
			"agosto":    time.August,
			"settembre": time.September,
			"ottobre":   time.October,
			"novembre":  time.November,
			"dicembre":  time.December,
			// Abbreviations
			"gen": time.January,
			"feb": time.February,
			"mar": time.March,
			"apr": time.April,
			"mag": time.May,
			"giu": time.June,
			"lug": time.July,
			"ago": time.August,
			"set": time.September, "sett": time.September,
			"ott": time.October,
			"nov": time.November,
			"dic": time.December,
		},
		Weekdays: map[string]time.Weekday{
			// Full names
			"lunedì": time.Monday, "lunedi": time.Monday,
			"martedì": time.Tuesday, "martedi": time.Tuesday,
			"mercoledì": time.Wednesday, "mercoledi": time.Wednesday,
			"giovedì": time.Thursday, "giovedi": time.Thursday,
			"venerdì": time.Friday, "venerdi": time.Friday,
			"sabato":   time.Saturday,
			"domenica": time.Sunday,
			// Abbreviations
			"lun": time.Monday,
			"mar": time.Tuesday,
			"mer": time.Wednesday,
			"gio": time.Thursday,
			"ven": time.Friday,
			"sab": time.Saturday,
			"dom": time.Sunday,
		},
		RelativeTerms: &RelativeTerms{
			Yesterday: "ieri",
			Today:     "oggi",
			Tomorrow:  "domani",
			Now:       "adesso",
			// "fa" for past (2 giorni fa = 2 days ago)
			Ago: []string{"fa"},
			// "tra/fra" for future (tra 3 settimane = in 3 weeks)
			In: []string{"tra", "fra", "in"},
			// Gender variations for "next" (prossimo/prossima)
			Next: []string{"prossimo", "prossima", "venturo", "ventura"},
			// Gender variations for "last" (scorso/scorsa, ultimo/ultima)
			Last: []string{"scorso", "scorsa", "ultimo", "ultima", "passato", "passata"},
			// Gender variations for "this" (questo/questa)
			This: []string{"questo", "questa"},
			// Time units with plural forms
			Second:    []string{"secondo", "secondi"},
			Minute:    []string{"minuto", "minuti"},
			Hour:      []string{"ora", "ore"},
			Day:       []string{"giorno", "giorni"},
			Week:      []string{"settimana", "settimane"},
			Fortnight: []string{"quindicina", "quindici giorni"},
			Month:     []string{"mese", "mesi"},
			Quarter:   []string{"trimestre", "trimestri"},
			Year:      []string{"anno", "anni"},
			Decade:    []string{"decennio", "decenni", "decade", "decadi"},
			// Period boundaries
			Beginning: []string{"inizio", "inizio", "principio"},
			End:       []string{"fine", "termine"},
			Start:     []string{"inizio", "avvio"},
			First:     []string{"primo", "prima"},
		},
		TimeTerms: &TimeTerms{
			Noon:     []string{"mezzogiorno", "mezzo giorno"},
			Midnight: []string{"mezzanotte", "mezza notte"},
			Quarter:  []string{"quarto"},
			Half:     []string{"mezzo", "mezza"},
			// "e" for past (3 e un quarto = quarter past 3)
			Past: []string{"e"},
			// "meno" for to (meno un quarto = quarter to)
			To:     []string{"meno"},
			OClock: []string{"in punto"},
			AM:     []string{"am", "a.m.", "di mattina", "del mattino"},
			PM:     []string{"pm", "p.m.", "di pomeriggio", "del pomeriggio", "di sera", "della sera"},
		},
	}
}
