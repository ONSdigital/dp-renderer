# dp-renderer

A rendering library for Digital Publishing frontend microservices. `dp-renderer` contains templates, localisations and model structs that are core to all `dp-frontend` services.

`dp-renderer` is intended to be used instead of calling `dp-frontend-renderer` to generate HTML, which consequently means more explicit domain encapsulation of our frontend apps and the removal of a single point of failure within the DP frontend architecture.

Note: although the `dp-frontend-renderer` is deprecated there is a transition period where updates will be needed in `dp-renderer` and the `dp-frontend-renderer`. See the README on the `dp-frontend-renderer` for the migration status.

## Installation

Other than `dp-renderer` itself, you will need a utility that can combine service-specific and `dp-renderer` assets. We currently use `go-bindata` for this process.

- `dp-renderer`: `go get github.com/ONSdigital/dp-renderer`

> You can specify a version of `dp-renderer` by appending a commit ID or semantic version number to this command. E.g., `go get github.com/ONSdigital/dp-renderer@31d8704`

- `go-bindata`: `go get github.com/kevinburke/go-bindata`

## Migrating from `dp-frontend-renderer` and `dp-frontend-models` to using `dp-renderer`

See the [migration](MIGRATION.md) guide for step-by-step details.

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

### Using the ONS Design System

`dp-renderer` supports use of the [ONS Design System](https://ons-design-system.netlify.app/) on a **per-page basis**. In order to use the ONS Design System, the page mapper function in the calling application should set the version number to the `FeatureFlags.ONSDesignSystemVersion` variable in the core page model struct, like so:

```go
func CreateExamplePage(basePage coreModel.Page) model.ExamplePage {
    p := model.ExamplePage{
        Page: basePage
    }

    // Load in specific version of ONS design system JS and CSS
    p.FeatureFlags.ONSDesignSystemVersion = "37.0.0"
    
    return p
}
```

### Referencing a local instance of dp-renderer in a docker container

If you are running and developing within a docker container that includes references to the dp-renderer, for example, the [cantabular import journey](https://github.com/ONSdigital/dp-compose/tree/main/cantabular-import). Follow these steps to use a local instance:

- Update the `go.mod` file in the relevant service to use the `replace statement` to point to your local dp-renderer instance
e.g.

```go
replace "github.com/ONSdigital/dp-renderer" => "/Users/{yourName}/{yourDirectory}/github.com/ONSdigital/dp-renderer"
```

- Add the volume of your local instance to the service's `.yaml` file in dp-compose
e.g.

To modify the [dp-frontend-dataset-controller](https://github.com/ONSdigital/dp-frontend-dataset-controller) to use a local dp-renderer instance in the cantabular import journey, modify `dp-frontend-dataset-controller.yml` volumes to include your local instance

```yml
version: '3.3'
services:
    dp-frontend-dataset-controller:
        build:
            context: ../../dp-frontend-dataset-controller
            dockerfile: Dockerfile.local
        command:
            - reflex
            - -d
            - none
            - -c
            - ./reflex
        volumes:
            - ../../dp-frontend-dataset-controller:/dp-frontend-dataset-controller
            - /Users/{yourName}/{yourDirectory}/github.com/ONSdigital/dp-renderer:/Users/{yourName}/{yourDirectory}/github.com/ONSdigital/dp-renderer
        ports:
            - 20200
        restart: unless-stopped
        environment:
            BIND_ADDR:                ":20200"
            DOWNLOAD_SERVICE_URL:     "http://dp-download-service:23600"
            API_ROUTER_URL:           "http://dp-api-router:23200/v1"
            DOWNLOADER_URL:           "http://dp-download-service:23400"
            ENABLE_CENSUS_PAGES:      "true"
```

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

## License

Copyright © 2021, Office for National Statistics (<https://www.ons.gov.uk>)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
