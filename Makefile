-include .env

GO := go
GIT := git
VERSION := v`cat version`

MAKEFLAGS += --silent

generate:
	$(GO) generate

test:
	$(GO) test -v

release:
	$(GIT) add --all
	$(GIT) commit -S

full-release: release
	$(GIT) tag -s $(VERSION)

push:
	$(GIT) push --follow-tags

.PHONY: generate test release full-release push