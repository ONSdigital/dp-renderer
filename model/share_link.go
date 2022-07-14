package model

import (
	"fmt"
	"net/url"
)

// Enumerates known social types that allow sharing a resource via a link URL
type SocialType int

const (
	SocialUnknown SocialType = iota
	SocialEmail
	SocialFacebook
	SocialLinkedin
	SocialTwitter
)

// Represents an instance of a link URL for a specific shared resource
type ShareLink struct {
	Type               SocialType
	Url                string
	RequiresJavaScript bool
}

// Stringer interface
func (s SocialType) String() string {
	var result string
	switch s {
	case SocialUnknown:
		result = "unknown"
	case SocialEmail:
		result = "email"
	case SocialFacebook:
		result = "facebook"
	case SocialLinkedin:
		result = "linkedin"
	case SocialTwitter:
		result = "twitter"
	}
	return result
}

func emailLink(title, target string) ShareLink {
	url := fmt.Sprintf("mailto:?subject=%s&body=%s%%0A%s", title, title, target)
	return ShareLink{
		Type:               SocialEmail,
		Url:                url,
		RequiresJavaScript: false,
	}
}

func facebookLink(target string) ShareLink {
	url := fmt.Sprintf("https://www.facebook.com/sharer.php?u=%s", target)
	return ShareLink{
		Type:               SocialFacebook,
		Url:                url,
		RequiresJavaScript: true,
	}
}

func linkedinLink(target string) ShareLink {
	url := fmt.Sprintf("https://www.linkedin.com/sharing/share-offsite/?url=%s", target)
	return ShareLink{
		Type:               SocialLinkedin,
		Url:                url,
		RequiresJavaScript: true,
	}
}

func twitterLink(title, target string) ShareLink {
	url := fmt.Sprintf("https://twitter.com/intent/tweet?text=%s&url=%s", title, target)
	return ShareLink{
		Type:               SocialTwitter,
		Url:                url,
		RequiresJavaScript: true,
	}
}

// Creates a ShareLink from the supplied resource title and target URL
func (s SocialType) CreateLink(title, target string) ShareLink {
	var result ShareLink
	escTitle := url.QueryEscape(title)
	escTarget := url.QueryEscape(target)

	switch s {
	case SocialEmail:
		result = emailLink(escTitle, escTarget)
	case SocialFacebook:
		result = facebookLink(escTarget)
	case SocialLinkedin:
		result = linkedinLink(escTarget)
	case SocialTwitter:
		result = twitterLink(escTitle, escTarget)
	}

	return result
}
