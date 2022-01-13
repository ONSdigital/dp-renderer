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
