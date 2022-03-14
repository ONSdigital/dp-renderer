package model

type ContentSection struct {
	Current bool   `json:"current"`
	Title   string `json:"title"`
}

type TableOfContents struct {
	AriaLabelLocaliseKey string                    `json:"aria_label_localise_key"`
	AriaLabel            string                    `json:"aria_label"`
	TitleLocaliseKey     string                    `json:"title_localise_key"`
	Title                string                    `json:"title"`
	Sections             map[string]ContentSection `json:"sections"`
	DisplayOrder         []string                  `json:"display_order"`
}
