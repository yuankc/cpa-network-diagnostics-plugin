PLUGIN_NAME := diagnostics
VERSION ?= 0.1.0
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
EXT := so
ifeq ($(GOOS),windows)
	EXT := dll
endif
ifeq ($(GOOS),darwin)
	EXT := dylib
endif

DIST_DIR := dist/$(GOOS)/$(GOARCH)
PLUGIN_FILE := $(DIST_DIR)/$(PLUGIN_NAME).$(EXT)
RELEASE_DIR := release
ZIP_FILE := $(RELEASE_DIR)/$(PLUGIN_NAME)_$(VERSION)_$(GOOS)_$(GOARCH).zip

.PHONY: all fmt vet test build package checksums clean

all: fmt vet test build

fmt:
	gofmt -w main.go

vet:
	go vet ./...

test:
	go test ./...

build:
	mkdir -p $(DIST_DIR)
	CGO_ENABLED=1 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -buildmode=c-shared -trimpath -ldflags "-s -w" -o $(PLUGIN_FILE) .

package: build
	mkdir -p $(RELEASE_DIR)
	rm -f $(ZIP_FILE)
	cd $(DIST_DIR) && zip -q ../../../$(ZIP_FILE) $(PLUGIN_NAME).$(EXT)

checksums: package
	cd $(RELEASE_DIR) && sha256sum $$(basename $(ZIP_FILE)) > $$(basename $(ZIP_FILE)).sha256
	cd $(RELEASE_DIR) && sha256sum *.zip > checksums.txt

clean:
	rm -rf dist release
