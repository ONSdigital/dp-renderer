package model

/* InputCheckbox provides an input for a checkbox.
Language should be passed in from the Page model.
Id is the id in the document.
Name is the name in the document.
Checked is flag that shows up it the value it true
Title states the purpose of the InputCheckbox.
*/
type InputCheckBox struct {
	Language    string       `json:"language"`
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	Checked     bool         `json:"checked"`
	Count       string       `json:"count"`
	Title       Localisation `json:"title"`
	LocaleKey   Localisation `json:"localeKey"`
	Description Localisation `json:"description"`
}
