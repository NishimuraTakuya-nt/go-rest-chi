.PHONY: lint lint-fix test test-ginkgo down build up clean wire fetch-graphql-schema clean-graphql-schema generate-graphql swagger generate-mocks  help

NAME := go-rest-chi
DC := docker compose
LDFLAGS := -ldflags="-s -w -extldflags \"-static\""

include .envrc
export

## CI #########################################################################################
lint: ## Run linter
	golangci-lint run

lint-fix: ## Run linter with fix
	golangci-lint run --fix

test: ## Run tests
	go test -v ./...

test-ginkgo: ## Run tests
	ginkgo -v -p ./...


## Container ##################################################################################
down: ## Stop container
	$(DC) down

build: ## Build image
	$(DC) build $(NAME)

up: ## Run container
	$(DC) up -d

clean: ## Clean up
	docker system prune -f


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
	swag init -g cmd/api/main.go -o docs/swagger
	npm run convert-openapi

generate-mock: ## Generate mock
	go generate ./internal/mock/...


help: ## display this help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
