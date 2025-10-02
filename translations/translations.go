// Package translations provides internationalization support for date parsing.
package translations

import (
	"regexp"
	"strings"
	"time"
)

// Language represents a supported language with its translation data.
type Language struct {
	Code             string
	Name             string
	Months           map[string]time.Month
	Weekdays         map[string]time.Weekday
	RelativeTerms    *RelativeTerms
	TimeTerms        *TimeTerms
	RelativePatterns []*LocalizedPattern
}

// RelativeTerms contains localized relative date keywords.
type RelativeTerms struct {
	// Simple terms
	Yesterday string
	Today     string
	Tomorrow  string
	Now       string

	// Directional terms (with variations for gender/number)
	Ago  []string // "ago", "hace"
	In   []string // "in", "en"
	Next []string // "next", "próximo", "próxima"
	Last []string // "last", "último", "última"
	This []string // "this", "este", "esta"

	// Period terms
	Second    []string // "second", "segundo"
	Minute    []string // "minute", "minuto"
	Hour      []string // "hour", "hora"
	Day       []string // "day", "día"
	Week      []string // "week", "semana"
	Fortnight []string // "fortnight", "quincena"
	Month     []string // "month", "mes"
	Quarter   []string // "quarter", "trimestre"
	Year      []string // "year", "año"
	Decade    []string // "decade", "década"

	// Period boundaries
	Beginning []string // "beginning", "inicio", "comienzo"
	End       []string // "end", "final", "fin"
	Start     []string // "start", "inicio"
	First     []string // "first", "primer", "primero"
}

// TimeTerms contains localized time-related keywords.
type TimeTerms struct {
	Noon     []string // "noon", "mediodía"
	Midnight []string // "midnight", "medianoche"
	Quarter  []string // "quarter", "cuarto"
	Half     []string // "half", "media", "medio"
	Past     []string // "past", "y"
	To       []string // "to", "menos"
	OClock   []string // "o'clock", "en punto"
	AM       []string // "am", "de la mañana"
	PM       []string // "pm", "de la tarde", "de la noche"
}

// LocalizedPattern represents a language-specific regex pattern.
type LocalizedPattern struct {
	Name        string
	Regex       *regexp.Regexp
	Example     string
	Description string
}

// Registry holds all registered language translations.
type Registry struct {
	languages map[string]*Language
	defaultVal string
}

// NewRegistry creates a new translation registry.
func NewRegistry() *Registry {
	r := &Registry{
		languages: make(map[string]*Language),
		defaultVal: "en",
	}

	// Register English by default
	r.Register(NewEnglishTranslation())

	return r
}

// Register adds a language to the registry.
func (r *Registry) Register(lang *Language) {
	r.languages[lang.Code] = lang
}

// Get retrieves a language by code.
func (r *Registry) Get(code string) *Language {
	if lang, ok := r.languages[code]; ok {
		return lang
	}
	return r.languages[r.defaultVal]
}

// GetMultiple retrieves multiple languages by codes.
func (r *Registry) GetMultiple(codes []string) []*Language {
	var langs []*Language
	for _, code := range codes {
		if lang, ok := r.languages[code]; ok {
			langs = append(langs, lang)
		}
	}
	if len(langs) == 0 {
		langs = append(langs, r.languages[r.defaultVal])
	}
	return langs
}

// DetectLanguage attempts to detect the language of the input string.
func (r *Registry) DetectLanguage(input string) string {
	input = strings.ToLower(input)

	// Check for language-specific indicators
	scores := make(map[string]int)

	for code, lang := range r.languages {
		score := 0

		// Check months
		for month := range lang.Months {
			if strings.Contains(input, strings.ToLower(month)) {
				score += 10
			}
		}

		// Check weekdays
		for weekday := range lang.Weekdays {
			if strings.Contains(input, strings.ToLower(weekday)) {
				score += 10
			}
		}

		// Check relative terms
		if lang.RelativeTerms != nil {
			terms := [][]string{
				{lang.RelativeTerms.Yesterday, lang.RelativeTerms.Today, lang.RelativeTerms.Tomorrow},
				lang.RelativeTerms.Ago,
				lang.RelativeTerms.In,
				lang.RelativeTerms.Next,
				lang.RelativeTerms.Last,
			}
			for _, termList := range terms {
				for _, term := range termList {
					if term != "" && strings.Contains(input, strings.ToLower(term)) {
						score += 5
					}
				}
			}
		}

		if score > 0 {
			scores[code] = score
		}
	}

	// Return language with highest score
	maxScore := 0
	detectedLang := r.defaultVal
	for code, score := range scores {
		if score > maxScore {
			maxScore = score
			detectedLang = code
		}
	}

	return detectedLang
}

// SupportedLanguages returns a list of all supported language codes.
func (r *Registry) SupportedLanguages() []string {
	codes := make([]string, 0, len(r.languages))
	for code := range r.languages {
		codes = append(codes, code)
	}
	return codes
}
