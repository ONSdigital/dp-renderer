package model_test

import (
	"strings"
	"testing"

	"github.com/ONSdigital/dp-renderer/v2/helper"
	"github.com/ONSdigital/dp-renderer/v2/model"
	. "github.com/smartystreets/goconvey/convey"
)

func mockPaginationLocales(name string) ([]byte, error) {
	locale := []string{
		"[Pagination]",
		"one = \"Pagination\"",
		"[PaginationPage]",
		"one = \"Page\"",
		"[PaginationOf]",
		"one = \"of\"",
		"[PaginationGoPrevious]",
		"one = \"Go to the previous page\"",
		"[PaginationGoFirst]",
		"one = \"Go to the first page\"",
		"[PaginationCurrentPage]",
		"one = \"Current page\"",
		"[PaginationGoLast]",
		"one = \"Go to the last page\"",
		"[PaginationGoNext]",
		"one = \"Go to the next page\"",
	}
	return []byte(strings.Join(locale, "\n")), nil
}

func TestPagination(t *testing.T) {
	helper.InitialiseLocalisationsHelper(mockPaginationLocales)

	Convey("FuncPhrasePageNOfTotal", t, func() {
		Convey("Should render a string like 'Page 1 of 10'", func() {
			pagination := model.Pagination{
				TotalPages: 10,
			}
			So(pagination.FuncPhrasePageNOfTotal(1, "en"), ShouldEqual, "Page 1 of 10")
		})
	})

	Convey("FuncPhrasePaginationProgress", t, func() {
		Convey("Should render a progress string in parentheses", func() {
			pagination := model.Pagination{}
			So(pagination.FuncPhrasePaginationProgress("progress", "en"), ShouldEqual, "Pagination (progress)")
		})
	})

	Convey("FuncPhraseGoToPreviousPage", t, func() {
		Convey("Should render direction with page number in parentheses", func() {
			pagination := model.Pagination{
				CurrentPage: 2,
			}
			So(pagination.FuncPhraseGoToPreviousPage("en"), ShouldEqual, "Go to the previous page (Page 1)")
		})
	})

	Convey("FuncPhraseGoToNextPage", t, func() {
		Convey("Should render direction with page number in parentheses", func() {
			pagination := model.Pagination{
				CurrentPage: 2,
			}
			So(pagination.FuncPhraseGoToNextPage("en"), ShouldEqual, "Go to the next page (Page 3)")
		})
	})

	Convey("FuncPhraseGoToFirstPage", t, func() {
		Convey("Should render direction with page number in parentheses", func() {
			pagination := model.Pagination{}
			So(pagination.FuncPhraseGoToFirstPage("en"), ShouldEqual, "Go to the first page (Page 1)")
		})
	})

	Convey("FuncPhraseCurrentPage", t, func() {
		Convey("Should render a progress string in parentheses", func() {
			pagination := model.Pagination{}
			So(pagination.FuncPhraseCurrentPage("progress", "en"), ShouldEqual, "Current page (progress)")
		})
	})

	Convey("FuncPhraseGoToLastPage", t, func() {
		Convey("Should render direction with page number in parentheses", func() {
			pagination := model.Pagination{
				TotalPages: 10,
			}
			So(pagination.FuncPhraseGoToLastPage("en"), ShouldEqual, "Go to the last page (Page 10)")
		})
	})

	Convey("FuncPickPreviousURL", t, func() {
		Convey("Should pick the URL of the previous page", func() {
			pagination := model.Pagination{
				PagesToDisplay: []model.PageToDisplay{
					{PageNumber: 1, URL: "previous"},
					{PageNumber: 2, URL: "current"},
				},
				CurrentPage: 2,
			}
			So(pagination.FuncPickPreviousURL(), ShouldEqual, "previous")
		})

		Convey("Should return an empty fragment when the previous page lies outside the display window", func() {
			pagination := model.Pagination{
				PagesToDisplay: []model.PageToDisplay{
					{PageNumber: 1, URL: "current"},
					{PageNumber: 2, URL: "2"},
				},
				CurrentPage: 1,
			}
			So(pagination.FuncPickPreviousURL(), ShouldEqual, "#")
		})
	})

	Convey("FuncPickNextURL", t, func() {
		Convey("Should pick the URL of the next page", func() {
			pagination := model.Pagination{
				PagesToDisplay: []model.PageToDisplay{
					{PageNumber: 1, URL: "current"},
					{PageNumber: 2, URL: "next"},
				},
				CurrentPage: 1,
			}
			So(pagination.FuncPickNextURL(), ShouldEqual, "next")
		})

		Convey("Should return an empty fragment when the next page lies outside the display window", func() {
			pagination := model.Pagination{
				PagesToDisplay: []model.PageToDisplay{
					{PageNumber: 1, URL: "1"},
					{PageNumber: 2, URL: "current"},
				},
				CurrentPage: 2,
			}
			So(pagination.FuncPickNextURL(), ShouldEqual, "#")
		})
	})

	Convey("FuncShowLinkToFirst", t, func() {
		Convey("Should be false when the display window begins with the first page", func() {
			pagination := model.Pagination{
				PagesToDisplay: []model.PageToDisplay{
					{PageNumber: 1, URL: ""},
				},
			}
			So(pagination.FuncShowLinkToFirst(), ShouldBeFalse)
		})

		Convey("Should be true otherwise", func() {
			pagination := model.Pagination{
				PagesToDisplay: []model.PageToDisplay{
					{PageNumber: 2, URL: ""},
				},
			}
			So(pagination.FuncShowLinkToFirst(), ShouldBeTrue)
		})
	})

	Convey("FuncShowLinkToLast", t, func() {
		Convey("Should be false when the display window ends with the last page", func() {
			pagination := model.Pagination{
				PagesToDisplay: []model.PageToDisplay{
					{PageNumber: 10, URL: ""},
				},
				TotalPages: 10,
			}
			So(pagination.FuncShowLinkToLast(), ShouldBeFalse)
		})

		Convey("Should be true otherwise", func() {
			pagination := model.Pagination{
				PagesToDisplay: []model.PageToDisplay{
					{PageNumber: 9, URL: ""},
				},
				TotalPages: 10,
			}
			So(pagination.FuncShowLinkToLast(), ShouldBeTrue)
		})
	})
}
