package translations

import (
	"time"
)

// NewSpanishTranslation creates the Spanish language translation.
func NewSpanishTranslation() *Language {
	return &Language{
		Code: "es",
		Name: "Spanish",
		Months: map[string]time.Month{
			// Full names
			"enero":      time.January,
			"febrero":    time.February,
			"marzo":      time.March,
			"abril":      time.April,
			"mayo":       time.May,
			"junio":      time.June,
			"julio":      time.July,
			"agosto":     time.August,
			"septiembre": time.September, "setiembre": time.September,
			"octubre":   time.October,
			"noviembre": time.November,
			"diciembre": time.December,
			// Abbreviations
			"ene": time.January,
			"feb": time.February,
			"mar": time.March,
			"abr": time.April,
			"may": time.May,
			"jun": time.June,
			"jul": time.July,
			"ago": time.August,
			"sep": time.September, "set": time.September,
			"oct": time.October,
			"nov": time.November,
			"dic": time.December,
		},
		Weekdays: map[string]time.Weekday{
			// Full names
			"lunes":     time.Monday,
			"martes":    time.Tuesday,
			"miércoles": time.Wednesday, "miercoles": time.Wednesday,
			"jueves":  time.Thursday,
			"viernes": time.Friday,
			"sábado":  time.Saturday, "sabado": time.Saturday,
			"domingo": time.Sunday,
			// Abbreviations
			"lun": time.Monday,
			"mar": time.Tuesday,
			"mié": time.Wednesday, "mie": time.Wednesday,
			"jue": time.Thursday,
			"vie": time.Friday,
			"sáb": time.Saturday, "sab": time.Saturday,
			"dom": time.Sunday,
		},
		RelativeTerms: &RelativeTerms{
			Yesterday: "ayer",
			Today:     "hoy",
			Tomorrow:  "mañana",
			Now:       "ahora",
			// "hace" for past (hace 2 días = 2 days ago)
			Ago: []string{"hace"},
			// "en" for future (en 3 semanas = in 3 weeks)
			In: []string{"en", "dentro de"},
			// Gender variations for "next"
			Next: []string{"próximo", "próxima", "proximo", "proxima", "siguiente"},
			// Gender variations for "last"
			Last: []string{"último", "última", "ultimo", "ultima", "pasado", "pasada"},
			// Gender variations for "this"
			This: []string{"este", "esta", "esto"},
			// Time units with plural forms
			Second:    []string{"segundo", "segundos"},
			Minute:    []string{"minuto", "minutos"},
			Hour:      []string{"hora", "horas"},
			Day:       []string{"día", "días", "dia", "dias"},
			Week:      []string{"semana", "semanas"},
			Fortnight: []string{"quincena", "quincenas"},
			Month:     []string{"mes", "meses"},
			Quarter:   []string{"trimestre", "trimestres"},
			Year:      []string{"año", "años", "ano", "anos"},
			Decade:    []string{"década", "décadas", "decada", "decadas"},
			// Period boundaries
			Beginning: []string{"comienzo", "inicio", "principio"},
			End:       []string{"fin", "final"},
			Start:     []string{"inicio", "comienzo"},
			First:     []string{"primer", "primero", "primera"},
		},
		TimeTerms: &TimeTerms{
			Noon:     []string{"mediodía", "mediodia", "medio día", "medio dia"},
			Midnight: []string{"medianoche", "media noche"},
			Quarter:  []string{"cuarto"},
			Half:     []string{"media", "medio"},
			// "y" for past (3 y cuarto = quarter past 3)
			Past: []string{"y"},
			// "menos" for to (menos cuarto = quarter to)
			To:     []string{"menos", "para"},
			OClock: []string{"en punto"},
			AM:     []string{"am", "a.m.", "de la mañana", "de la manana"},
			PM:     []string{"pm", "p.m.", "de la tarde", "de la noche"},
		},
	}
}
