package model

type InputDate struct {
	Language        string `json:"language"`
	Id              string `json:"id"`
	InputNameDay    string `json:"input_name_day"`
	InputNameMonth  string `json:"input_name_month"`
	InputNameYear   string `json:"input_name_year"`
	InputValueDay   string `json:"input_value_day"`
	InputValueMonth string `json:"input_value_month"`
	InputValueYear  string `json:"input_value_year"`
	Title           string `json:"title"`
	Description     string `json:"description"`
}
