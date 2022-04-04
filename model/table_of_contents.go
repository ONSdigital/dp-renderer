package model

type ContentSection struct {
	Current bool   `json:"current"`
	Title   string `json:"title"`
}

type TableOfContents struct {
	AriaLabel    Localisation              `json:"aria_label"`
	Title        Localisation              `json:"title"`
	Sections     map[string]ContentSection `json:"sections"`
	DisplayOrder []string                  `json:"display_order"`
}
