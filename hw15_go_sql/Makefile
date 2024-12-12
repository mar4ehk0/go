MAKEFLAGS += --silent

.PHONY: *
SHELL=/bin/bash -o pipefail

COLOR="\033[32m%-25s\033[0m %s\n"

ifneq ($(wildcard .env),)
	include .env
endif

PROJECT_DIR = $(shell pwd)
PROJECT_BIN = $(PROJECT_DIR)/bin
PROJECT_TMP = $(PROJECT_DIR)/tmp

.PHONY: help
help: ## Show this help
	@printf "\033[33m%s:\033[0m\n" 'Available commands'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "  "${COLOR}, $$1, $$2}' ${MAKEFILE_LIST}

install-deps: ## Setup install deps
	GOBIN=$(PROJECT_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest

migration-status: ## Migration status
	$(PROJECT_BIN)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} status -v

migration-add: ## Create migration
	$(PROJECT_BIN)/goose -dir ${MIGRATION_DIR} create $(name) sql

migration-up: ## Migration up
	$(PROJECT_BIN)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} up -v

migration-down: ## Migration down
	$(PROJECT_BIN)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} down -v




# TITLE ?= lorem
# BODY ?= body

# run_server: ## Run server
# 	go run server/cmd/main.go -p=8080 -i="127.0.0.1"

# client_send_get: ## Client send GET /posts/@TITLE, use TITLE=lorem
# 	go run client/cmd/main.go -u="127.0.0.1:8080" -p="posts/${TITLE}" -m="GET" 

# client_send_post: ## Client send POST /posts, use TITLE=lorem BODY=body
# 	go run client/cmd/main.go -u="127.0.0.1:8080" -p="posts" -m="POST" -b='{"title": "${TITLE}","body": "${BODY}"}'

# lint: ## Run linter
# 	golangci-lint run -c ../.golangci.yml


# Global
.DEFAULT_GOAL := help
