.PHONY: lint lint-fix test test-ginkgo go-download go-build docker-down docker-build docker-up wire fetch-graphql-schema clean-graphql-schema generate-graphql swagger generate-mocks  help

# 環境変数
NAME := go-rest-chi
REVISION   := $(shell git rev-parse --short HEAD)
DC := docker compose
DOCKER_BUILDKIT := 1
export DOCKER_BUILDKIT

# ビルド設定
LDFLAGS := -ldflags='-s -w -X "main.Revision=$(REVISION)" -extldflags "-static"'
CGO_ENABLED ?= 0
GOOS ?= linux
GOARCH ?= amd64

#include .envrc
#export

## CI #########################################################################################
lint: ## Run linter
	golangci-lint run -v

lint-fix: ## Run linter with fix
	golangci-lint run -v --fix

test: ## Run tests
	go test -v ./...

test-ginkgo: ## Run tests
	ginkgo -v -p --randomize-all ./...

test-ginkgo-coverage: ## Run tests with coverage
	ginkgo --race --cover --covermode=atomic \
		--output-dir=./coverage \
		--coverprofile=coverage.txt \
		--randomize-all \
		-p -v ./...

go-download: ## Download dependencies
	go mod download
	go mod verify
	go mod tidy

go-build: ## Build binary
	CGO_ENABLED=$(CGO_ENABLED) \
	GOOS=$(GOOS) \
	GOARCH=$(GOARCH) \
	go build -tags netgo -installsuffix netgo $(LDFLAGS) \
	-o bin/$(NAME) ./cmd/api


## Container ##################################################################################
docker-down: ## Stop container
	$(DC) down

docker-build: ## Build image
	$(DC) build $(NAME)

docker-up: ## Run container
	$(DC) up -d


## Generate ###################################################################################
wire: ## Generate wire
	wire ./cmd/api

fetch-graphql-schema: ## Fetch graphql schema
	GRAPHQL_ENDPOINT=$(GRAPHQL_ENDPOINT) npm run fetch-schema

clean-graphql-schema: ## Clean graphql schema
	npm run fetch-schema:clean

generate-graphql: ## Generate graphql
	go run github.com/Khan/genqlient

swagger: ## Generate swagger
	go run github.com/swaggo/swag/cmd/swag init -g cmd/api/main.go -o docs/openapi
	npm run convert-openapi

generate-mock: ## Generate mock
	go generate ./internal/mock/...


help: ## display this help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
