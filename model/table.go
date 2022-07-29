package model

// Defines the form button visible in a data cell
type TableFormButton struct {
	Text    string `json:"text"`    // Label text, required
	Id      string `json:"id"`      // HTML id attribute
	Name    string `json:"name"`    // HTML name attribute
	Value   string `json:"value"`   // HTML value attribute
	Url     string `json:"url"`     // Link destination
	Classes string `json:"classes"` // Additional classes appended to the HTML class attribute
}

// Defines hidden form input fields
type TableFormHiddenField struct {
	Name  string `json:"name"` // HTML input name attribute
	Value string `json:"vale"` // HTML input value attribute
}

// TableForm can be placed in data cells to present a TableFormButton
type TableForm struct {
	Method           string                 `json:"method"`             // HTML method attribute, defaults to "POST"
	Action           string                 `json:"action"`             // HTML action attribute, required
	Button           TableFormButton        `json:"button"`             // Button shown to the user, required
	HiddenFormFields []TableFormHiddenField `json:"hidden_form_fields"` // Hidden form inputs
}

// Defines a column footer
type TableFooter struct {
	Value string `json:"value"` // Cell content in the footer, required
}

// TableData defines an individual data cell
type TableData struct {
	TdClasses string     `json:"td_classes"` // Additional classes appended to the HTML class attribute
	Id        string     `json:"id"`         // HTML id attribute
	Name      string     `json:"name"`       // HTML name attribute
	Data      string     `json:"data"`       // Required by the responsive variant, sets the data-th attribute for this cell to its matching column header
	DataSort  string     `json:"data_sort"`  // Required by the sortable variant, sets the numerical order of a table cell in a column
	Value     string     `json:"value"`      // Cell content
	Numeric   bool       `json:"numeric"`    // When true, right aligns content for easier comparison of numeric data
	Form      *TableForm `json:"form"`       // Settings for a form within the cell
}

// TableHeader defines column headers
type TableHeader struct {
	ThClasses string `json:"th_classes"` // Additional classes appended to the HTML class attribute
	AriaSort  string `json:"aria_sort"`  // Set to “ascending” or “descending” to set the default order of a table column when the page loads when setting variants to “sortable”. Defaults to “none”
	Value     string `json:"value"`      // Header content
	Numeric   bool   `json:"numeric"`    // When true, right aligns content for easier comparison of numeric data
}

// TableRow holds rows of data
type TableRow struct {
	TableData []TableData `json:"table_data"` // An array of data cells populating the row, required
	Id        string      `json:"id"`         // HTML id attribute
	Name      string      `json:"name"`       // HTML name attribute
	Highlight bool        `json:"highlight"`  // When true, highlights the row
}

// Table displays data in a variety of tabular styles that can be combined in Variants
type Table struct {
	Variants     []string      `json:"variants"`      // Possible variants: "compact", "responsive", "scrollable", "sortable", and "row-hover"
	TableClasses string        `json:"table_classes"` // Additional classes appended to the HTML class attribute
	Id           string        `json:"id"`            // HTML id attribute
	Caption      string        `json:"caption"`       // Content of the HTML caption
	HideCaption  bool          `json:"hide_caption"`  // When true, visually hides the caption
	AriaLabel    string        `json:"aria_label"`    // The ARIA label to be added if ”scrollable” variant set, to inform screen reader users that the table can be scrolled. Defaults to “Scrollable table“.
	TableHeaders []TableHeader `json:"table_headers"` // An array of column headers, required
	TableRows    []TableRow    `json:"trs"`           // An array or data rows, required
	SortBy       string        `json:"sort_by"`       // Required the "sortable" variant, sets the data-aria-sort attribute for the table. Used as a prefix for the aria-label to announce to screen readers when the table is sorted by a column. For example, “Sort by Date, ascending”
	AriaAsc      string        `json:"aria_asc"`      // Required by "sortable" variant, sets the data-aria-asc attribute for the table. Used to update aria-sort attribute to announce to screen readers how a table is sorted by a column, for example, “Sort by Date, ascending“
	AriaDesc     string        `json:"aria_desc"`     // Required by "sortable" variant, sets the data-aria-desc attribute for the table. Used to update aria-sort attribute to announce to screen readers how a table is sorted by a column, for example, “Sort by Date, descending“
	TableFooters []TableFooter `json:"table_footers"` // An array of cells for the footer of each column
}

// Tests for the presence of a known variant in the options supplied by the user
func (table Table) FuncContainsVariant(variant string) bool {
	for _, v := range table.Variants {
		if v == variant {
			return true
		}
	}
	return false
}
