package translations

import (
	"time"
)

// NewDutchTranslation creates the Dutch (Netherlands) language translation.
func NewDutchTranslation() *Language {
	return &Language{
		Code: "nl",
		Name: "Dutch",
		Months: map[string]time.Month{
			// Full names
			"januari":   time.January,
			"februari":  time.February,
			"maart":     time.March,
			"april":     time.April,
			"mei":       time.May,
			"juni":      time.June,
			"juli":      time.July,
			"augustus":  time.August,
			"september": time.September,
			"oktober":   time.October,
			"november":  time.November,
			"december":  time.December,
			// Abbreviations
			"jan": time.January,
			"feb": time.February,
			"mrt": time.March,
			"apr": time.April,
			"jun": time.June,
			"jul": time.July,
			"aug": time.August,
			"sep": time.September, "sept": time.September,
			"okt": time.October,
			"nov": time.November,
			"dec": time.December,
		},
		Weekdays: map[string]time.Weekday{
			// Full names
			"maandag":   time.Monday,
			"dinsdag":   time.Tuesday,
			"woensdag":  time.Wednesday,
			"donderdag": time.Thursday,
			"vrijdag":   time.Friday,
			"zaterdag":  time.Saturday,
			"zondag":    time.Sunday,
			// Abbreviations
			"ma": time.Monday,
			"di": time.Tuesday,
			"wo": time.Wednesday,
			"do": time.Thursday,
			"vr": time.Friday,
			"za": time.Saturday,
			"zo": time.Sunday,
		},
		RelativeTerms: &RelativeTerms{
			Yesterday: "gisteren",
			Today:     "vandaag",
			Tomorrow:  "morgen",
			Now:       "nu",
			// "geleden" for past (2 dagen geleden = 2 days ago)
			Ago: []string{"geleden"},
			// "over" for future (over 3 weken = in 3 weeks)
			In: []string{"over", "in"},
			// "volgende" for next (common gender) and "volgend" (neuter, used with "jaar")
			Next: []string{"volgende", "volgend", "komende", "komend", "aanstaande", "aanstaand"},
			// "vorige" for last (common gender) and "vorig" (neuter, used with "jaar")
			Last: []string{"vorige", "vorig", "afgelopen", "voorgaande", "voorgaand", "laatste"},
			// "deze" for this
			This: []string{"deze", "dit"},
			// Time units with plural forms
			Second:    []string{"seconde", "seconden"},
			Minute:    []string{"minuut", "minuten"},
			Hour:      []string{"uur", "uren"},
			Day:       []string{"dag", "dagen"},
			Week:      []string{"week", "weken"},
			Fortnight: []string{"veertien dagen", "twee weken"},
			Month:     []string{"maand", "maanden"},
			Quarter:   []string{"kwartaal", "kwartalen"},
			Year:      []string{"jaar", "jaren"},
			Decade:    []string{"decennium", "decennia", "tien jaar"},
			// Period boundaries
			Beginning: []string{"begin", "start"},
			End:       []string{"einde", "eind"},
			Start:     []string{"begin", "start"},
			First:     []string{"eerste"},
		},
		TimeTerms: &TimeTerms{
			Noon:     []string{"middag", "twaalf uur 's middags"},
			Midnight: []string{"middernacht", "twaalf uur 's nachts"},
			Quarter:  []string{"kwart", "kwartier"},
			Half:     []string{"half", "halve"},
			// "over" for past (kwart over 3 = quarter past 3)
			Past: []string{"over"},
			// "voor" for to (kwart voor 3 = quarter to 3)
			To:     []string{"voor"},
			OClock: []string{"uur"},
			AM:     []string{"am", "a.m.", "'s ochtends", "'s morgens", "ochtend", "morgen"},
			PM:     []string{"pm", "p.m.", "'s middags", "'s avonds", "'s nachts", "middag", "avond", "nacht"},
		},
	}
}
