package translations

import (
	"time"
)

// NewFrenchTranslation creates the French (France) language translation.
func NewFrenchTranslation() *Language {
	return &Language{
		Code: "fr",
		Name: "French",
		Months: map[string]time.Month{
			// Full names
			"janvier": time.January,
			"février": time.February, "fevrier": time.February,
			"mars":    time.March,
			"avril":   time.April,
			"mai":     time.May,
			"juin":    time.June,
			"juillet": time.July,
			"août":    time.August, "aout": time.August,
			"septembre": time.September,
			"octobre":   time.October,
			"novembre":  time.November,
			"décembre":  time.December, "decembre": time.December,
			// Abbreviations
			"janv": time.January,
			"févr": time.February, "fevr": time.February,
			"avr":  time.April,
			"juil": time.July,
			"sept": time.September,
			"oct":  time.October,
			"nov":  time.November,
			"déc":  time.December, "dec": time.December,
		},
		Weekdays: map[string]time.Weekday{
			// Full names
			"lundi":    time.Monday,
			"mardi":    time.Tuesday,
			"mercredi": time.Wednesday,
			"jeudi":    time.Thursday,
			"vendredi": time.Friday,
			"samedi":   time.Saturday,
			"dimanche": time.Sunday,
			// Abbreviations
			"lun": time.Monday,
			"mar": time.Tuesday,
			"mer": time.Wednesday,
			"jeu": time.Thursday,
			"ven": time.Friday,
			"sam": time.Saturday,
			"dim": time.Sunday,
		},
		RelativeTerms: &RelativeTerms{
			Yesterday: "hier",
			Today:     "aujourd'hui",
			Tomorrow:  "demain",
			Now:       "maintenant",
			Ago:       []string{"il y a"},
			In:        []string{"dans", "en"},
			Last:      []string{"dernier", "dernière", "derniere"},
			Next:      []string{"prochain", "prochaine"},
			This:      []string{"ce", "cet", "cette"},
			Second:    []string{"seconde", "secondes"},
			Minute:    []string{"minute", "minutes"},
			Hour:      []string{"heure", "heures"},
			Day:       []string{"jour", "jours"},
			Week:      []string{"semaine", "semaines"},
			Fortnight: []string{"quinzaine", "quinzaines"},
			Month:     []string{"mois"},
			Quarter:   []string{"trimestre", "trimestres"},
			Year:      []string{"an", "ans", "année", "années", "annee", "annees"},
			Decade:    []string{"décennie", "décennies", "decennie", "decennies"},
			Beginning: []string{"début", "debut", "commencement"},
			End:       []string{"fin"},
			Start:     []string{"début", "debut"},
			First:     []string{"premier", "première", "premiere"},
		},
		TimeTerms: &TimeTerms{
			Noon:     []string{"midi"},
			Midnight: []string{"minuit"},
			Quarter:  []string{"quart"},
			Half:     []string{"demi", "demie"},
			Past:     []string{"et"},
			To:       []string{"moins"},
			OClock:   []string{"heure", "heures"},
			AM:       []string{"du matin", "matin"},
			PM:       []string{"de l'après-midi", "après-midi", "apres-midi", "du soir", "soir"},
		},
	}
}
