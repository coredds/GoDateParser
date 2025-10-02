package translations

import (
	"time"
)

// NewJapaneseTranslation creates the Japanese (Japan) language translation.
func NewJapaneseTranslation() *Language {
	return &Language{
		Code: "ja",
		Name: "Japanese",
		Months: map[string]time.Month{
			// Full names with 月 (gatsu)
			"一月": time.January, "1月": time.January, "1がつ": time.January,
			"二月": time.February, "2月": time.February, "2がつ": time.February,
			"三月": time.March, "3月": time.March, "3がつ": time.March,
			"四月": time.April, "4月": time.April, "4がつ": time.April,
			"五月": time.May, "5月": time.May, "5がつ": time.May,
			"六月": time.June, "6月": time.June, "6がつ": time.June,
			"七月": time.July, "7月": time.July, "7がつ": time.July,
			"八月": time.August, "8月": time.August, "8がつ": time.August,
			"九月": time.September, "9月": time.September, "9がつ": time.September,
			"十月": time.October, "10月": time.October, "10がつ": time.October,
			"十一月": time.November, "11月": time.November, "11がつ": time.November,
			"十二月": time.December, "12月": time.December, "12がつ": time.December,
		},
		Weekdays: map[string]time.Weekday{
			// Full names with 曜日 (youbi)
			"月曜日": time.Monday, "月曜": time.Monday, "げつようび": time.Monday,
			"火曜日": time.Tuesday, "火曜": time.Tuesday, "かようび": time.Tuesday,
			"水曜日": time.Wednesday, "水曜": time.Wednesday, "すいようび": time.Wednesday,
			"木曜日": time.Thursday, "木曜": time.Thursday, "もくようび": time.Thursday,
			"金曜日": time.Friday, "金曜": time.Friday, "きんようび": time.Friday,
			"土曜日": time.Saturday, "土曜": time.Saturday, "どようび": time.Saturday,
			"日曜日": time.Sunday, "日曜": time.Sunday, "にちようび": time.Sunday,
		},
		RelativeTerms: &RelativeTerms{
			Yesterday: "昨日",                    // kinou
			Today:     "今日",                    // kyou
			Tomorrow:  "明日",                    // ashita/asu
			Now:       "今",                     // ima
			Ago:       []string{"前"},           // mae (e.g., 3日前 = 3 days ago)
			In:        []string{"後", "あと"},     // ato/go (e.g., 3日後 = in 3 days)
			Next:      []string{"来", "次", "翌"}, // rai/tsugi/yoku
			Last:      []string{"先", "前", "昨"}, // sen/mae/saku
			This:      []string{"今", "本"},      // kon/hon
			// Time units
			Second:    []string{"秒", "秒間", "びょう"},
			Minute:    []string{"分", "分間", "ふん"},
			Hour:      []string{"時間", "じかん"},
			Day:       []string{"日", "日間", "にち"},
			Week:      []string{"週", "週間", "しゅう"},
			Fortnight: []string{"二週間", "2週間"},
			Month:     []string{"月", "ヶ月", "か月", "ケ月"},
			Quarter:   []string{"四半期", "クォーター"},
			Year:      []string{"年", "年間", "ねん"},
			Decade:    []string{"十年", "10年"},
			// Period boundaries
			Beginning: []string{"初", "始", "頭"},
			End:       []string{"末", "終", "終わり"},
			Start:     []string{"初", "始め"},
			First:     []string{"初", "最初"},
		},
		TimeTerms: &TimeTerms{
			Noon:     []string{"正午", "昼", "12時"},
			Midnight: []string{"真夜中", "夜中", "0時"},
			Quarter:  []string{"15分"},
			Half:     []string{"半", "30分"},
			Past:     []string{"過ぎ"},
			To:       []string{"前"},
			OClock:   []string{"時"},
			AM:       []string{"午前", "朝"},
			PM:       []string{"午後", "夜"},
		},
	}
}
