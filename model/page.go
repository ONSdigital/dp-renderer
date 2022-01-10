package model

// Page contains data re-used for each page type a Data struct for data specific to the page type
type Page struct {
	Count                            int             `json:"count"`
	Type                             string          `json:"type"`
	DatasetId                        string          `json:"dataset_id"`
	DatasetTitle                     string          `json:"dataset_title"`
	URI                              string          `json:"uri"`
	Taxonomy                         []TaxonomyNode  `json:"taxonomy"`
	Breadcrumb                       []TaxonomyNode  `json:"breadcrumb"`
	IsInFilterBreadcrumb             bool            `json:"is_in_filter_breadcrumb"`
	ServiceMessage                   string          `json:"service_message"`
	Metadata                         Metadata        `json:"metadata"`
	SearchDisabled                   bool            `json:"search_disabled"`
	SiteDomain                       string          `json:"-"`
	PatternLibraryAssetsPath         string          `json:"-"`
	Language                         string          `json:"language"`
	IncludeAssetsIntegrityAttributes bool            `json:"-"`
	ReleaseDate                      string          `json:"release_date"`
	BetaBannerEnabled                bool            `json:"beta_banner_enabled"`
	CookiesPreferencesSet            bool            `json:"cookies_preferences_set"`
	CookiesPolicy                    CookiesPolicy   `json:"cookies_policy"`
	HasJSONLD                        bool            `json:"has_jsonld"`
	FeatureFlags                     FeatureFlags    `json:"feature_flags"`
	Error                            Error           `json:"error"`
	EmergencyBanner                  EmergencyBanner `json:"emergency_banner"`
}

// FeatureFlags contains toggles for certain features on the website
type FeatureFlags struct {
	HideCookieBanner       bool   `json:"hide_cookie_banner"`
	ONSDesignSystemVersion string `json:"ons_design_system_version"`
	SixteensVersion        string `json:"legacy_sixteens_version"`
}

// NewPage instantiates the base Page type with configurable fields
func NewPage(path, domain string) Page {
	return Page{
		PatternLibraryAssetsPath: path,
		SiteDomain:               domain,
	}
}
