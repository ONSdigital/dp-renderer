package model

import "html/template"

type InputType int

const (
	Text InputType = iota
	Email
	Tel
	Url
)

// FuncGetInputType returns the input type as a string
func (i Input) FuncGetInputType() (inputType string) {
	switch i.Type {
	case Text:
		return "text"
	case Email:
		return "email"
	case Tel:
		return "tel"
	case Url:
		return "url"
	}
	return inputType
}

/*
Input represents the common attributes and elements for html input.
Some properties are not rendered if they are invalid attributes for the type.
*/
type Input struct {
	Autocomplete   string          `json:"autocomplete"`    // Renders the autocomplete browser functionality for the input
	DataAttributes []DataAttribute `json:"data_attributes"` // Renders a key/value pair of data attributes automatically prepended with 'data-'
	Description    Localisation    `json:"description"`     // Human readable additional content to support the label
	ID             string          `json:"id"`              // Unique element ID
	IsChecked      bool            `json:"is_checked"`      // Boolean representing whether the element is checked
	IsDisabled     bool            `json:"is_disabled"`     // Boolean representing whether the element is disabled
	IsRequired     bool            `json:"is_required"`     // Boolean representing whether the element is required
	Label          Localisation    `json:"label"`           // Human readable label
	Language       string          `json:"language"`        // Passed from the Page model
	Name           string          `json:"name"`            // Name attribute used for model binding
	Type           InputType       `json:"type"`            // Input type - default 'text'
	Value          string          `json:"value"`           // Value sent to the server
}

// DataAttribute represents the data struct required to display a data attribute
type DataAttribute struct {
	Key   template.HTMLAttr `json:"key"`
	Value Localisation      `json:"value"`
}
