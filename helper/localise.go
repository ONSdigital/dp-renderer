package helper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/ONSdigital/dp-net/v3/request"
	"github.com/ONSdigital/log.go/v2/log"
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

// InitLocaleBundle is used to initialise the locale bundle
func initLocaleBundle(assetFn func(name string) ([]byte, error)) (*i18n.Bundle, error) {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	commonLocaliseNames := []string{"core", "service"}

	for _, locale := range request.SupportedLanguages {
		for _, fileName := range commonLocaliseNames {
			filePath := fmt.Sprintf("locales/%s.%s.toml", fileName, locale)
			asset, err := assetFn(filePath)
			if err != nil {
				log.Error(context.Background(), "failed to get locale file", err)
			}
			_, err = bundle.ParseMessageFileBytes(asset, filePath)
			if err != nil {
				log.Error(context.Background(), "failed to parse message file", err)
			}
		}
	}

	return bundle, nil
}

// InitLocalizer is used to initialise the localizer
func initLocalizer(bundle *i18n.Bundle) map[string]*i18n.Localizer {
	m := make(map[string]*i18n.Localizer)
	for _, locale := range request.SupportedLanguages {
		m[locale] = i18n.NewLocalizer(bundle, locale)

	}
	return m
}

func Localise(key string, language string, plural int, templateArguments ...string) string {
	if key == "" {
		err := fmt.Errorf("key %s not found in locale file", key)
		log.Error(context.Background(), "no locale look up key provided", err)
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
