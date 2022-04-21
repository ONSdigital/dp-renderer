package model

// CompactSearch data
type CompactSearch struct {
	ElementId  string       `json:"element_id"`
	InputName  string       `json:"input_name"`
	Language   string       `json:"language"`
	Label      Localisation `json:"label"`
	SearchTerm string       `json:"search_term"`
}
