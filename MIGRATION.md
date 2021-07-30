# Migrating from `dp-frontend-renderer` to `dp-renderer`

This doc is a guide for migrating an existing frontend service from using `dp-frontend-renderer` to the `dp-renderer` library instead.

## Context: why migrate?

As the web team has grown, the workflow around `dp-frontend-renderer` has become more cumbersome. In some instances, developers across teams may be working on new frontend features, such as new pages for the website, that require updates across three repositories: `dp-frontend-renderer`, `dp-frontend-models` and the relevant frontend controller. This risks bottlenecks and greater external dependency across teams when it comes to releases.

`dp-renderer` mitigates a lot of these workflow overheads by providing greater domain encapsulation for our frontend services. Instead of having models and assets in separate repositories, we can instead house them all in the specific frontend service, with `dp-renderer` being imported to handle template rendering duties instead of requiring an additional network request to do this for us.

Consequently, `dp-renderer` simplifies our frontend architecture and reduces pressure on teams to coordinate work in this space.

## Migrating assets and models

The frontend service is responsible for setting up the assets binary and source path.

You should also store all assets and models within the service itself with the following structure:

```md
.
├── assets                   # relevant templates & localisations from dp-frontend-renderer
│   ├── templates          
│   ├── localisations  
|   |   ├──  service.en.toml
|   └── └──  service.cy.toml
└── model                    # relevant models from dp-frontend-models
```

## Updating the Makefile

For `dp-renderer` to work correctly once the assets have been migrated over, we use `go-bindata` to generate a combined assets source file.

Update the frontend service's `Makefile` with the following new commands so that `go-bindata` will generate this file:

```Makefile
.PHONY: fetch-renderer-lib
fetch-renderer-lib:
ifeq ($(LOCAL_DP_RENDERER_IN_USE), 1)
 $(eval CORE_ASSETS_PATH = $(shell grep -w "\"github.com/ONSdigital/dp-renderer\" =>" go.mod | awk -F '=> ' '{print $$2}'))
else
 $(eval APP_RENDERER_VERSION=$(shell grep "github.com/ONSdigital/dp-renderer" go.mod | cut -d ' ' -f2 ))
 $(eval CORE_ASSETS_PATH = $(shell go get github.com/ONSdigital/dp-renderer@$(APP_RENDERER_VERSION) && go list -f '{{.Dir}}' -m github.com/ONSdigital/dp-renderer))
endif


.PHONY: generate-debug
generate-debug: fetch-renderer-lib
 cd assets; go run github.com/kevinburke/go-bindata/go-bindata -prefix $(CORE_ASSETS_PATH)/assets -debug -o data.go -pkg assets locales/... templates/... $(CORE_ASSETS_PATH)/assets/locales/... $(CORE_ASSETS_PATH)/assets/templates/...
 { echo "// +build debug\n"; cat assets/data.go; } > assets/debug.go.new
 mv assets/debug.go.new assets/data.go

.PHONY: generate-prod
generate-prod: fetch-renderer-lib
 cd assets; go run github.com/kevinburke/go-bindata/go-bindata -prefix $(CORE_ASSETS_PATH)/assets -o data.go -pkg assets locales/... templates/... $(CORE_ASSETS_PATH)/assets/locales/... $(CORE_ASSETS_PATH)/assets/templates/...
 { echo "// +build production\n"; cat assets/data.go; } > assets/data.go.new
 mv assets/data.go.new assets/data.go
```

Due to having distributed assets that are combined with `go-bindata`, we require the `get-renderer-version` and `fetch-renderer-version` tasks to ensure the version of `dp-renderer` as specified in `go.mod` is used.

The existing `build` and `debug` tasks should then be updated to use the relevant `generate-` command as a prerequisite:

```Makefile
.PHONY: build
build: generate-prod

.PHONY: debug
debug: generate-debug
```

## Updating `config.go`

`config.go` should be updated to include three new properties: `PatternLibraryAssetsPath`, `SiteDomain`, and `Debug`.

