package helper

import (
	"fmt"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/ONSdigital/go-ns/common"
	"github.com/ONSdigital/log.go/log"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle
var localizers map[string]*i18n.Localizer

// InitialiseLocalisationsHelper sets up the core and service specific localisations for use in the templates
// with the Localise helper function
func InitialiseLocalisationsHelper(assetFn func(name string) ([]byte, error)) {
	bundle, _ = initLocaleBundle(assetFn)
	localizers = initLocalizer(bundle)
}

// InitLocalizer is used to initialise the localizer
func initLocalizer(bundle *i18n.Bundle) map[string]*i18n.Localizer {
	m := make(map[string]*i18n.Localizer)
	for _, locale := range common.SupportedLanguages {
		m[locale] = i18n.NewLocalizer(bundle, locale)

	}
	return m
}

// InitLocaleBundle is used to initialise the locale bundle
func initLocaleBundle(assetFn func(name string) ([]byte, error)) (*i18n.Bundle, error) {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	commonLocaliseNames := []string{"core", "service"}

	for _, locale := range common.SupportedLanguages {
		for _, fileName := range commonLocaliseNames {
			filePath := fmt.Sprintf("locales/%s.%s.toml", fileName, locale)
			asset, err := assetFn(filePath)
			if err != nil {
				log.Event(nil, "failed to get locale file", log.Error(err), log.ERROR)
			}
			bundle.ParseMessageFileBytes(asset, filePath)
		}
	}

	return bundle, nil
}

func Localise(key string, language string, plural int, templateArguments ...string) string {
	if key == "" {
		err := fmt.Errorf("key " + key + " not found in locale file")
		log.Event(nil, "no locale look up key provided", log.Error(err), log.ERROR)
		return ""
	}
	if language == "" {
		language = "en"
	}

	// Configure template data for arguments in strings
	templateData := make(map[string]string)
	for i, argument := range templateArguments {
		stringIndex := strconv.Itoa(i)
		key := "arg" + stringIndex
		templateData[key] = argument
	}

	loc := localizers[language]
	translation := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    key,
		PluralCount:  plural,
		TemplateData: templateData,
	})
	return translation
}
