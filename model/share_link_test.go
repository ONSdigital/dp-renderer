package model_test

import (
	"net/url"
	"testing"

	"github.com/ONSdigital/dp-renderer/v2/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestShareLink(t *testing.T) {
	Convey("SocialType", t, func() {
		Convey("Should implement the Stringer interface", func() {
			So(model.SocialEmail.String(), ShouldEqual, "email")
		})
	})

	Convey("CreateLink", t, func() {
		Convey("SocialUnknown", func() {
			title := "Test Title"
			target := "https://www.example.com"
			shareLink := model.SocialUnknown.CreateLink(title, target)

			Convey("Should create a link of the specified type", func() {
				So(shareLink.Type, ShouldEqual, model.SocialUnknown)
			})

			Convey("Should indicate the medium's dependency on JavaScript", func() {
				So(shareLink.RequiresJavaScript, ShouldBeFalse)
			})

			Convey("Should have an empty link", func() {
				So(shareLink.Url, ShouldEqual, "")
			})
		})

		Convey("SocialEmail", func() {
			title := "Test Title"
			target := "https://www.example.com"
			shareLink := model.SocialEmail.CreateLink(title, target)

			Convey("Should create a link of the specified type", func() {
				So(shareLink.Type, ShouldEqual, model.SocialEmail)
			})

			Convey("Should indicate the medium's dependency on JavaScript", func() {
				So(shareLink.RequiresJavaScript, ShouldBeFalse)
			})

			Convey("Should contain a well formed link", func() {
				parsedUrl, err := url.Parse(shareLink.Url)
				So(err, ShouldBeNil)
				params := parsedUrl.Query()

				Convey("Should create a link with the correct protocol", func() {
					So(parsedUrl.Scheme, ShouldEqual, "mailto")
				})

				Convey("Should create a link containing the title", func() {
					So(params.Get("subject"), ShouldEqual, title)
				})

				Convey("Should create a link containing the target", func() {
					So(params.Get("body"), ShouldContainSubstring, target)
				})
			})
		})

		Convey("SocialFacebook", func() {
			title := "Test Title"
			target := "https://www.example.com"
			shareLink := model.SocialFacebook.CreateLink(title, target)

			Convey("Should create a link of the specified type", func() {
				So(shareLink.Type, ShouldEqual, model.SocialFacebook)
			})

			Convey("Should indicate the medium's dependency on JavaScript", func() {
				So(shareLink.RequiresJavaScript, ShouldBeTrue)
			})

			Convey("Should contain a well formed link", func() {
				parsedUrl, err := url.Parse(shareLink.Url)
				So(err, ShouldBeNil)
				params := parsedUrl.Query()

				Convey("Should create a link with the correct protocol", func() {
					So(parsedUrl.Scheme, ShouldEqual, "https")
				})

				Convey("Should create a link with the correct domain", func() {
					So(parsedUrl.Hostname(), ShouldEqual, "www.facebook.com")
				})

				Convey("Should create a link containing the target", func() {
					So(params.Get("u"), ShouldEqual, target)
				})
			})
		})

		Convey("SocialLinkedin", func() {
			title := "Test Title"
			target := "https://www.example.com"
			shareLink := model.SocialLinkedin.CreateLink(title, target)

			Convey("Should create a link of the specified type", func() {
				So(shareLink.Type, ShouldEqual, model.SocialLinkedin)
			})

			Convey("Should indicate the medium's dependency on JavaScript", func() {
				So(shareLink.RequiresJavaScript, ShouldBeTrue)
			})

			Convey("Should contain a well formed link", func() {
				parsedUrl, err := url.Parse(shareLink.Url)
				So(err, ShouldBeNil)
				params := parsedUrl.Query()

				Convey("Should create a link with the correct protocol", func() {
					So(parsedUrl.Scheme, ShouldEqual, "https")
				})

				Convey("Should create a link with the correct domain", func() {
					So(parsedUrl.Hostname(), ShouldEqual, "www.linkedin.com")
				})

				Convey("Should create a link containing the target", func() {
					So(params.Get("url"), ShouldEqual, target)
				})
			})
		})

		Convey("SocialTwitter", func() {
			title := "Test Title"
			target := "https://www.example.com"
			shareLink := model.SocialTwitter.CreateLink(title, target)

			Convey("Should create a link of the specified type", func() {
				So(shareLink.Type, ShouldEqual, model.SocialTwitter)
			})

			Convey("Should indicate the medium's dependency on JavaScript", func() {
				So(shareLink.RequiresJavaScript, ShouldBeTrue)
			})

			Convey("Should contain a well formed link", func() {
				parsedUrl, err := url.Parse(shareLink.Url)
				So(err, ShouldBeNil)
				params := parsedUrl.Query()

				Convey("Should create a link with the correct protocol", func() {
					So(parsedUrl.Scheme, ShouldEqual, "https")
				})

				Convey("Should create a link with the correct domain", func() {
					So(parsedUrl.Hostname(), ShouldEqual, "twitter.com")
				})

				Convey("Should create a link containing the title", func() {
					So(params.Get("text"), ShouldEqual, title)
				})

				Convey("Should create a link containing the target", func() {
					So(params.Get("url"), ShouldEqual, target)
				})
			})
		})
	})
}
