package translations

import (
	"time"
)

// NewChineseTranslation creates the Chinese Simplified (China) language translation.
func NewChineseTranslation() *Language {
	return &Language{
		Code: "zh",
		Name: "Chinese",
		Months: map[string]time.Month{
			// Full names (numeric + 月)
			"一月": time.January, "1月": time.January,
			"二月": time.February, "2月": time.February,
			"三月": time.March, "3月": time.March,
			"四月": time.April, "4月": time.April,
			"五月": time.May, "5月": time.May,
			"六月": time.June, "6月": time.June,
			"七月": time.July, "7月": time.July,
			"八月": time.August, "8月": time.August,
			"九月": time.September, "9月": time.September,
			"十月": time.October, "10月": time.October,
			"十一月": time.November, "11月": time.November,
			"十二月": time.December, "12月": time.December,
		},
		Weekdays: map[string]time.Weekday{
			// Full names
			"星期一": time.Monday, "周一": time.Monday, "礼拜一": time.Monday,
			"星期二": time.Tuesday, "周二": time.Tuesday, "礼拜二": time.Tuesday,
			"星期三": time.Wednesday, "周三": time.Wednesday, "礼拜三": time.Wednesday,
			"星期四": time.Thursday, "周四": time.Thursday, "礼拜四": time.Thursday,
			"星期五": time.Friday, "周五": time.Friday, "礼拜五": time.Friday,
			"星期六": time.Saturday, "周六": time.Saturday, "礼拜六": time.Saturday,
			"星期日": time.Sunday, "星期天": time.Sunday, "周日": time.Sunday, "周天": time.Sunday, "礼拜日": time.Sunday, "礼拜天": time.Sunday,
		},
		RelativeTerms: &RelativeTerms{
			Yesterday: "昨天",
			Today:     "今天",
			Tomorrow:  "明天",
			Now:       "现在",
			// "前" for past (3天前 = 3 days ago)
			Ago: []string{"前", "之前"},
			// "后" for future (3天后 = in 3 days)
			In: []string{"后", "之后", "内"},
			// "下" for next
			Next: []string{"下", "下个", "下一个"},
			// "上" for last
			Last: []string{"上", "上个", "上一个"},
			// "这" for this
			This: []string{"这", "这个", "本"},
			// Time units
			Second:    []string{"秒", "秒钟"},
			Minute:    []string{"分钟", "分"},
			Hour:      []string{"小时", "个小时", "钟头"},
			Day:       []string{"天", "日"},
			Week:      []string{"周", "星期", "礼拜"},
			Fortnight: []string{"两周", "两星期"},
			Month:     []string{"月", "个月"},
			Quarter:   []string{"季度", "季"},
			Year:      []string{"年"},
			Decade:    []string{"十年"},
			// Period boundaries
			Beginning: []string{"初", "开始", "始"},
			End:       []string{"末", "底", "尾", "结束"},
			Start:     []string{"初", "开始"},
			First:     []string{"第一"},
		},
		TimeTerms: &TimeTerms{
			Noon:     []string{"中午", "正午"},
			Midnight: []string{"午夜", "半夜", "凌晨"},
			Quarter:  []string{"一刻", "刻"},
			Half:     []string{"半"},
			Past:     []string{},    // Chinese doesn't use "past" in the same way
			To:       []string{"差"}, // 差10分3点 = 10 minutes to 3
			OClock:   []string{"点", "点钟"},
			AM:       []string{"上午", "早上", "凌晨"},
			PM:       []string{"下午", "晚上", "傍晚"},
		},
	}
}
