package model

// CookiesPolicy contains data for the users cookie policy
type CookiesPolicy struct {
	Communications bool `json:"communications"`
	Essential      bool `json:"essential"`
	Settings       bool `json:"settings"`
	Usage          bool `json:"usage"`
}
