SHELL := /bin/bash
BASEDIR = $(shell pwd)

.PHONY: help
help:
		@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: install
install: ## Install dependencies
		@go mod download
		@go mod vendor

.PHONY: dev
dev: ## Run with Dev
		@test -f conf/config.local.yml || cp conf/config.local.yml.example conf/config.local.yml
		@go run cmd/todomvc/todomvc.go

.PHONY: build
build: ## Build todomvc
		@go build -o build/todomvc cmd/todomvc/todomvc.go

clean: ### Remove build dir
		@rm -fr build