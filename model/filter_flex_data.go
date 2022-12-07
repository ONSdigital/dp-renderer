package model

// FilterFlex ...
type FilterFlex struct {
	Dimensions       []string `json:"dimensions"`
	AreaType         string   `json:"area_type"`
	CoverageAreaType string   `json:"coverage_area_type"`
	Coverage         []string `json:"coverage"`
}
