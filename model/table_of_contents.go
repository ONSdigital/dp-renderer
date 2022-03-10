package model

type ContentSection struct {
	Current bool
	Title   string
}

type TableOfContents struct {
	AriaLabel    string
	Title        string
	Sections     map[string]ContentSection
	DisplayOrder []string
}
