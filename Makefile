.DEFAULT_GOAL := help
.PHONY: help vet

PWD=$(shell pwd)
APP_NAME?=$(shell pwd | xargs basename)
APP_DIR=$(shell echo /${APP_NAME})
INTERACTIVE_OR_DETACH:=$(shell [ -t 0 ] && echo --interactive || echo --detach)
PROJECT_FILES=$(shell find . -type f -name '*.go' -not -path "./vendor/*")
HTTP_PORT:=80
GRPC_PORT:=6666
DEBUG_PORT:=2345

welcome:
	@printf "\033[33m    ______                 __     ____             __				\n"
	@printf "\033[33m   / ____/   _____  ____  / /_   / __ )_________  / /_____  _____	\n"
	@printf "\033[33m  / __/ | | / / _ \/ __ \/ __/  / __  / ___/ __ \/ //_/ _ \/ ___/	\n"
	@printf "\033[33m / /___ | |/ /  __/ / / / /_   / /_/ / /  / /_/ / ,< /  __/ /		\n"
	@printf "\033[33m/_____/ |___/\___/_/ /_/\__/  /_____/_/   \____/_/|_|\___/_/		\n"
	@printf "\n"

server-debug: welcome .env vendor ## Runs gRPC server in debug mode
	@echo 'Running on http://localhost:${GRPC_PORT}'

	@docker run \
		${INTERACTIVE_OR_DETACH} \
		--tty \
		--rm \
		--volume ${PWD}:${APP_DIR} \
		--expose ${GRPC_PORT} \
		--expose ${DEBUG_PORT} \
		--publish ${GRPC_PORT}:${HTTP_PORT} \
		--publish ${DEBUG_PORT}:${DEBUG_PORT} \
		--env DEBUG_PORT=${DEBUG_PORT} \
		--name ${APP_NAME} \
		${APP_NAME} \
		modd -f ./cmd/server/debug_modd.conf

server-release: welcome .env release ## Runs http server in release mode
	@echo 'Running on http://localhost:${GRPC_PORT}'

	@docker run \
		${INTERACTIVE_OR_DETACH} \
		--tty \
		--rm \
		--volume ${PWD}:${APP_DIR} \
		--expose ${HTTP_PORT} \
		--publish ${GRPC_PORT}:${HTTP_PORT} \
		--name ${APP_NAME} \
		${APP_NAME}

release: welcome .env
	@docker build \
		--target release \
		--tag ${APP_NAME} \
		.

test: welcome vendor
	@go clean --testcache

	@docker run \
		${INTERACTIVE_OR_DETACH} \
		--tty \
		--rm \
		--volume ${PWD}:${APP_DIR} \
		--name ${APP_NAME}-test \
		${APP_NAME} \
		go clean --testcache && \
		go test ./... -race

lint: welcome .env ## Code verifier
	@docker run \
		${INTERACTIVE_OR_DETACH} \
		--tty \
		--rm \
		--volume ${PWD}:${APP_DIR} \
		--name ${APP_NAME}-lint \
		${APP_NAME} \
		golangci-lint run --print-resources-usage --timeout=180s --out-format=tab ./...

.env:
	@cp .env.default .env

vendor:
	@go mod vendor

setup: welcome .env vendor ## Install dependencies

clean: ## Cleans vendor and temp files
	@-rm -rf vendor* _vendor* coverage.xml

format: ## Runs automatic and built-in code formatter
	@command -v goimports >/dev/null 2>&1 || go get -u golang.org/x/tools/cmd/goimports
	@goimports -l -w -d ${PROJECT_FILES}
	@gofmt -l -s -w ${PROJECT_FILES}

vet: ## Reports suspicious constructs
	@go vet ./...

help: welcome
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | grep ^help -v | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-22s\033[0m %s\n", $$1, $$2}'
