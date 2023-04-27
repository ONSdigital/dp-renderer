package model_test

import (
	"strings"
	"testing"

	"github.com/ONSdigital/dp-renderer/v2/helper"
	"github.com/ONSdigital/dp-renderer/v2/model"
	. "github.com/smartystreets/goconvey/convey"
)

func mockAssetFunction(name string) ([]byte, error) {
	if strings.Contains(name, ".cy.toml") {
		return []byte("[Home]\none = \"Hafan\"\n"), nil
	}
	return []byte("[Home]\none = \"Home\"\n"), nil
}

func TestLocalise(t *testing.T) {
	helper.InitialiseLocalisationsHelper(mockAssetFunction)

	Convey("Should localise in the presence of a LocaleKey", t, func() {
		localisation := model.Localisation{
			LocaleKey: "Home",
			Plural:    1,
			Text:      "Home",
		}

		result := localisation.FuncLocalise("cy")

		So(result, ShouldEqual, "Hafan")
	})

	Convey("Should default to Text in the absence of a LocaleKey", t, func() {
		localisation := model.Localisation{
			Text: "Home",
		}

		result := localisation.FuncLocalise("cy")

		So(result, ShouldEqual, "Home")
	})
}
