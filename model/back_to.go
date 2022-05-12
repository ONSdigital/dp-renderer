package model

/* BackTo maps the back to component.
AnchorFragment refers to the anchor fragment on the page to link to, leave empty to display the default '#'.
Text refers to the display text that can be either a 'Localisation.Text' or a 'Localisation.LocaleKey'.
*/
type BackTo struct {
	AnchorFragment string       `json:"anchor_fragment"`
	Text           Localisation `json:"text"`
}
