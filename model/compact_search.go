package model

/* CompactSearch provides an input for a search term.
ElementId is the id in the document.
InputName is the name submitted in a form.
Language should be passed in from the Page model.
Label states the purpose of the search term.
SearchTerm is the user input submitted in a form.
*/
type CompactSearch struct {
	ElementId  string       `json:"element_id"`
	InputName  string       `json:"input_name"`
	Language   string       `json:"language"`
	Label      Localisation `json:"label"`
	SearchTerm string       `json:"search_term"`
}
