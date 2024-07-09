GO_VERSION := 1.22
TAG := $(shell git describe --abbrev=0 --tags --always)
HASH := $(shell git rev-parse HEAD)
DATE := $(shell date +%Y-%m-%d.%H:%M:%S)
LDFLAGS := -w -X github.com/luzeduardo/shipping-go/handlers.hash=$(HASH) \
	-X github.com/luzeduardo/shipping-go/handlers.tag=$(TAG) \
	-X github.com/luzeduardo/shipping-go/handlers.date=$(DATE)

build:
	go build -ldflags "$(LDFLAGS)" -o api cmd/main.go

test:
	go test -v -short ./...

coverage:
	go test ./... -cover

report:
	touch coverage.out && go tool cover -html=coverage.out -o cover.xhtml
	
