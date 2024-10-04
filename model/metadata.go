package model

// Metadata ...
type Metadata struct {
	Title        string   `json:"title"`
	DisplayTitle string   `json:"display_title"`
	Description  string   `json:"description"`
	ServiceName  string   `json:"serviceName"`
	Keywords     []string `json:"keywords"`
}
