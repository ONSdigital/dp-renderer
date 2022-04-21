package model

/* Localisation offers text or a localised substitute.
Text is displayed as-is.
LocaleKey is a key into the toml files found in assets/locales.
Plural chooses the singular or plural form of the phrase selected by LocaleKey.
*/
type Localisation struct {
	Text      string `json:"text"`
	LocaleKey string `json:"locale_key"`
	Plural    int    `json:"plural"`
}
