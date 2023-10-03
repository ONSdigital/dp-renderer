package model_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/v2/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFuncFeedback(t *testing.T) {
	Convey("Given a page request", t, func() {
		Convey("When FuncFeedback is called", func() {
			p := model.Page{}
			Convey("Then the expected mapped model is returned", func() {
				expected := model.Feedback{
					Description: model.TextareaField{
						Input: model.Input{
							Autocomplete: "off",
							DataAttributes: []model.DataAttribute{
								{
									Key: "value-missing",
									Value: model.Localisation{
										LocaleKey: "ImproveThisPageInvalid",
										Plural:    1,
									},
								},
							},
							ID:         "description-field",
							IsRequired: true,
							Label: model.Localisation{
								LocaleKey: "ImproveThisPage",
								Plural:    1,
							},
							Language: p.Language,
							Name:     "description",
						},
					},
					NameInput: model.TextField{
						Input: model.Input{
							Autocomplete: "name",
							ID:           "name-field",
							Label: model.Localisation{
								LocaleKey: "NameOpt",
								Plural:    1,
							},
							Name:     "name",
							Type:     model.Text,
							Language: p.Language,
						},
					},
					EmailInput: model.TextField{
						Input: model.Input{
							Autocomplete: "email",
							DataAttributes: []model.DataAttribute{
								{
									Key: "type-mismatch",
									Value: model.Localisation{
										LocaleKey: "EmailInvalid",
										Plural:    1,
									},
								},
							},
							ID: "email-field",
							Label: model.Localisation{
								LocaleKey: "EmailOpt",
								Plural:    1,
							},
							Name:     "email",
							Type:     model.Email,
							Language: p.Language,
						},
					},
				}
				So(p.FuncFeedback(), ShouldEqual, expected)
			})
		})

		Convey("When the language is set to 'en' on the page model", func() {
			p := model.Page{
				Language: "en",
			}
			Convey("Then the set language is mapped to the feedback model", func() {
				So(p.FuncFeedback().Description.Input.Language, ShouldEqual, "en")
			})
		})
		Convey("When the language is set to 'cy' on the page model", func() {
			p := model.Page{
				Language: "cy",
			}
			Convey("Then the set language is mapped to the feedback model", func() {
				So(p.FuncFeedback().Description.Input.Language, ShouldEqual, "cy")
			})
		})
	})
}
