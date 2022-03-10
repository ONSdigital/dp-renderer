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
