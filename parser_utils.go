package godateparser

import (
	"time"

	"github.com/coredds/GoDateParser/translations"
)

// monthNameToNumberWithLangs converts a month name using specific languages.
func monthNameToNumberWithLangs(name string, langs []*translations.Language) time.Month {
	if month, ok := translations.ParseMonth(name, langs...); ok {
		return month
	}
	return 0
}
