-include .env

GO := go
GIT := git
VERSION := v`cat version`

MAKEFLAGS += --silent

generate:
	$(GO) generate

test:
	$(GO) test

release:
	$(GIT) add --all
	$(GIT) commit -S
	$(GIT) tag -s $(VERSION)

push:
	$(GIT) push --follow-tags

.PHONY: generate test release push