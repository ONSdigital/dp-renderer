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

## Collapsible

To instatiate the [collapsible](https://ons-design-system.netlify.app/components/collapsible/) UI component in your service:

- In the `mapper.go` file in your service, populate the relevant fields
e.g.

```go
p.Collapsible = coreModel.Collapsible{
  // You can either use a localisation key or populate the `Text` string
  Title: coreModel.Localisation{
    LocaleKey: "VariablesExplanation",
    Plural:    4,
  },
  CollapsibleItems: []coreModel.CollapsibleItem{
   {
    Subheading: "This is a subheading",
    Content:    []string{"a string"},
   },
   {
    Subheading: "This is another subheading",
    Content:    []string{"another string", "and another"},
   },
  },
 }
```

- In the template file within your service, reference the `collapsible.tmpl` file
e.g.

```tmpl
<div>Some html...</div>
{{ template "partials/collapsible" . }}
<div>Some more html</div>
```

## Table Of Contents

To instatiate the [table-of-contents](https://ons-design-system.netlify.app/components/table-of-contents/) UI component in your service:

- In the `mapper.go` file in your service, populate the relevant fields:

```go
page.TableOfContents = TableOfContents{
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

To instatiate the [pagination](https://ons-design-system.netlify.app/components/pagination/) UI component in your service:

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

To instatiate the [Dates](https://ons-design-system.netlify.app/patterns/dates/) UI pattern in your service:

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

To instatiate the 'back to' UI component in your service:

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

To instatiate the [Table](https://ons-design-system.netlify.app/components/table/) UI component in your service needs an entry in the mapper and a template. However as there are many variations of this component this will be expanded on in the next section.

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
