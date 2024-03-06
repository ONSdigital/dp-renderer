package model

// Radio defines the fields for a radio input
type Radio struct {
	Input      Input `json:"input"`       // HTML input attributes
	OtherInput Input `json:"other_input"` // Conditionally displays an additional text input
}

// RadioFieldset defines the HTML fieldset for radio inputs
type RadioFieldset struct {
	Language      string        `json:"language"`       // Passed in from the Page model
	Legend        Localisation  `json:"legend"`         // Content within legend html element
	Radios        []Radio       `json:"radios"`         // Radios within the fieldset
	ValidationErr ValidationErr `json:"validation_err"` // Fields for validation model against the fieldset
}

// TextareaField defines the fields for a textarea element
type TextareaField struct {
	Input         Input         `json:"input"`          // HTML input attributes
	ValidationErr ValidationErr `json:"validation_err"` // fields for validation model against the field
}

// TextField defines the fields for a text input element
type TextField struct {
	Input         Input         `json:"input"`          // HTML input attributes
	ValidationErr ValidationErr `json:"validation_err"` // fields for validation model against the field
}

// ValidationErr defines the fields for a field validation error
type ValidationErr struct {
	ErrorItem        ErrorItem `json:"error_item"`         // Fields for an error item
	HasValidationErr bool      `json:"has_validation_err"` // Bool check to display additional html required for a field error
}
