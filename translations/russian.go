package translations

import (
	"time"
)

// NewRussianTranslation creates the Russian language translation.
func NewRussianTranslation() *Language {
	return &Language{
		Code: "ru",
		Name: "Russian",
		Months: map[string]time.Month{
			// Full names (nominative case)
			"январь":   time.January,
			"февраль":  time.February,
			"март":     time.March,
			"апрель":   time.April,
			"май":      time.May,
			"июнь":     time.June,
			"июль":     time.July,
			"август":   time.August,
			"сентябрь": time.September,
			"октябрь":  time.October,
			"ноябрь":   time.November,
			"декабрь":  time.December,
			// Genitive case (used with dates: "15 января")
			"января":   time.January,
			"февраля":  time.February,
			"марта":    time.March,
			"апреля":   time.April,
			"мая":      time.May,
			"июня":     time.June,
			"июля":     time.July,
			"августа":  time.August,
			"сентября": time.September,
			"октября":  time.October,
			"ноября":   time.November,
			"декабря":  time.December,
			// Abbreviations
			"янв": time.January,
			"фев": time.February,
			"мар": time.March,
			"апр": time.April,
			"июн": time.June,
			"июл": time.July,
			"авг": time.August,
			"сен": time.September, "сент": time.September,
			"окт": time.October,
			"ноя": time.November,
			"дек": time.December,
		},
		Weekdays: map[string]time.Weekday{
			// Full names (nominative case)
			"понедельник": time.Monday,
			"вторник":     time.Tuesday,
			"среда":       time.Wednesday,
			"четверг":     time.Thursday,
			"пятница":     time.Friday,
			"суббота":     time.Saturday,
			"воскресенье": time.Sunday,
			// Prepositional case (used with "в": "в понедельнике")
			"понедельнике":   time.Monday,
			"вторнике":       time.Tuesday,
			"среде":          time.Wednesday,
			"четверге":       time.Thursday,
			"пятнице":        time.Friday,
			"субботе":        time.Saturday,
			"в понедельнике": time.Monday,
			"в вторнике":     time.Tuesday,
			"в среде":        time.Wednesday,
			"в четверге":     time.Thursday,
			"в пятнице":      time.Friday,
			"в субботе":      time.Saturday,
			// Note: воскресенье is the same in prepositional case, already included above
			// Abbreviations
			"пн": time.Monday,
			"вт": time.Tuesday,
			"ср": time.Wednesday,
			"чт": time.Thursday,
			"пт": time.Friday,
			"сб": time.Saturday,
			"вс": time.Sunday,
		},
		RelativeTerms: &RelativeTerms{
			Yesterday: "вчера",
			Today:     "сегодня",
			Tomorrow:  "завтра",
			Now:       "сейчас",
			// "назад" for past (2 дня назад = 2 days ago)
			Ago: []string{"назад", "тому назад"},
			// "через" for future (через 3 недели = in 3 weeks)
			In: []string{"через", "спустя"},
			// Gender/number variations for "next" (nominative and genitive cases)
			Next: []string{"следующий", "следующая", "следующее", "следующие", "следующего", "следующей", "следующих", "будущий", "будущая", "будущее", "будущие", "будущего", "будущей", "будущих"},
			// Gender/number variations for "last" (nominative and genitive cases)
			Last: []string{"прошлый", "прошлая", "прошлое", "прошлые", "прошлого", "прошлой", "прошлых", "последний", "последняя", "последнее", "последние", "последнего", "последней", "последних", "предыдущий", "предыдущая", "предыдущее", "предыдущие", "предыдущего", "предыдущей", "предыдущих"},
			// Gender/number variations for "this"
			This: []string{"этот", "эта", "это", "эти", "текущий", "текущая", "текущее", "текущие"},
			// Time units with plural forms (nominative, genitive singular, genitive plural, accusative)
			Second:    []string{"секунда", "секунды", "секунд", "секунду"},
			Minute:    []string{"минута", "минуты", "минут", "минуту"},
			Hour:      []string{"час", "часа", "часов"},
			Day:       []string{"день", "дня", "дней"},
			Week:      []string{"неделя", "недели", "недель", "неделю"},
			Fortnight: []string{"две недели", "двух недель"},
			Month:     []string{"месяц", "месяца", "месяцев"},
			Quarter:   []string{"квартал", "квартала", "кварталов"},
			Year:      []string{"год", "года", "лет"},
			Decade:    []string{"десятилетие", "десятилетия", "десятилетий", "декада", "декады", "декад"},
			// Period boundaries
			Beginning: []string{"начало", "начала"},
			End:       []string{"конец", "конца"},
			Start:     []string{"начало", "начала"},
			First:     []string{"первый", "первая", "первое", "первые"},
		},
		TimeTerms: &TimeTerms{
			Noon:     []string{"полдень", "полудень", "12 часов дня"},
			Midnight: []string{"полночь", "полуночь", "12 часов ночи", "0 часов"},
			Quarter:  []string{"четверть"},
			Half:     []string{"половина", "пол", "полчаса"},
			// No direct equivalent to "past" in Russian time expressions
			Past: []string{},
			// No direct equivalent to "to" in Russian time expressions
			To:     []string{"без"},
			OClock: []string{"часов", "час", "часа"},
			AM:     []string{"утра", "ночи"},
			PM:     []string{"дня", "вечера"},
		},
	}
}
