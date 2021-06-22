BINPATH ?= build

BUILD_TIME=$(shell date +%s)
GIT_COMMIT=$(shell git rev-parse HEAD)
VERSION ?= $(shell git tag --points-at HEAD | grep ^v | head -n 1)

LDFLAGS = -ldflags "-X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT) -X main.Version=$(VERSION)"

.PHONY: all
all: audit test build

.PHONY: audit
audit:
	go list -json -m all | nancy sleuth --exclude-vulnerability-file ./.nancy-ignore

.PHONY: build
build:
	go build -tags 'production' $(LDFLAGS) -o $(BINPATH)/dp-renderer

.PHONY: debug
debug:
	go build -tags 'debug' $(LDFLAGS) -o $(BINPATH)/dp-renderer
	HUMAN_LOG=1 DEBUG=1 $(BINPATH)/dp-renderer

.PHONY: generate-prod
generate-prod:
	# build the production version
	cd assets; go run github.com/kevinburke/go-bindata/go-bindata -o data.go -pkg assets locales/...
	{ echo "// +build production\n"; cat assets/data.go; } > assets/data.go.new
	mv assets/data.go.new assets/data.go

.PHONY: test
test: generate-prod
	go test -race -cover -tags 'production' ./...

.PHONY: convey
convey:
	goconvey ./...

