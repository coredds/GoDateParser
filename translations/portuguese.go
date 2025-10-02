package translations

import (
	"time"
)

// NewPortugueseTranslation creates the Portuguese (Brazil) language translation.
func NewPortugueseTranslation() *Language {
	return &Language{
		Code: "pt",
		Name: "Portuguese",
		Months: map[string]time.Month{
			// Full names
			"janeiro":   time.January,
			"fevereiro": time.February,
			"março":     time.March, "marco": time.March,
			"abril":    time.April,
			"maio":     time.May,
			"junho":    time.June,
			"julho":    time.July,
			"agosto":   time.August,
			"setembro": time.September,
			"outubro":  time.October,
			"novembro": time.November,
			"dezembro": time.December,
			// Abbreviations
			"jan": time.January,
			"fev": time.February,
			"mar": time.March,
			"abr": time.April,
			"mai": time.May,
			"jun": time.June,
			"jul": time.July,
			"ago": time.August,
			"set": time.September,
			"out": time.October,
			"nov": time.November,
			"dez": time.December,
		},
		Weekdays: map[string]time.Weekday{
			// Full names
			"segunda-feira": time.Monday, "segunda": time.Monday,
			"terça-feira": time.Tuesday, "terca-feira": time.Tuesday, "terça": time.Tuesday, "terca": time.Tuesday,
			"quarta-feira": time.Wednesday, "quarta": time.Wednesday,
			"quinta-feira": time.Thursday, "quinta": time.Thursday,
			"sexta-feira": time.Friday, "sexta": time.Friday,
			"sábado": time.Saturday, "sabado": time.Saturday,
			"domingo": time.Sunday,
			// Abbreviations
			"seg": time.Monday,
			"ter": time.Tuesday,
			"qua": time.Wednesday,
			"qui": time.Thursday,
			"sex": time.Friday,
			"sáb": time.Saturday, "sab": time.Saturday,
			"dom": time.Sunday,
		},
		RelativeTerms: &RelativeTerms{
			Yesterday: "ontem",
			Today:     "hoje",
			Tomorrow:  "amanhã",
			Now:       "agora",
			// "atrás" or "há" for past (há 2 dias = 2 days ago)
			Ago: []string{"atrás", "atras", "há", "ha"},
			// "em" or "daqui a" for future (em 3 semanas = in 3 weeks, daqui a 3 dias = 3 days from now)
			In: []string{"em", "daqui a", "daqui"},
			// Gender variations for "next"
			Next: []string{"próximo", "próxima", "proximo", "proxima"},
			// Gender variations for "last"
			Last: []string{"último", "última", "ultimo", "ultima", "passado", "passada"},
			// Gender variations for "this"
			This: []string{"este", "esta", "esse", "essa", "isto", "isso"},
			// Time units with plural forms
			Second:    []string{"segundo", "segundos"},
			Minute:    []string{"minuto", "minutos"},
			Hour:      []string{"hora", "horas"},
			Day:       []string{"dia", "dias"},
			Week:      []string{"semana", "semanas"},
			Fortnight: []string{"quinzena", "quinzenas"},
			Month:     []string{"mês", "meses", "mes"},
			Quarter:   []string{"trimestre", "trimestres"},
			Year:      []string{"ano", "anos"},
			Decade:    []string{"década", "décadas", "decada", "decadas"},
			// Period boundaries
			Beginning: []string{"começo", "comeco", "início", "inicio", "princípio", "principio"},
			End:       []string{"fim", "final"},
			Start:     []string{"início", "inicio", "começo", "comeco"},
			First:     []string{"primeiro", "primeira"},
		},
		TimeTerms: &TimeTerms{
			Noon:     []string{"meio-dia", "meio dia", "meiodia"},
			Midnight: []string{"meia-noite", "meia noite", "meianoite"},
			Quarter:  []string{"quarto", "um quarto", "uma quarto", "quinze"},
			Half:     []string{"meia", "meio", "trinta"},
			// "e" for past (3 e meia = half past 3)
			Past: []string{"e"},
			// "para" or "menos" for to (quinze para as 3 = quarter to 3, menos quinze = minus 15)
			To:     []string{"para", "menos"},
			OClock: []string{"em ponto", "horas"},
			AM:     []string{"am", "a.m.", "da manhã", "da manha", "de manhã", "de manha"},
			PM:     []string{"pm", "p.m.", "da tarde", "de tarde", "da noite", "de noite"},
		},
	}
}
