.PHONY: clean check test build release-test

export GO111MODULE=on

TAG_NAME := $(shell git tag -l --contains HEAD)
SHA := $(shell git rev-parse --short HEAD)
VERSION := $(if $(TAG_NAME),$(TAG_NAME),$(SHA))
BUILD_DATE := $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')

BIN_OUTPUT := $(if $(filter $(shell go env GOOS), windows), mixtus.exe, mixtus)

default: clean check test build

test: clean
	go test -v -cover ./...

clean:
	rm -rf dist/ cover.out

build: clean
	@echo Version: $(VERSION) $(BUILD_DATE)
	CGO_ENABLED=0 go build -v -trimpath -ldflags '-s -w -X "main.version=${VERSION}" -X "main.commit=${SHA}" -X "main.date=${BUILD_DATE}"' -o ${BIN_OUTPUT} .

check:
	golangci-lint run

release-test:
	goreleaser --skip-publish --snapshot --rm-dist