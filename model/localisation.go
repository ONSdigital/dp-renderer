package model

import (
	"github.com/ONSdigital/dp-renderer/v2/helper"
)

/*
Localisation offers text or a localised substitute.
Text is displayed as-is.
LocaleKey is a key into the toml files found in assets/locales.
Plural chooses the singular or plural form of the phrase selected by LocaleKey.
*/
type Localisation struct {
	Text      string `json:"text"`
	LocaleKey string `json:"locale_key"`
	Plural    int    `json:"plural"`
}

// Localise by preference when a LocaleKey is provided, defaulting to Text otherwise.
func (localisation Localisation) FuncLocalise(language string) string {
	var result string

	if localisation.LocaleKey != "" {
		result = helper.Localise(localisation.LocaleKey, language, localisation.Plural)
	} else {
		result = localisation.Text
	}

	return result
}
