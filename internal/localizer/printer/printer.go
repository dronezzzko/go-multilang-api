package printer

import (
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	_ "github.com/dronezzzko/go-multilang-api/internal/translation"
)

const defaultLocaleID = "en-US"

type (
	// internalLocaleID is service internal language ID.
	internalLocaleID string
	// localeID is BCP 47 language ID.
	localeID string
)

var locales = map[internalLocaleID]localeID{
	"en-us": "en-US",
	"fr-fr": "fr-FR",
}

// Make instantiates a new Localizer for provided locale ID.
// If ID is not supported a localizer for default locale (defaultLocaleID) will be returned.
func Make(ID string) (*message.Printer, bool) {
	normalizedID := internalLocaleID(strings.ToLower(ID))
	if langID, ok := locales[normalizedID]; ok {
		return message.NewPrinter(language.MustParse(string(langID))), true
	}

	return message.NewPrinter(language.MustParse(defaultLocaleID)), false
}
