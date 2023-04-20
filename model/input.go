package model

/*
Input represents the common attributes for html input elements.
Some properties are not rendered if they are invalid attributes for the type.
Description is human readable additional content to support the label.
ID is the unique element ID.
IsChecked is a boolean representing whether the element is checked.
IsDisabled is a boolean representing whether the element is disabled.
Label is the human readable label.
Language should be passed in from the Page model.
Name is the name attribute used for model binding.
Value is the value sent to the server.
*/
type Input struct {
	Description Localisation `json:"description"`
	ID          string       `json:"id"`
	IsChecked   bool         `json:"is_checked"`
	IsDisabled  bool         `json:"is_disabled"`
	Label       Localisation `json:"label"`
	Language    string       `json:"language"`
	Name        string       `json:"name"`
	Value       string       `json:"value"`
}
