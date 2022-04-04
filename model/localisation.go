package model

type Localisation struct {
	Text      string `json:"text"`
	LocaleKey string `json:"locale_key"`
	Plural    int    `json:"plural"`
}
