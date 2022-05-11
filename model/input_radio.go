package model

/* InputRadioBox provides an input for a checkbox.
Language should be passed in from the Page model.
Id is the id in the document.
Name is the name in the document.
Checked is flag that shows up it the value it true
Title states the purpose of the InputCheckbox.
*/
type InputRadioBox struct {
	Language  string       `json:"language"`
	Id        string       `json:"id"`
	Name      string       `json:"name"`
	Value     string       `json:"value"`
	Checked   bool         `json:"checked"`
	IsLabel   bool         `json:"isLabel"`
	Count     string       `json:"count"`
	Title     Localisation `json:"title"`
	LocaleKey Localisation `json:"localeKey"`
}
