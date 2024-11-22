.PHONY: lint lint-fix test down build up clean wire swagger generate-mocks fetch-graphql-schema clean-graphql-schema generate-graphql help

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

swagger: ## Generate swagger
	swag init -g cmd/api/main.go -o docs/swagger
	npm run convert-openapi

generate-mock:
	go generate ./internal/mock/...

fetch-graphql-schema:
	GRAPHQL_ENDPOINT=$(GRAPHQL_ENDPOINT) npm run fetch-schema

clean-graphql-schema:
	npm run fetch-schema:clean

generate-graphql:
	go run github.com/Khan/genqlient


help: ## display this help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
