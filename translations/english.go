package translations

import (
	"time"
)

// NewEnglishTranslation creates the English language translation.
func NewEnglishTranslation() *Language {
	return &Language{
		Code: "en",
		Name: "English",
		Months: map[string]time.Month{
			"january": time.January, "jan": time.January,
			"february": time.February, "feb": time.February,
			"march": time.March, "mar": time.March,
			"april": time.April, "apr": time.April,
			"may":  time.May,
			"june": time.June, "jun": time.June,
			"july": time.July, "jul": time.July,
			"august": time.August, "aug": time.August,
			"september": time.September, "sep": time.September, "sept": time.September,
			"october": time.October, "oct": time.October,
			"november": time.November, "nov": time.November,
			"december": time.December, "dec": time.December,
		},
		Weekdays: map[string]time.Weekday{
			"monday": time.Monday, "mon": time.Monday,
			"tuesday": time.Tuesday, "tue": time.Tuesday, "tues": time.Tuesday,
			"wednesday": time.Wednesday, "wed": time.Wednesday,
			"thursday": time.Thursday, "thu": time.Thursday, "thur": time.Thursday, "thurs": time.Thursday,
			"friday": time.Friday, "fri": time.Friday,
			"saturday": time.Saturday, "sat": time.Saturday,
			"sunday": time.Sunday, "sun": time.Sunday,
		},
		RelativeTerms: &RelativeTerms{
			Yesterday: "yesterday",
			Today:     "today",
			Tomorrow:  "tomorrow",
			Now:       "now",
			Ago:       []string{"ago"},
			In:        []string{"in"},
			Next:      []string{"next"},
			Last:      []string{"last"},
			This:      []string{"this"},
			Second:    []string{"second", "seconds"},
			Minute:    []string{"minute", "minutes"},
			Hour:      []string{"hour", "hours"},
			Day:       []string{"day", "days"},
			Week:      []string{"week", "weeks"},
			Fortnight: []string{"fortnight", "fortnights"},
			Month:     []string{"month", "months"},
			Quarter:   []string{"quarter", "quarters"},
			Year:      []string{"year", "years"},
			Decade:    []string{"decade", "decades"},
			Beginning: []string{"beginning", "start"},
			End:       []string{"end"},
			Start:     []string{"start"},
			First:     []string{"first"},
		},
		TimeTerms: &TimeTerms{
			Noon:     []string{"noon"},
			Midnight: []string{"midnight"},
			Quarter:  []string{"quarter"},
			Half:     []string{"half"},
			Past:     []string{"past", "after"},
			To:       []string{"to", "before"},
			OClock:   []string{"o'clock"},
			AM:       []string{"am", "a.m."},
			PM:       []string{"pm", "p.m."},
		},
	}
}
