package model

import "html/template"

// Page contains data re-used for each page type a Data struct for data specific to the page type
type Page struct {
	Count                            int              `json:"count"`
	Type                             string           `json:"type"`
	DatasetId                        string           `json:"dataset_id"`
	DatasetTitle                     string           `json:"dataset_title"`
	URI                              string           `json:"uri"`
	Taxonomy                         []TaxonomyNode   `json:"taxonomy"`
	Breadcrumb                       []TaxonomyNode   `json:"breadcrumb"`
	IsInFilterBreadcrumb             bool             `json:"is_in_filter_breadcrumb"`
	ServiceMessage                   string           `json:"service_message"`
	Metadata                         Metadata         `json:"metadata"`
	SearchDisabled                   bool             `json:"search_disabled"`
	SearchBarFormAction              string           `json:"search_bar_form_action"`
	SearchBarLocaliseKey             string           `json:"search_bar_localise_key"`
	SiteDomain                       string           `json:"-"`
	PatternLibraryAssetsPath         string           `json:"-"`
	Language                         string           `json:"language"`
	IncludeAssetsIntegrityAttributes bool             `json:"-"`
	ReleaseDate                      string           `json:"release_date"`
	BetaBannerEnabled                bool             `json:"beta_banner_enabled"`
	CookiesPreferencesSet            bool             `json:"cookies_preferences_set"`
	CookiesPolicy                    CookiesPolicy    `json:"cookies_policy"`
	HasJSONLD                        bool             `json:"has_jsonld"`
	FeatureFlags                     FeatureFlags     `json:"feature_flags"`
	Error                            Error            `json:"error"`
	EmergencyBanner                  EmergencyBanner  `json:"emergency_banner"`
	Collapsible                      Collapsible      `json:"collapsible"`
	Pagination                       Pagination       `json:"pagination"`
	TableOfContents                  TableOfContents  `json:"table_of_contents"`
	BackTo                           BackTo           `json:"back_to"`
	SearchNoIndexEnabled             bool             `json:"search_no_index_enabled"`
	NavigationContent                []NavigationItem `json:"navigation_content"`
	PreGTMJavaScript                 []template.JS    `json:"pre_gtm_javascript"`
}

// NavigationItem contains all information needed to render the navigation bar and submenus
type NavigationItem struct {
	Uri      string           `json:"uri"`
	Label    string           `json:"label"`
	SubItems []NavigationItem `json:"sub_items"`
}

// FeatureFlags contains toggles for certain features on the website
type FeatureFlags struct {
	EnableCensusBanner     bool   `json:"enable_census_banner"`
	EnableCensusTile       bool   `json:"enable_census_tile"`
	EnableGetDataCard      bool   `json:"enable_get_data_card"`
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
