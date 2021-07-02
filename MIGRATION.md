# Migrating from `dp-frontend-renderer` to `dp-renderer`

## Migrating assets and models

The frontend service is responsible for setting up the assets binary and source path.

You should also store all assets and models within the application itself with the following structure:

```md
.
├── asset                    # relevant templates & localisations from dp-frontend-renderer
│   ├── templates          
│   ├── localisations  
|   |   ├──  service.en.toml
|   └── └──  service.cy.toml
└── model                    # relevant models from dp-frontend-models
```

## Updating the Makefile

In order for `dp-renderer` to work correctly once the assets have been migrated over, we use `go-bindata` to generate a combined assets source file.

Update the frontend app's `Makefile` with the following new commands so that `go-bindata` will generate this file:

```Makefile
.PHONY: get-renderer-version
get-renderer-version:
 $(eval APP_RENDERER_VERSION=$(shell grep "github.com/ONSdigital/dp-renderer" go.mod | cut -d ' ' -f2 ))

.PHONY: fetch-renderer-lib
fetch-renderer-lib: get-renderer-version
 $(eval CORE_ASSETS_PATH = $(shell go get github.com/ONSdigital/dp-renderer@$(APP_RENDERER_VERSION) && go list -f '{{.Dir}}' -m github.com/ONSdigital/dp-renderer))

.PHONY: generate-debug
generate-debug: fetch-renderer-lib
 # fetch the renderer library and build the dev version
 cd assets; go run github.com/kevinburke/go-bindata/go-bindata -prefix $(CORE_ASSETS_PATH)/assets -debug -o data.go -pkg assets locales/... templates/... $(CORE_ASSETS_PATH)/assets/locales/... $(CORE_ASSETS_PATH)/assets/templates/...
 { echo "// +build debug\n"; cat assets/data.go; } > assets/debug.go.new
 mv assets/debug.go.new assets/data.go

.PHONY: generate-prod
generate-prod: fetch-renderer-lib
 # fetch the renderer library and build the prod version
 cd assets; go run github.com/kevinburke/go-bindata/go-bindata -prefix $(CORE_ASSETS_PATH)/assets -o data.go -pkg assets locales/... templates/... $(CORE_ASSETS_PATH)/assets/locales/... $(CORE_ASSETS_PATH)/assets/templates/...
 { echo "// +build production\n"; cat assets/data.go; } > assets/data.go.new
 mv assets/data.go.new assets/data.go
```

Due to having distributed assets that are combined with `go-bindata`, we require `get-renderer-version` and `fetch-renderer-version` to ensure the version of `dp-renderer` as specified in `go.mod` is used.

The existing `build` and `debug` tasks should then be updated to use the relevant `generate-` command as a prerequisite:

```Makefile
.PHONY: build
build: generate-prod

.PHONY: debug
debug: generate-debug
```

## Updating `config.go`

`config.go` should be updated to include two new properties: `PatternLibraryAssetsPath` and `SiteDomain`. These values are used when instantiating the `Render` struct.

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

Once the combined assets generation, `Makefile` and `config.go` changes are made, you can now start using `dp-renderer`. Check the [usage](/README.md) notes in the readme for more information.
