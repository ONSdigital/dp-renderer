# Using patterns or components in your service

Patterns and Components from the ONS Design System are made available in the
`dp-renderer` as Go Templates and their associated data models.

Templates are also often referred to as 'partials' as they live in
`assets/templates/partials`.

Data models, or simply 'models', live in `model` and by convention carry the
same name as the template they drive.

Models may implement methods. These are used within the component template,
rather than by the application importing the template. When used within a
component template they keep the template concise and encourage re-use of logic.
The Go Template syntax makes it difficult to distinguish a property reference
from a method call, so within the `dp-renderer` the convention of prefixing
methods intended for use within templates with `Func` has been adopted for
clarity.

For example, this is a (ficticious) method call:

```
{{ .Component.FuncIsSubmitVisible }}
```

and this is a property reference:

```
{{ .Component.IsSubmitVisible }}
```

Without the naming convention these would be indistinguishable, unless a
function takes parameters which makes it a little more obvious:

```
{{ .Component.FuncIsSubmitVisible .State .Language }}
```

## Contents

- [Collapsible](#collapsible)
- [Table of Contents](#table-of-contents)
- [Pagination](#pagination)
- [Input date](#inputdate)
- [Back to](#backto)
- [Table](#table)
- [Correct errors](#correct-errors)
- [Fields](#fields)

## Collapsible

To instantiate the [collapsible](https://service-manual.ons.gov.uk/design-system/components/accordion) (renamed to accordion) UI component in your service:

- In the `mapper.go` file in your service, populate the relevant fields
  e.g.

```go
p.Collapsible = coreModel.Collapsible{
  // You can either use a localisation key or populate the `Text` string
  Title: coreModel.Localisation{
    LocaleKey: "VariablesExplanation",
    Plural:    4,
  },
  // You can use `content` or `safeHTML` to inject content into the section
  CollapsibleItems: []coreModel.CollapsibleItem{
   {
    Subheading: "This is a subheading",
    Content:    []string{"a string"},
   },
   {
    Subheading: "This is another subheading",
    Content:    []string{"another string", "and another"},
   },
   {
    Subheading: "A third subheading",
    SafeHTML: coreModel.Localisation{
      LocaleKey: "LocaleKey",
      Plural:    1,
    },
   },
  },
 }
```

- In the template file within your service, reference the `collapsible.tmpl` file
  e.g.

```tmpl
<div>Some html...</div>
{{ template "partials/collapsible" .Collapsible }}
<div>Some more html</div>
```

## Table of Contents

To instantiate the [table-of-contents](https://service-manual.ons.gov.uk/design-system/components/table-of-contents) UI component in your service:

- In the `mapper.go` file in your service, populate the relevant fields:

```go
page.TableOfContents = TableOfContents{
  Id: "toc",
  AriaLabel: "Table of contents",
  Title: "Contents",
  Sections: map[string]ContentSection{
    "trees": {
      Current: false,
      Title: "All about trees",
    },
    "flowers": {
      Current: true,
      Title: "All about flowers",
    },
  },
  DisplayOrder: []string{
    "flowers",
    "trees",
  }
}
```

The `Id` is optional.

The keys of the Sections map must match the entries in DisplayOrder,
and these keys are used as the fragment IDs in the anchor tags.

Omitting a section's key from the DisplayOrder will prevent that
section from being listed in the table of contents.

- In the template file within your service, reference the
  `table-of-contents.tmpl` file, where `.` is the Page model:

```tmpl
<div>Some html...</div>
{{ template "partials/table-of-contents" . }}
<div>Some more html</div>
```

### Localisation

Entries in `assets/locales/core.<lang>.toml` can be referenced by their
keys to enable localisation:

- `AriaLabel` may be overridden by `AriaLabelLocaliseKey`
- `Title` may be overridden by `TitleLocaliseKey`

## Pagination

To instantiate the [pagination](https://service-manual.ons.gov.uk/design-system/components/pagination) UI component in your service:

- In the `mapper.go` file in your service, populate the relevant fields:

```go
page.Pagination = Pagination{
  CurrentPage: 5,
  TotalPages: 10,
  Limit: 10,
  LimitOptions: []int{
    10,
    20,
  },
  FirstAndLastPages: []PageToDisplay{
    {
      PageNumber: 1,
      URL: "results/1",
    },
    {
      PageNumber: 10,
      URL: "results/10",
    },
  },
  PagesToDisplay: []PageToDisplay{
    {
      PageNumber: 4,
      URL: "results/4",
    },
    {
      PageNumber: 5,
      URL: "results/5",
    },
    {
      PageNumber: 6,
      URL: "results/6",
    },
  },
}
```

`Limit` represents the number of results shown at once a.k.a.
results per page.

`LimitOptions` is currently unused, and is therefore optional. It
represents the available result limits. If used, `Limit` should
be set to one of the available options in `LimitOptions`.

Offering the user the choice of how many results per page should be
shown could be done by extending the `Pagination` component or
creating a new component intended to be used alongside the
`Pagination` component that consumes this same `Pagination` model.

`FirstAndLastPages` and `PagesToDisplay` interact with each other
to govern the way available pages are displayed.

`FirstAndLastPages` are used to bookend the row of page numbers.
They are always visible.

`PagesToDisplay` represents a moving window into all possible
pages. When this window overlaps a bookend, there will be no
ellipsis (...) shown between the window and the bookend.

The size of this window is up to you. In the examples that follow
it is set to three.

For example, with 1 and 10 as bookends defined in `FirstAndLastPages`,
setting `PagesToDisplay = { 1, 2, 3 }` closes the gap between
1 (a bookend) and 2:

```
1 2 3 ... 10
```

Similarly, setting `PagesToDisplay = { 8, 9, 10 }` closes the
gap between 9 and 10 (a bookend):

```
1 ... 8 9 10
```

When the window defined by `PagesToDisplay` does not overlap with
the bookends defined in `FirstAndLastPages`, the page options are
rendered with ellipsis between both bookends. For example, setting
`PagesToDisplay = { 4, 5, 6 }` renders as:

```
1 ... 4 5 6 ... 10
```

- In the template file within your service, reference the
  `pagination.tmpl` file, where `.` is the Page model:

```tmpl
<div>Some html...</div>
{{ template "partials/pagination" . }}
<div>Some more html</div>
```

### Localisation

All translations live in `assets/locales/core.<language>.toml` and
are prefixed with `Pagination`.

## InputDate

To instantiate the [Dates](https://service-manual.ons.gov.uk/design-system/patterns/dates) UI pattern in your service:

- In the `mapper.go` file in your service, populate the relevant fields:

```go
page.PublicationDate = InputDate{
  Language:        page.Language,
  Id:              "publication-date",
  InputNameDay:    "publication-day",
  InputNameMonth:  "publication-month",
  InputNameYear:   "publication-year",
  InputValueDay:   "5",
  InputValueMonth: "6",
  InputValueYear:  "1950",
  Title:           "Publication date",
  Description:     "For example: 2006 or 19/07/2010",
}
```

- In the template file within your service, reference the
  `input-date.tmpl` file:

```tmpl
<div>Some html...</div>
{{ template "partials/input-date" .PublicationDate }}
<div>Some more html</div>
```

### Localisation

All translations live in `assets/locales/core.<language>.toml` and
are prefixed with `InputDate`.

## BackTo

To instantiate the 'back to' UI component in your service:

- In the `mapper.go` file in your service, populate the relevant fields:

```go
p.BackTo = coreModel.BackTo{
  Text: coreModel.Localisation{
    LocaleKey: "BackToContents",
    Plural:    4,
  },
  AnchorFragment: "toc",
}
```

- In the template file within your service, reference the
  `back-to.tmpl` file:

```tmpl
<div>Some html...</div>
{{ template "partials/back-to" . }}
<div>Some more html</div>
```

## Table

To instantiate the [Table](https://service-manual.ons.gov.uk/design-system/components/table) UI component in your service needs an entry in the mapper and a template. However as there are many variations of this component this will be expanded on in the next section.

- Basic `mapper.go` example:

```go
p.ExampleTable = coreModel.Table{
  Caption: "Example caption",
  TableHeaders: []coreModel.TableHeader{
    {
      Value: "Column A",
    },
  },
  TableRows: []coreModel.TableRow{
    {
      TableData: []coreModel.TableData{
        {
          Value: "Cell A",
        },
      },
    },
  },
}
```

- In the template file within your service, reference the `partials/table.tmpl` file:

```tmpl
{{ template "partials/table" .ExampleTable }}
```

### Variations

The Table component supports many variations. These are optional, and can be
combined:

- "compact"
- "responsive"
- "scrollable"
- "sortable"
- "row-hover"

All variations are configured through the mapper.

#### Table with footer

```go
p.ExampleTable = return coreModel.Table{
  Caption: "A basic table with a footer",
  TableHeaders: []coreModel.TableHeader{
    {
      Value: "Column A",
    },
    {
      Value: "Column B",
    },
    {
      Value: "Column C",
    },
  },
  TableRows: []coreModel.TableRow{
    {
      TableData: []coreModel.TableData{
        {
          Value: "Cell A1",
        },
        {
          Value: "Cell B1",
        },
        {
          Value: "Cell C1",
        },
      },
    },
    {
      TableData: []coreModel.TableData{
        {
          Value: "Cell A2",
        },
        {
          Value: "Cell B2",
        },
        {
          Value: "Cell C2",
        },
      },
    },
  },
  TableFooters: []coreModel.TableFooter{
    {
      Value: "Column summary A",
    },
    {
      Value: "Column summary B",
    },
    {
      Value: "Column summary C",
    },
  },
}
```

#### Compact table

```go
p.ExampleTable = coreModel.Table{
  Variants: []string{"compact", "row-hover"},
  Caption:  "A compacted table with a large number of columns",
  TableHeaders: []coreModel.TableHeader{
    {
      Value: "Column A",
    },
    {
      Value: "Column B",
    },
    {
      Value: "Column C",
    },
    {
      Value: "Column D",
    },
    {
      Value: "Column E",
    },
    {
      Value: "Column F",
    },
  },
  TableRows: []coreModel.TableRow{
    {
      TableData: []coreModel.TableData{
        {
          Value: "Cell A1",
        },
        {
          Value: "Cell B1",
        },
        {
          Value: "Cell C1",
        },
        {
          Value: "Cell D1",
        },
        {
          Value: "Cell E1",
        },
        {
          Value: "Cell F1",
        },
      },
    },
    {
      TableData: []coreModel.TableData{
        {
          Value: "Cell A2",
        },
        {
          Value: "Cell B2",
        },
        {
          Value: "Cell C2",
        },
        {
          Value: "Cell D2",
        },
        {
          Value: "Cell E1",
        },
        {
          Value: "Cell F1",
        },
      },
    },
  },
}
```

#### Numeric table

```go
p.ExampleTable = coreModel.Table{
  Caption: "A basic table with numeric values",
  TableHeaders: []coreModel.TableHeader{
    {
      Value: "Country",
    },
    {
      Value:   "Population mid-2020",
      Numeric: true,
    },
    {
      Value:   "% change 2019 to 2020",
      Numeric: true,
    },
  },
  TableRows: []coreModel.TableRow{
    {
      TableData: []coreModel.TableData{
        {
          Value: "England",
        },
        {
          Value:   "67,081,000",
          Numeric: true,
        },
        {
          Value:   "0.43",
          Numeric: true,
        },
      },
    },
    {
      TableData: []coreModel.TableData{
        {
          Value: "Northern Ireland",
        },
        {
          Value:   "1,896,000",
          Numeric: true,
        },
        {
          Value:   "0.10",
          Numeric: true,
        },
      },
    },
    {
      TableData: []coreModel.TableData{
        {
          Value: "Scotland",
        },
        {
          Value:   "5,466,000",
          Numeric: true,
        },
        {
          Value:   "0.05",
          Numeric: true,
        },
      },
    },
    {
      TableData: []coreModel.TableData{
        {
          Value: "Wales",
        },
        {
          Value:   "3,170,000",
          Numeric: true,
        },
        {
          Value:   "0.53",
          Numeric: true,
        },
      },
    },
  },
}
```

#### Responsive table

```go
p.ExampleTable = coreModel.Table{
  Variants: []string{"responsive"},
  Caption:  "Responsive table with stacked rows for small viewports",
  TableHeaders: []coreModel.TableHeader{
    {
      Value: "Country",
    },
    {
      Value: "Highest mountain",
    },
    {
      Value:   "Height in metres",
      Numeric: true,
    },
  },
  TableRows: []coreModel.TableRow{
    {
      TableData: []coreModel.TableData{
        {
          Value: "Scotland",
          Data:  "Country",
        },
        {
          Value: "Ben Nevis",
          Data:  "Highest mountain",
        },
        {
          Value:   "1,345",
          Data:    "Height",
          Numeric: true,
        },
      },
    },
    {
      TableData: []coreModel.TableData{
        {
          Value: "Wales",
          Data:  "Country",
        },
        {
          Value: "Snowdon",
          Data:  "Highest mountain",
        },
        {
          Value:   "1,085",
          Data:    "Height",
          Numeric: true,
        },
      },
    },
    {
      TableData: []coreModel.TableData{
        {
          Value: "England",
          Data:  "Country",
        },
        {
          Value: "Scafell Pike",
          Data:  "Highest mountain",
        },
        {
          Value:   "978",
          Data:    "Height",
          Numeric: true,
        },
      },
    },
    {
      TableData: []coreModel.TableData{
        {
          Value: "Northern Ireland",
          Data:  "Country",
        },
        {
          Value: "Slieve Donard",
          Data:  "Highest mountain",
        },
        {
          Value:   "850",
          Data:    "Height",
          Numeric: true,
        },
      },
    },
  },
}
```

#### Scrollable table

```go
p.ExampleTable = coreModel.Table{
  Variants:  []string{"scrollable"},
  Caption:   "Scrollable table",
  AriaLabel: "There are 7 columns in this table. Some of the table may be off screen. Scroll or drag horizontally to bring into view.",
  TableHeaders: []coreModel.TableHeader{
    {
      Value: "ID",
    },
    {
      Value: "Title",
    },
    {
      Value: "Abbreviation",
    },
    {
      Value: "Legal basis",
    },
    {
      Value: "Frequency",
    },
    {
      Value: "Date",
    },
    {
      Value: "Status",
    },
  },
  TableRows: []coreModel.TableRow{
    {
      TableData: []coreModel.TableData{
        {
          Value: "023",
        },
        {
          Value: "Monthly Business Survey - Retail Sales Index",
        },
        {
          Value: "RSI",
        },
        {
          Value: "Statistics of Trade Act 1947",
        },
        {
          Value: "Monthly",
        },
        {
          Value: "20 Jan 2018",
        },
        {
          Value: "<span class='ons-status ons-status--success'>Ready</span>",
        },
      },
    },
    {
      TableData: []coreModel.TableData{
        {
          Value: "112",
        },
        {
          Value: "Annual Inward Foreign Direct Investment Survey",
        },
        {
          Value: "AIFDI",
        },
        {
          Value: "Statistics of Trade Act 1947",
        },
        {
          Value: "Annually",
        },
        {
          Value: "26 Feb 2018",
        },
        {
          Value: "<span class='ons-status ons-status--dead'>Not ready</span>",
        },
      },
    },
    {
      TableData: []coreModel.TableData{
        {
          Value: "332",
        },
        {
          Value: "Business Register and Employment Survey",
        },
        {
          Value: "BRES",
        },
        {
          Value: "Statistics of Trade Act 1947",
        },
        {
          Value: "Annually",
        },
        {
          Value: "23 Jan 2013",
        },
        {
          Value: "<span class='ons-status ons-status--info'>In progress</span>",
        },
      },
    },
    {
      TableData: []coreModel.TableData{
        {
          Value: "654",
        },
        {
          Value: "Quartely Survey of Building Materials Sand and Gravel",
        },
        {
          Value: "QBMS",
        },
        {
          Value: "Statistics of Trade Act 1947 - BEIS",
        },
        {
          Value: "Quartely",
        },
        {
          Value: "24 Jan 2015",
        },
        {
          Value: "<span class='ons-status ons-status--error'>Issue</span>",
        },
      },
    },
    {
      TableData: []coreModel.TableData{
        {
          Value: "765",
        },
        {
          Value: "Monthly Survey of Building Materials Concrete Building Blocks",
        },
        {
          Value: "MSBB",
        },
        {
          Value: "Voluntary",
        },
        {
          Value: "Monthly",
        },
        {
          Value: "25 Jan 2014",
        },
        {
          Value: "<span class='ons-status ons-status--success'>Ready</span>",
        },
      },
    },
  },
}
```

#### Sortable table

```go
p.ExampleTable = coreModel.Table{
  Variants: []string{"sortable"},
  Caption:  "Javascript enhanced sortable table",
  SortBy:   "Sort by",
  AriaAsc:  "ascending",
  AriaDesc: "descending",
  TableHeaders: []coreModel.TableHeader{
    {
      Value: "ID",
    },
    {
      Value: "Title",
    },
    {
      Value: "Abbreviation",
    },
    {
      Value: "Legal basis",
    },
    {
      Value: "Frequency",
    },
    {
      Value: "Date",
    },
    {
      Value: "Status",
    },
  },
  TableRows: []coreModel.TableRow{
    {
      TableData: []coreModel.TableData{
        {
          Value: "023",
        },
        {
          Value: "Monthly Business Survey - Retail Sales Index",
        },
        {
          Value: "RSI",
        },
        {
          Value: "Statistics of Trade Act 1947",
        },
        {
          Value:    "Monthly",
          DataSort: "1",
        },
        {
          Value:    "20 Jan 2018",
          DataSort: "2018-01-20",
        },
        {
          Value:    "<span class='ons-status ons-status--success'>Ready</span>",
          DataSort: "0",
        },
      },
    },
    {
      TableData: []coreModel.TableData{
        {
          Value: "112",
        },
        {
          Value: "Annual Inward Foreign Direct Investment Survey",
        },
        {
          Value: "AIFDI",
        },
        {
          Value: "Statistics of Trade Act 1947",
        },
        {
          Value:    "Annually",
          DataSort: "12",
        },
        {
          Value:    "26 Feb 2018",
          DataSort: "2018-02-26",
        },
        {
          Value:    "<span class='ons-status ons-status--dead'>Not ready</span>",
          DataSort: "1",
        },
      },
    },
    {
      TableData: []coreModel.TableData{
        {
          Value: "332",
        },
        {
          Value: "Business Register and Employment Survey",
        },
        {
          Value: "BRES",
        },
        {
          Value: "Statistics of Trade Act 1947",
        },
        {
          Value:    "Annually",
          DataSort: "12",
        },
        {
          Value:    "23 Jan 2013",
          DataSort: "2013-01-23",
        },
        {
          Value:    "<span class='ons-status ons-status--info'>In progress</span>",
          DataSort: "2",
        },
      },
    },
    {
      TableData: []coreModel.TableData{
        {
          Value: "654",
        },
        {
          Value: "Quartely Survey of Building Materials Sand and Gravel",
        },
        {
          Value: "QBMS",
        },
        {
          Value: "Statistics of Trade Act 1947 - BEIS",
        },
        {
          Value:    "Quartely",
          DataSort: "3",
        },
        {
          Value:    "24 Jan 2015",
          DataSort: "2015-01-24",
        },
        {
          Value:    "<span class='ons-status ons-status--error'>Issue</span>",
          DataSort: "3",
        },
      },
    },
    {
      TableData: []coreModel.TableData{
        {
          Value: "765",
        },
        {
          Value: "Monthly Survey of Building Materials Concrete Building Blocks",
        },
        {
          Value: "MSBB",
        },
        {
          Value: "Voluntary",
        },
        {
          Value:    "Monthly",
          DataSort: "1",
        },
        {
          Value:    "25 Jan 2014",
          DataSort: "2014-01-25",
        },
        {
          Value:    "<span class='ons-status ons-status--success'>Ready</span>",
          DataSort: "0",
        },
      },
    },
  },
}
```

### Forms

Forms can be embedded in any data cell.

```go
linkForm := coreModel.TableForm{
  Method: "POST",
  Action: "#",
  Button: coreModel.TableFormButton{
    Text:  "Link",
    Id:    "form-link",
    Name:  "form-link-name",
    Value: "form-link-value",
    Url:   "https://example.com/foo",
  },
  HiddenFormFields: []coreModel.TableFormHiddenField{
    {
      Name:  "hidden-name-1",
      Value: "hidden-value-1",
    },
    {
      Name:  "hidden-name-2",
      Value: "hidden-value-2",
    },
  },
}

buttonForm := coreModel.TableForm{
  Method: "POST",
  Action: "#",
  Button: coreModel.TableFormButton{
    Text:  "Button",
    Id:    "form-button-id",
    Name:  "form-button-name",
    Value: "form-button-value",
  },
  HiddenFormFields: []coreModel.TableFormHiddenField{
    {
      Name:  "hidden-name-1",
      Value: "hidden-value-1",
    },
    {
      Name:  "hidden-name-2",
      Value: "hidden-value-2",
    },
  },
}

p.ExampleTable = coreModel.Table{
  Caption: "A basic table with a caption",
  Id:      "basic-table",
  TableHeaders: []coreModel.TableHeader{
    {
      Value: "Column A",
    },
    {
      Value: "Column B",
    },
    {
      Value: "Column C",
    },
  },
  TableRows: []coreModel.TableRow{
    {
      TableData: []coreModel.TableData{
        {
          Value: "Cell A1",
          Name:  "cell-name",
        },
        {
          Value: "Cell B1",
        },
        {
          Value: "Cell C1",
        },
      },
    },
    {
      TableData: []coreModel.TableData{
        {
          Value: "Cell A2",
        },
        {
          Form: &linkForm,
        },
        {
          Form: &buttonForm,
        },
      },
    },
  },
}
```

### Localisation

All translations live in `assets/locales/core.<language>.toml` and
are prefixed with `Table`.

## Correct errors

To instantiate the [correct errors](https://service-manual.ons.gov.uk/design-system/patterns/correct-errors) UI component in your service:

- In the `mapper.go` file in your service, populate the relevant fields
  e.g.

```go
p.Page.Error = coreModel.Error{
   Title: p.Metadata.Title,
   ErrorItems: []coreModel.ErrorItem{
    {
      // You can either use a localisation key or populate the `Text` string
      Description: coreModel.Localisation{
        LocaleKey: "CoverageSelectDefault",
        Plural:    1,
      },
      // URL can be an in-page link or a different page
      URL: "#coverage-error",
    },
   },
   Language: lang,
  }
```

- If there is more than one error to correct, add further `ErrorItem` types. By doing so, the view will automatically toggle to an ordered numbered list
  e.g.

```go
p.Page.Error = coreModel.Error{
   Title: p.Metadata.Title,
   ErrorItems: []coreModel.ErrorItem{
    {
      // You can either use a localisation key or populate the `Text` string
      Description: coreModel.Localisation{
        LocaleKey: "CoverageSelectDefault",
        Plural:    1,
      },
      // URL can be an in-page link or a different page
      URL: "#coverage-error",
    },
    {
      Description: coreModel.Localisation{
        Text: "A different error",
      },
      // URL can be an in-page link or a different page
      URL: "#another-error",
    },
   },
   Language: lang,
  }
```

- In the template file within your service, reference the `error-summary.tmpl` file below the breadcrumbs, before the `<h1>` and pass in the `error` struct
  e.g.

```tmpl
{{ template "partials/breadcrumb" . }}
{{ if .Page.Error.Title }}
    {{ template "partials/error-summary" .Page.Error }}
{{ end}}
<h1>The header</h1>
```

- You will need to manually add the [error panel](https://service-manual.ons.gov.uk/design-system/components/panel/#error-details) yourself but it will resemble something like the example. This is due to the panel adding two additional containing `<div>` elements and different input fields have different variants on implementing the error. Alternatively, if your fields are not complex then use one of the [fields](#fields) patterns which provide the necessary error panel surrounding the input field.
  e.g.

```tmpl
{{ if .Page.Error.Title }}
<div class="ons-panel ons-panel--error ons-panel--no-title" id="coverage-error">
  <span class="ons-u-vh">
      {{ localise "Error" .Language 1 }}:
  </span>
  <div class="ons-panel__body">
      <p class="ons-panel__error">
          <strong>Enter a value</strong>
      </p>
{{ end }}
...Input field...
{{ if .Page.Error.Title }}
  </div>
</div>
{{ end }}
```

## Fields

The field components provides a consistent way of adding a simple valid html input element with corresponding field validation by populating relevant properties on the model.

Fields available:

- [text](#text-input)
- [textarea](#textarea)
- [radios fieldset](#radios-fieldset)

The inputs contained within the fields share a common model `input.go`, this allows [attributes](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes) to be set in a consistent way. However, there are circumstances where the attributes are not permitted and subsequently not rendered. Use the table below as a guide:

| Attribute                                                                                 | Checkbox | Radio | Text | Textarea |
| ----------------------------------------------------------------------------------------- | -------- | ----- | ---- | -------- |
| [Autocomplete](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/autocomplete) | No       | No    | Yes  | Yes      |
| ID (mandatory)                                                                            | Yes      | Yes   | Yes  | Yes      |
| Checked                                                                                   | Yes      | Yes   | No   | No       |
| Disabled                                                                                  | Yes      | Yes   | Yes  | Yes      |
| Name (mandatory)                                                                          | Yes      | Yes   | Yes  | Yes      |
| Value                                                                                     | Yes      | Yes   | Yes  | Yes      |

### Text input

To instatiate the [text input](https://service-manual.ons.gov.uk/design-system/components/input) UI component in your service:

- In the `mapper.go` file in your service, populate the relevant fields
  e.g.

```go
p.TextInput = core.TextField{
   Input: core.Input{
    ID:    "text-field",
    Name:  "text",
    Value: "", // Only required to rebind to the model after validation
    Label: core.Localisation{
     LocaleKey: "LocaleKey",
     Plural:    1,
    },
   },
  },
```

- The above snippet will render the following html:

```html
<div class="ons-field">
  <label class="ons-label" for="name-field" id="name-field-label">
    Locale lookup label
  </label>
  <input
    class="ons-input ons-input--text ons-input-type__input"
    type="text"
    id="text-field"
    value=""
    name="text"
  />
</div>
```

#### Text input validation

If the field fails form validation, users need to be given the opportunity to [correct the error](https://service-manual.ons.gov.uk/design-system/patterns/correct-errors). The following mapped fields will render the field with a validation error:

```go
p.AnotherInput = core.TextField{
  Input: core.Input{
  ID:    "another-field",
  Name:  "another",
  Label: core.Localisation{
    Text: "Label text",
  },
  Value: "user entered value",
  },
  ValidationErr: core.ValidationErr{ // Use in conjunction with the 'correct errors' pattern
  HasValidationErr: true,
  ErrorItem: core.ErrorItem{
    Description: core.Localisation{
    LocaleKey: "LocaleKey",
    Plural:    1,
    },
    ID: "another-field-error", // Linked from the 'correct errors' pattern
    },
  },
}
```

- The above snippet will render the following html

```html
<div
  class="ons-panel ons-panel--error ons-panel--no-title"
  id="another-field-error"
>
  <span class="ons-u-vh">Error: </span>
  <div class="ons-panel__body">
    <p class="ons-panel__error">
      <strong>A locale lookup meaningful message</strong>
    </p>
    <div class="ons-field">
      <label class="ons-label" for="another-field" id="another-field-label">
        Label text
      </label>
      <input
        class="ons-input ons-input--text ons-input-type__input"
        type="text"
        id="another-field"
        value="user entered value"
        name="another"
      />
    </div>
  </div>
</div>
```

### Textarea

To instatiate the [textarea input](https://service-manual.ons.gov.uk/design-system/components/textarea) UI component in your service:

- In the `mapper.go` file in your service, populate the relevant fields
  e.g.

```go
p.TextareaField = core.TextareaField{
  Input: core.Input{
   Description: core.Localisation{
    LocaleKey: "DescriptionLocaleKey", // Can be text instead of a locale
    Plural:    1,
   },
   ID: "textarea-field",
   Label: core.Localisation{
    LocaleKey: "LabelLocaleKey", // Can be text instead of a locale
    Plural:    1,
   },
   Language: lang, // required for the localisation
   Name:     "textarea",
  },
 }
```

- The above snippet will render the following html

```html
<div class="ons-field">
  <label
    class="ons-label ons-label--with-description"
    aria-describedby="textarea-field-hint"
    for="textarea-field"
    id="textarea-field-label"
  >
    The label text
  </label>
  <span
    id="textarea-field-hint"
    class="ons-label__description ons-input--with-description"
  >
    Meaningful hint text
  </span>
  <textarea
    id="textarea-field"
    class="ons-input ons-input--textarea"
    name="textarea"
    rows="8"
  >
  </textarea>
</div>
```

#### Textarea validation

If the field fails form validation, users need to be given the opportunity to [correct the error](https://service-manual.ons.gov.uk/design-system/patterns/correct-errors). The following mapped fields will render the field with a validation error:

```go
p.TextareaField = core.TextareaField{
  Input: core.Input{
   // As above (omitted for brevity)
   Value: " ", // user entered whitespace
  },
  ValidationErr: core.ValidationErr{
   HasValidationErr: true,
   ErrorItem: core.ErrorItem{
    Description: core.Localisation{
     LocaleKey: "TextareaErrorLocale",
     Plural:    1,
    },
    ID: "textarea-error",
   },
  },
 }
```

- The above snippet will render the following html:

```html
<div class="ons-panel ons-panel--error ons-panel--no-title" id="textarea-error">
  <span class="ons-u-vh">Error: </span>
  <div class="ons-panel__body">
    <p class="ons-panel__error">
      <strong>Enter some text</strong>
    </p>
    <div class="ons-field">
      <label
        class="ons-label ons-label--with-description"
        aria-describedby="textarea-field-hint"
        for="textarea-field"
        id="textarea-field-label"
      >
        The label text
      </label>
      <span
        id="textarea-field-hint"
        class="ons-label__description ons-input--with-description"
      >
        Meaningful hint text
      </span>
      <textarea
        id="textarea-field"
        class="ons-input ons-input--textarea"
        name="textarea"
        rows="8"
      >
      </textarea>
    </div>
  </div>
</div>
```

### Radios fieldset

To instatiate a [radios fieldset](https://service-manual.ons.gov.uk/design-system/components/radios) without a border that allows a conditionally revealed text input UI component in your service:

- In the `mapper.go` file in your service, populate the relevant fields
  e.g.

```go
p.Radios = core.RadioFieldset{
  Legend: core.Localisation{
   LocaleKey: "LegendLocaleKey", // Can be text instead of a locale
   Plural:    1,
  },
  Radios: []core.Radio{
   {
    Input: core.Input{
     ID:        "radio-one",
     IsChecked: true, // Can be false or omitted (= false)
     Label: core.Localisation{
      LocaleKey: "LabelOneLocaleKey", // Can be text instead of a locale
      Plural:    1,
     },
     Name:  "radios",
     Value: "radio one value",
    },
   },
   {
    Input: core.Input{
     ID:        "radio-two",
     Label: core.Localisation{
      Text: "Label two text"
     },
     Name:  "radios",
     Value: "radio two value",
    },
    OtherInput: core.Input{
     ID:    "other-input",
     Name:  "other",
     Label: core.Localisation{
      Text: "other input label",
     },
    },
   },
  },
 }
```

- The above snippet will render the following html

```html
<fieldset class="ons-fieldset">
  <legend class="ons-fieldset__legend">A locale legend</legend>
  <div class="ons-radios__items">
    <div class="ons-radios__item ons-radios__item--no-border ons-u-fw">
      <div class="ons-radio ons-radio--no-border">
        <input
          type="radio"
          id="radio-one"
          name="radios"
          class="ons-radio__input ons-js-radio"
          value="radio one value"
          checked
        />
        <label class="ons-radio__label" for="radio-one" id="radio-one-label">
          Radio one locale label
        </label>
      </div>
    </div>
    <div class="ons-radios__item ons-radios__item--no-border ons-u-fw">
      <div class="ons-radio ons-radio--no-border">
        <input
          type="radio"
          id="radio-two"
          name="radios"
          class="ons-radio__input ons-js-radio"
          value="radio two value"
        />
        <label class="ons-radio__label" for="radio-two" id="radio-two-label">
          Label two text
        </label>
        <div class="ons-radio__other">
          <label class="ons-label" for="other-input" id="other-input-label">
            Other input label
          </label>
          <input
            class="ons-input ons-input--text ons-input-type__input"
            type="text"
            id="other-input"
            value=""
            name="other"
          />
        </div>
      </div>
    </div>
  </div>
</fieldset>
```

#### Radio fieldset validation

If the field fails form validation, users need to be given the opportunity to [correct the error](https://service-manual.ons.gov.uk/design-system/patterns/correct-errors). The following mapped fields will render the field with a validation error:

- In the `mapper.go` file in your service, populate the relevant fields
  e.g.

```go
p.Radios = core.RadioFieldset{
  Legend: core.Localisation{
   // As above (omitted for brevity)
  },
  Radios: []core.Radio{
   // As above (omitted for brevity)
  },
  ValidationErr: core.ValidationErr{
   HasValidationErr: true,
   ErrorItem: core.ErrorItem{
    Description: core.Localisation{
     Text: "Radio error text",
    },
    ID: "radio-error",
   },
  },
 }
```

- The above snippet will render the following html

```html
<div class="ons-panel ons-panel--error ons-panel--no-title" id="radio-error">
  <span class="ons-u-vh">Error: </span>
  <div class="ons-panel__body">
    <p class="ons-panel__error">
      <strong>Radio error text</strong>
    </p>
    <fieldset class="ons-fieldset">
      <legend class="ons-fieldset__legend">A locale legend</legend>
      <div class="ons-radios__items">
        <div class="ons-radios__item ons-radios__item--no-border ons-u-fw">
          <div class="ons-radio ons-radio--no-border">
            <input
              type="radio"
              id="radio-one"
              name="radios"
              class="ons-radio__input ons-js-radio"
              value="radio one value"
              checked
            />
            <label
              class="ons-radio__label"
              for="radio-one"
              id="radio-one-label"
            >
              Radio one locale label
            </label>
          </div>
        </div>
        <div class="ons-radios__item ons-radios__item--no-border ons-u-fw">
          <div class="ons-radio ons-radio--no-border">
            <input
              type="radio"
              id="radio-two"
              name="radios"
              class="ons-radio__input ons-js-radio"
              value="radio two value"
            />
            <label
              class="ons-radio__label"
              for="radio-two"
              id="radio-two-label"
            >
              Label two text
            </label>
            <div class="ons-radio__other">
              <label class="ons-label" for="other-input" id="other-input-label">
                Other input label
              </label>
              <input
                class="ons-input ons-input--text ons-input-type__input"
                type="text"
                id="other-input"
                value=""
                name="other"
              />
            </div>
          </div>
        </div>
      </div>
    </fieldset>
  </div>
</div>
```
