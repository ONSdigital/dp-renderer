package model

// DateFieldset defines the HTML for a date input
type DateFieldset struct {
	ErrorID                  string         `json:"error_id"`                   // HTML id attribute used within a field validation error
	Input                    InputDate      `json:"input"`                      // HTML date input attributes
	Language                 string         `json:"language"`                   // Passed in from the Page model
	ValidationErrDescription []Localisation `json:"validation_err_description"` // String array that describes the validation error
}

// FuncHasDateValidationErr helper function that returns true if any of the date input fields have a validation error
func (d DateFieldset) FuncHasDateValidationErr() bool {
	return d.Input.HasYearValidationErr || d.Input.HasMonthValidationErr || d.Input.HasDayValidationErr
}
