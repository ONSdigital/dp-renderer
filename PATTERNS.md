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
