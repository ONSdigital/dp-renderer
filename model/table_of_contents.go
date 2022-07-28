package model

/* ContentSection maps the content details.
The visible text can be either a 'Localisation.Text' or a 'Localisation.LocaleKey'.
The 'LocaleKey' has to correspond to the localisation key found in the toml files within assets/locales, otherwise the page will error.
Plural refers to the plural int used in the toml file.
*/
type ContentSection struct {
	Current bool         `json:"current"`
	Title   Localisation `json:"title"`
}

// TableOfContents contains the contents of the page
type TableOfContents struct {
	Id           string                    `json:"id"`
	AriaLabel    Localisation              `json:"aria_label"`
	Title        Localisation              `json:"title"`
	Sections     map[string]ContentSection `json:"sections"`
	DisplayOrder []string                  `json:"display_order"`
}
