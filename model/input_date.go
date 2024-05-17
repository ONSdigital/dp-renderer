package model

// TODO: There is significant duplication on model props, refactor for v3

/*
InputDate provides an input for a date.
Language should be passed in from the Page model.
Id is the id in the document.
InputNameDay is the name submitted for the day in a form.
InputNameMonth is the name submitted for the month in a form.
InputNameYear is the name submitted for the year in a form.
InputValueDay is the user input submitted for the day in a form.
InputValueMonth is the user input submitted for the month in a form.
InputValueYear is the user input submitted for the year in a form.
HasDayValidationErr displays a validation error for the day input field.
HasMonthValidationErr displays a validation error for the month input field.
HasYearValidationErr displays a validation error for the year input field.
Title states the purpose of the date.
Description offers a further explanation for the purpose of the date.
Renders a key/value pair of data attributes automatically prepended with 'data-'
*/
type InputDate struct {
	Language              string          `json:"language"`
	Id                    string          `json:"id"`
	InputNameDay          string          `json:"input_name_day"`
	InputNameMonth        string          `json:"input_name_month"`
	InputNameYear         string          `json:"input_name_year"`
	InputValueDay         string          `json:"input_value_day"`
	InputValueMonth       string          `json:"input_value_month"`
	InputValueYear        string          `json:"input_value_year"`
	HasDayValidationErr   bool            `json:"has_day_validation_err"`
	HasMonthValidationErr bool            `json:"has_month_validation_err"`
	HasYearValidationErr  bool            `json:"has_year_validation_err"`
	Title                 Localisation    `json:"title"`
	Description           Localisation    `json:"description"`
	DataAttributes        []DataAttribute `json:"data_attributes"`
	DayDataAttributes     []DataAttribute `json:"day_data_attributes"`
	MonthDataAttributes   []DataAttribute `json:"month_data_attributes"`
	YearDataAttributes    []DataAttribute `json:"year_data_attributes"`
}
