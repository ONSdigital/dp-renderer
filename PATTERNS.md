# Using patterns or components in your service

## Collapsible

To instatiate the [collapsible](https://ons-design-system.netlify.app/components/collapsible/) UI component in your service:

- In the `mapper.go` file in your service, populate the relevant fields
e.g.

```go
p.Collapsible = coreModel.Collapsible{
  // You can either use a localisation key or populate the `Title` string
  LocaliseKey:       "VariablesExplanation",
  LocalisePluralInt: 4,
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

For example, with 1 and 10 as bookends defined in `FirstAndLastPages`,
setting `PagesToDisplay = { 1, 2, 3, 4 }` closes the gap between
1 (a bookend) and 2:

```
1 2 3 4 ... 10
```

Similarly, setting `PagesToDisplay = { 7, 8, 9 10 }` closes the
gap between 9 and 10 (a bookend):

1 ... 7 8 9 10

When the window defined by `PagesToDisplay` does not overlap with
the bookends defined in `FirstAndLastPages`, the page options are
rendered as:

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
