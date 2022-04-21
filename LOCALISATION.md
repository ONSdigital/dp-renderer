# Localisation

Components (template partials) can be localised via
- dp-renderer's `assets/locales`, or
- your frontend's `assets/locales`, or
- string literals passed into the component model

## Localisation model

Wherever a component uses the Localisation model, for example:

```
package model

type Button struct {
  Label Localisation
}
```

The `Label` can be set in one of two ways. Either the `LocaleKey` and `Plural`
are set, or its `Text` can be set.

When the `LocaleKey` and `Plural` are provided, these take precedence over
`Text`.

### LocaleKey and Plural

Providing a `LocaleKey` and `Plural` allows a component template to query the
locale files for a translation.

In your frontend mapper this would look like:

```
page.Button = model.Button{
  Label: model.Localisation{
    LocaleKey: "ButtonClickMe",
    Plural: 1,
  },
}
```

The dp-renderer's locales and your frontend's locales are effectively combined,
sharing a namespace. This makes the components in dp-renderer more easily
reusable because they can refer to locale keys found in the frontend.

The drawback of locales sharing a namespace is that a frontend must take care
not to use a key already present in the dp-renderer's locales.

### Text

Setting `Text` to a string is useful when
- No localisation is needed, or
- The text is already localised (perhaps retrieved from an external source)

In your frontend mapper this would look like:

```
page.Button = model.Button{
  Label: model.Localisation{
    Text: "Click me"
  },
}
```

## Template usage

The `Localisation` model has a `Localise` method that takes care of
choice between rendering the `LocaleKey` or `Text` for you.

Continuing the `Button` example, this component's template might
look like this:

```
<button type="button">
  {{- .Button.Label.Localise .Language -}}
</button>
```
