package model

/* InputDate provides an input for a date.
Language should be passed in from the Page model.
Id is the id in the document.
InputNameDay is the name submitted for the day in a form.
InputNameMonth is the name submitted for the month in a form.
InputNameYear is the name submitted for the year in a form.
InputValueDay is the user input submitted for the day in a form.
InputValueMonth is the user input submitted for the month in a form.
InputValueYear is the user input submitted for the year in a form.
Title states the purpose of the date.
Description offers a further explanation for the purpose of the date.
*/
type InputDate struct {
	Language        string       `json:"language"`
	Id              string       `json:"id"`
	InputNameDay    string       `json:"input_name_day"`
	InputNameMonth  string       `json:"input_name_month"`
	InputNameYear   string       `json:"input_name_year"`
	InputValueDay   string       `json:"input_value_day"`
	InputValueMonth string       `json:"input_value_month"`
	InputValueYear  string       `json:"input_value_year"`
	Title           Localisation `json:"title"`
	Description     Localisation `json:"description"`
}
