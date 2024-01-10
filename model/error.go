package model

// N.B. This model is shared between 'error.tmpl' and 'partials/error-summary.tmpl'

// Error contains data to display a client page error
type Error struct {
	Title       string      `json:"title"`       // The error title and populates the <title> element
	Description string      `json:"description"` // Free text to describe the error
	ErrorItems  []ErrorItem `json:"error_items"` // Array of error item
	Language    string      `json:"language"`    // User defined language
	ErrorCode   int         `json:"error_code"`  // Http error code e.g. 401, 404, 500
}

// ErrorItem represents an error item.
type ErrorItem struct {
	Description Localisation `json:"description"` // Can be either a 'Localisation.Text' or a 'Localisation.LocaleKey'
	Language    string       `json:"language"`    // User defined language
	ID          string       `json:"id"`          // HTML id attribute used within a field validation error
	URL         string       `json:"url"`         // The href to the error
}
