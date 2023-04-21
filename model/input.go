package model

/*
Input represents the common attributes and elements for html input.
Some properties are not rendered if they are invalid attributes for the type.
*/
type Input struct {
	Autocomplete string       `json:"autocomplete"` // Renders the autocomplete browser functionality for the input
	Description  Localisation `json:"description"`  // Human readable additional content to support the label
	ID           string       `json:"id"`           // Unique element ID
	IsChecked    bool         `json:"is_checked"`   // Boolean representing whether the element is checked
	IsDisabled   bool         `json:"is_disabled"`  // Boolean representing whether the element is disabled
	Label        Localisation `json:"label"`        // Human readable label
	Language     string       `json:"language"`     // Passed from the Page model
	Name         string       `json:"name"`         // Name attribute used for model binding
	Value        string       `json:"value"`        // Value sent to the server
}
