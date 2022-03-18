package model

type CompactSearch struct {
	ElementId        string `json:"element_id"`
	InputName        string `json:"input_name"`
	Language         string `json:"language"`
	LabelLocaliseKey string `json:"label_localise_key"`
	Label            string `json:"label"`
	SearchTerm       string `json:"search_term"`
}