You will also need to add additional logic to `config.go` to handle the path for pattern library assets when running `make debug`. During local development, we point to our local version of the pattern library instead.

Example set up in `config.go`:

```go
type Config struct {
 BindAddr                   string        `envconfig:"BIND_ADDR"`
 Debug                      bool          `envconfig:"DEBUG"`
 APIRouterURL               string        `envconfig:"API_ROUTER_URL"`
 SiteDomain                 string        `envconfig:"SITE_DOMAIN"`
 PatternLibraryAssetsPath   string        `envconfig:"PATTERN_LIBRARY_ASSETS_PATH"`
 GracefulShutdownTimeout    time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
 HealthCheckInterval        time.Duration `envconfig:"HEALTHCHECK_INTERVAL"`
 HealthCheckCriticalTimeout time.Duration `envconfig:"HEALTHCHECK_CRITICAL_TIMEOUT"`
}

var cfg *Config

func Get() (*Config, error) {
 cfg, err := get()
 if err != nil {
  return nil, err
 }

 if cfg.Debug {
  cfg.PatternLibraryAssetsPath = "http://localhost:9000/dist"
 } else {
  cfg.PatternLibraryAssetsPath = "//cdn.ons.gov.uk/sixteens/f80be2c"
 }
 return cfg, nil
}

func get() (*Config, error) {
 if cfg != nil {
  return cfg, nil
 }

 cfg = &Config{
  BindAddr:                   ":24100",
  Debug:                      false,
  APIRouterURL:               "http://localhost:22400",
  SiteDomain:                 "localhost",
  GracefulShutdownTimeout:    5 * time.Second,
  HealthCheckInterval:        30 * time.Second,
  HealthCheckCriticalTimeout: 90 * time.Second,
 }

 return cfg, envconfig.Process("", cfg)
}
```

## Refer to usage guide

At this point you can start implementing `dp-renderer`. Check the [usage](/README.md) notes for more information. Once complete, you can return to this migration doc which will cover changes to the handlers, mapper and `RenderClient` interface.

## Updating `RenderClient` interface

You will need to update the frontend service's `RenderClient` interface in order to implement the new methods that `dp-renderer` exposes.

Before:

```golang
type RenderClient interface {
  Do(string, []byte) ([]byte, error)
}
```

After:

```golang
type RenderClient interface {
  BuildPage(w io.Writer, pageModel interface{}, templateName string)
  NewBasePageModel() model.Page
}
```

The compiler will throw errors due to this change. The following sections will cover how to resolve a number of them in the handler and mapper functions.

## Updating handler and mapper functions

### Handlers

There is a lot of error handling, logging, `Write` and `Marshal` logic that is no longer needed in your handlers.

An example handler:

```golang
func getCookiePreferencePage(w http.ResponseWriter, rendC RenderClient, cp cookies.Policy, isUpdated bool, lang string) {
  // create a new base page model that inject SiteDomain and PatternLibraryAssetsPath into the page struct
  basePage := rendC.NewBasePageModel()

  // Mapper function is updated to accept the base page as an argument
  m := mapper.CreateCookieSettingPage(basePage, cp, isUpdated, lang)

  // send the mapped data, with ResponseWriter and template name defined by the actual template file name (e.g. cookies-preferences.tmpl) to the render lib
  rendC.BuildPage(w, m, "cookies-preferences")
```

### Mappers

To continue with the above example, the only changes made to the mapper are updating imports to include the model package from `dp-renderer`, passing in a new page model argument to the mapper function itself, and then setting that as the `Page` property.

```golang
import (
  "dp-frontend-cookie-controller/model"

  "github.com/ONSdigital/dp-cookies/cookies"

  coreModel "github.com/ONSdigital/dp-renderer/model"
)

func CreateCookieSettingPage(basePage coreModel.Page, policy cookies.Policy, isUpdated bool, lang string) model.CookiesPreference {
  page := model.CookiesPreference{
    Page: basePage,
  }
  // rest of mapper function logic
}
```

Once you have made these updates, you should be able to run `make debug` and see that pages handled by your frontend service are presented without requiring `dp-frontend-renderer`.
