MAKEFLAGS += --silent

.PHONY: *
SHELL=/bin/bash -o pipefail

COLOR="\033[32m%-25s\033[0m %s\n"

TITLE ?= lorem
BODY ?= body

.PHONY: help
help: ## Show this help
	@printf "\033[33m%s:\033[0m\n" 'Available commands'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "  "${COLOR}, $$1, $$2}' ${MAKEFILE_LIST}

.PHONY: run_server
run_server: ## Run server
	go run server/cmd/main.go -p=8080 -i="127.0.0.1"

client_send_get: ## Client send GET /posts/@TITLE, use TITLE=lorem
	go run client/cmd/main.go -u="127.0.0.1:8080" -p="posts/${TITLE}" -m="GET" 

client_send_post: ## Client send POST /posts
	go run client/cmd/main.go -u="127.0.0.1:8080" -p="posts" -m="POST" -b='{"title": "${TITLE}","body": "${BODY}"}'

# Global
.DEFAULT_GOAL := help
