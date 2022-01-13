package model

/* Collapsible maps the collapsible UI component.
The title text can be either a 'Title' or a 'LocaliseKey', the question mark will always render at the end.
The 'LocaliseKey' has to correspond to the localisation key found in the toml files within assets/locales, otherwise the page will error.
LocalisePluralInt refers to the plural int used in the toml file.
*/
type Collapsible struct {
	Title             string            `json:"title"`
	LocaliseKey       string            `json:"localise_key"`
	LocalisePluralInt int               `json:"localise_plural_int"`
	CollapsibleItems  []CollapsibleItem `json:"collapsible_item"`
}

// CollapsibleItem is an individual representation of the data required in a collapsible item
type CollapsibleItem struct {
	Subheading string   `json:"subheading"`
	Content    []string `json:"content"`
}
