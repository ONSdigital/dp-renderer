package helper

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/ONSdigital/dp-net/v2/request"
	"github.com/ONSdigital/log.go/v2/log"
)

func DomainSetLang(domain string, uri string, language string) string {
	languageSupported := false
	for _, locale := range request.SupportedLanguages {
		if locale == language {
			languageSupported = true
		}
	}

	// uri comes in inconsistently, remove domain and port if they come through in uri param
	var findEndpointRE = regexp.MustCompile(`https?://[^/]+(.*)`)
	if endpoint := findEndpointRE.FindStringSubmatch(uri); len(endpoint) == 2 {
		uri = endpoint[1]
	}

	url := domain + uri

	strippedURL := strings.Replace(url, "https://", "", 1)
	strippedURL = strings.Replace(strippedURL, "www.", "", 1)

	for _, locale := range request.SupportedLanguages {
		possibleLocaleURLPrefix := strippedURL[0:len(locale)]

		if possibleLocaleURLPrefix == locale {
			trimLength := len(locale) + 1
			strippedURL = strippedURL[trimLength:]
			break
		}
	}

	domainWithTranslation := ""
	if !languageSupported {
		err := fmt.Errorf("Language: " + language + " is not supported resolving to " + request.DefaultLang)
		log.Error(context.Background(), "language fail", err)
	}
	if language == request.DefaultLang || !languageSupported {
		domainWithTranslation = "https://www." + strippedURL
	} else {
		domainWithTranslation = "https://" + language + "." + strippedURL
	}

	return domainWithTranslation
}
