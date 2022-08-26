package model

// Figure represents a figure (charts, tables)
type Figure struct {
	Title           string        `json:"title"`
	Filename        string        `json:"filename"`
	Version         string        `json:"version"`
	URI             string        `json:"uri"`
	Type            string        `json:"type,omitempty"`
	ChartType       string        `json:"chartType,omitempty"`
	Subtitle        string        `json:"subtitle,omitempty"`
	Source          string        `json:"source,omitempty"`
	Notes           string        `json:"notes,omitempty"`
	AltText         string        `json:"altText,omitempty"`
	LabelInterval   string        `json:"labelInterval,omitempty"`
	DecimalPlaces   string        `json:"decimalPlaces,omitempty"`
	Unit            string        `json:"unit,omitempty"`
	AspectRatio     string        `json:"aspectRatio,omitempty"`
	Files           []SidecarFile `json:"files,omitempty"`
	Series          []string      `json:"series,omitempty"`
	Content         string
	Attribution     string
	Align           string
	DownloadFormats []string
	// TODO: files, categories, series, headers, data
}

type SidecarFile struct {
	Type     string `json:"type"`
	Filename string `json:"filename"`
	FileType string `json:"fileType"`
	Content  string
	FileSize string
}
