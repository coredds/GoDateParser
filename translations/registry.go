package translations

import "sync"

var (
	// GlobalRegistry is the default translation registry
	GlobalRegistry *Registry
	once           sync.Once
)

// init initializes the global registry with all supported languages
func init() {
	once.Do(func() {
		GlobalRegistry = NewRegistry()

		// Register all supported languages
		GlobalRegistry.Register(NewEnglishTranslation())
		GlobalRegistry.Register(NewSpanishTranslation())
		GlobalRegistry.Register(NewPortugueseTranslation())
		GlobalRegistry.Register(NewFrenchTranslation())
		GlobalRegistry.Register(NewGermanTranslation())
		GlobalRegistry.Register(NewChineseTranslation())
		GlobalRegistry.Register(NewJapaneseTranslation())
	})
}

// GetLanguage is a convenience function to get a language from the global registry
func GetLanguage(code string) *Language {
	return GlobalRegistry.Get(code)
}

// DetectLanguage is a convenience function to detect language from the global registry
func DetectLanguage(input string) string {
	return GlobalRegistry.DetectLanguage(input)
}

// SupportedLanguages returns all supported language codes
func SupportedLanguages() []string {
	return GlobalRegistry.SupportedLanguages()
}
