# dp-renderer

A rendering library for Digital Publishing frontend microservices. `dp-renderer` contains templates, localisations and model structs that are core to all `dp-frontend` services.

`dp-renderer` is intended to be used instead of calling `dp-frontend-renderer` to generate HTML, which consequently means more explicit domain encapsulation of our frontend apps and the removal of a single point of failure within the DP frontend architecture.

## Installation

Other than `dp-renderer` itself, you will need `go-bindata` in order to combine service-specific and `dp-renderer` assets:

- `dp-renderer`: `go get github.com/ONSdigital/dp-renderer`

> You can specify a version of `dp-renderer` by appending a commit ID or semantic version number to this command. E.g., `go get github.com/ONSdigital/dp-renderer@31d8704`

- `go-bindata`: `go get github.com/kevinburke/go-bindata`

## Migrating from `dp-frontend-renderer` and `dp-frontend-models` to using `dp-renderer`

See the [migration](/migration.md) guide for step-by-step details.

## Usage

### Instantiation

Assuming you have `go-bindata` set up to generate the relevant asset helper functions, you can instantiate the renderer with a default client (in this case, the default client is [`unrolled`](https://github.com/unrolled/render)).

```go
rend := render.NewWithDefaultClient(asset.Asset, asset.AssetNames, cfg.PatternLibraryAssetsPath, cfg.SiteDomain)
```

You can also instantiate a `Render` struct without a default client by using `New()`. This requires a rendering client that fulfills the `Renderer` interface to be passed in as well.

```go
rend := render.New(rendereringClient, patternLibPath, siteDomain)
```

### Mapping data and building a page

When mapping data to a page model, you can use `NewBasePageModel` to instantiate a base page model with its `PatternLibraryAssetsPath` and `SiteDomain` properties auto-populated via the `Render` struct:

```go
basePage := rendC.NewBasePageModel()
mappedPageData := mapper.CreateExamplePage(basePage)
```

In order to generate HTML from a page model and template, use `BuildPage`, passing in the `ResponseWriter`, mapped data, and the name of the template:

```go
rend.BuildPage(w, mappedPageData, "name-of-template-file-without-extension")
```

If an error occurs during page build, either because of an incorrect template name or incorrect data mapping, `dp-renderer` will write an error via an `errorResponse` struct.

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

## License

Copyright © 2021, Office for National Statistics (<https://www.ons.gov.uk>)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
