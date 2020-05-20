-include .env

GO := go
GIT := git

MAKEFLAGS += --silent

generate:
	$(GO) generate

test:
	$(GO) test

release:
	$(GIT) add --all
	$(GIT) commit -S
	$(GIT) tag -s

push: clean
	$(GIT) push --follow-tags

.PHONY: generate test release push