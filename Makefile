#!/usr/bin/make

PKG := .
BIN := $(shell basename `pwd`)
GO  := $(realpath ./go)

DEPS := $(shell $(GO) list -f '{{join .Deps "\n"}}' $(PKG) \
	| sort | uniq | grep -v "^_")

.PHONY: %

default: test

all: build
build: deps
	$(GO) build $(PKG)
lint: vet
vet: deps
	$(GO) get code.google.com/p/go.tools/cmd/vet
	$(GO) vet $(PKG)
fmt:
	$(GO) fmt $(PKG)
test: test-deps
	$(GO) test $(PKG)
cover: test-deps
	$(GO) test -cover $(PKG)
clean:
	$(GO) clean -i $(PKG)
clean-all:
	$(GO) clean -i -r $(PKG)
deps:
	$(GO) get -d $(PKG)
	$(GO) install $(DEPS)
test-deps: deps
	$(GO) get -d -t $(PKG)
	$(GO) test -i $(PKG)
run: all
	./$(BIN)

