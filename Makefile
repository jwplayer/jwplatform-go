APPNAME = jwplatform-go
VERSION := $(shell head CHANGELOG.md | grep -e '^[0-9]' | head -n 1 | cut -f 1 -d ' ')
GOPATH  ?= $(curdir)/.gopath

export GOPATH

all: test

test:
	@echo -e '\e[01;34mRunning Go unit tests\e[0m'
	@go test -cover

release: version.go | test
	go mod tidy
	@echo "Releasing $(APPNAME) v$(VERSION)"
	git tag v$(VERSION)
	git push --tags

clean:
	@rm -rf build

distclean: clean
	@rm -rf Gopkg.lock

.PHONY: clean distclean
