package model

import (
	"fmt"

	"github.com/ONSdigital/dp-renderer/v2/helper"
)

// Pagination represents all information regarding pagination of search results
type Pagination struct {
	CurrentPage       int             `json:"current_page"`
	PagesToDisplay    []PageToDisplay `json:"pages_to_display"`
	FirstAndLastPages []PageToDisplay `json:"first_and_last_pages"`
	TotalPages        int             `json:"total_pages"`
	Limit             int             `json:"limit"`
	LimitOptions      []int           `json:"limit_options,omitempty"`
}

// PageToDisplay represents a page to display in pagination with their corresponding URL
type PageToDisplay struct {
	PageNumber int    `json:"page_number"`
	URL        string `json:"url"`
}

// Produces a string of the form "Page 1 of 10"
func (pagination Pagination) FuncPhrasePageNOfTotal(n int, language string) string {
	phrasePage := helper.Localise("PaginationPage", language, 1)
	phraseOf := helper.Localise("PaginationOf", language, 1)
	return fmt.Sprintf("%s %d %s %d", phrasePage, n, phraseOf, pagination.TotalPages)
}

// Produces a string of the form "Pagination (Page 1 of 10)"
func (pagination Pagination) FuncPhrasePaginationProgress(progress, language string) string {
	phrasePagination := helper.Localise("Pagination", language, 1)
	return fmt.Sprintf("%s (%s)", phrasePagination, progress)
}

// Produces a string of the form "Go to the previous page (Page 4)"
func (pagination Pagination) FuncPhraseGoToPreviousPage(language string) string {
	phrasePage := helper.Localise("PaginationPage", language, 1)
	phraseGoPrevious := helper.Localise("PaginationGoPrevious", language, 1)
	pageNumber := pagination.CurrentPage - 1
	return fmt.Sprintf("%s (%s %d)", phraseGoPrevious, phrasePage, pageNumber)
}

// Produces a string of the form "Go to the next page (Page 6)"
func (pagination Pagination) FuncPhraseGoToNextPage(language string) string {
	phrasePage := helper.Localise("PaginationPage", language, 1)
	phraseGoNext := helper.Localise("PaginationGoNext", language, 1)
	pageNumber := pagination.CurrentPage + 1
	return fmt.Sprintf("%s (%s %d)", phraseGoNext, phrasePage, pageNumber)
}

// Produces a string of the form "Go to the first page (Page 1)"
func (pagination Pagination) FuncPhraseGoToFirstPage(language string) string {
	phrasePage := helper.Localise("PaginationPage", language, 1)
	phraseGoFirst := helper.Localise("PaginationGoFirst", language, 1)
	return fmt.Sprintf("%s (%s 1)", phraseGoFirst, phrasePage)
}

// Produces a string of the form "Current page (Page 5 of 10)"
func (pagination Pagination) FuncPhraseCurrentPage(progress, language string) string {
	phrasePagination := helper.Localise("PaginationCurrentPage", language, 1)
	return fmt.Sprintf("%s (%s)", phrasePagination, progress)
}

// Produces a string of the form "Go to the last page (Page 10)"
func (pagination Pagination) FuncPhraseGoToLastPage(language string) string {
	phrasePage := helper.Localise("PaginationPage", language, 1)
	phraseGoLast := helper.Localise("PaginationGoLast", language, 1)
	return fmt.Sprintf("%s (%s %d)", phraseGoLast, phrasePage, pagination.TotalPages)
}

func (pagination Pagination) FuncPickPreviousURL() string {
	for _, pageToDisplay := range pagination.PagesToDisplay {
		if pageToDisplay.PageNumber == pagination.CurrentPage-1 {
			return pageToDisplay.URL
		}
	}

	return "#"
}

func (pagination Pagination) FuncPickNextURL() string {
	for _, pageToDisplay := range pagination.PagesToDisplay {
		if pageToDisplay.PageNumber == pagination.CurrentPage+1 {
			return pageToDisplay.URL
		}
	}

	return "#"
}

func (pagination Pagination) FuncShowLinkToFirst() bool {
	return pagination.PagesToDisplay[0].PageNumber > 1
}

func (pagination Pagination) FuncShowLinkToLast() bool {
	lastPage := len(pagination.PagesToDisplay) - 1
	return pagination.PagesToDisplay[lastPage].PageNumber != pagination.TotalPages
}
