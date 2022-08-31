package model

// N.B. This model is shared between 'error.tmpl' and 'partials/error-summary.tmpl'

/* Error contains data to display a client page error
Title is the error title and populates the <title> element
Description is free text
ErrorItems is an array of page errors
Language is the user defined language
*/
type Error struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	ErrorItems  []ErrorItem `json:"error_items"`
	Language    string      `json:"language"`
}

/* ErrorItem represents an error item.
The description can be either a 'Localisation.Text' or a 'Localisation.LocaleKey'.
The 'LocaleKey' has to correspond to the localisation key found in the toml files within assets/locales, otherwise the page will error.
Plural refers to the plural int used in the toml file.
URL is the href to the error
*/
type ErrorItem struct {
	Description Localisation `json:"description"`
	URL         string       `json:"url"`
}
