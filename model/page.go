package model

import (
	"html/template"
)

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
	SiteDomain                       string           `json:"-"`
	PatternLibraryAssetsPath         string           `json:"pattern_library_assets_path"`
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
	RemoveGalleryBackground          bool             `json:"remove_gallery_background"`
	Feedback                         Feedback         `json:"feedback"`
	Enable500ErrorPageStyling        bool             `json:"enable_500_error_page_styling"` // flag for hiding standard page "furniture" (header, nav, etc.)
	ABTest
}

// ABTest contains all information needed for ABTesting - this is separated for expansion in future.
type ABTest struct {
	GTMKey string `json:"abtest_gtm_key"` // key for GTM to differentiate test pages.
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
	EnableFeedbackAPI      bool   `json:"enable_feedback_api"` // Deprecated: EnableFeedbackAPI should not be used, it will be removed soon
	FeedbackAPIURL         string `json:"feedback_api_url"`    // technically not a feature flag, but used exclusivly with one
	IsPublishing           bool   `json:"is_publishing"`
}

// NewPage instantiates the base Page type with configurable fields
func NewPage(path, domain string) Page {
	return Page{
		PatternLibraryAssetsPath: path,
		SiteDomain:               domain,
	}
}
