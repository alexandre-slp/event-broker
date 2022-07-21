.DEFAULT_GOAL := help
.PHONY: help vet vendor proto

PWD=$(shell pwd)
APP_NAME?=$(shell pwd | xargs basename)
APP_DIR=$(shell echo /${APP_NAME})
INTERACTIVE_OR_DETACH:=$(shell [ -t 0 ] && echo --interactive || echo --detach)
PROJECT_FILES=$(shell find . -type f -name '*.go' -not -path "./vendor/*")
HAS_DEBUG_IMAGE:=$(shell docker images --quiet ${APP_NAME}-debug)
HAS_DEV_IMAGE:=$(shell docker images --quiet ${APP_NAME}-dev)
PROTO_FILES:=$(shell ls infra/api/v1 | grep proto)
PROTO_PATH:=infra/api/v1
PROTO_MODULE:=${APP_NAME}/${PROTO_PATH}
PROTO_SERIALIZER:=${PROTO_PATH}/serializer
PROTO_GRPC:=${PROTO_PATH}/gRPC
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

proto: ## Generate gRPC and Serializer from Proto files
	@protoc --proto_path=${PROTO_PATH} \
		    --go_out=${PROTO_SERIALIZER} \
		    --go_opt=module=${PROTO_MODULE} \
		    --go-grpc_out=${PROTO_GRPC} \
		    --go-grpc_opt=module=${PROTO_MODULE} \
		    ${PROTO_FILES}

debug: welcome .env vendor build-debug ## Run gRPC server in debug mode
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
		--name ${APP_NAME}-debug \
		${APP_NAME}-debug \
		modd -f ./cmd/server/debug_modd.conf

dev: welcome .env vendor build-dev ## Run gRPC server in dev mode
	@echo 'Running on http://localhost:${GRPC_PORT}'
	@docker run \
		${INTERACTIVE_OR_DETACH} \
		--tty \
		--rm \
		--volume ${PWD}:${APP_DIR} \
		--expose ${GRPC_PORT} \
		--publish ${GRPC_PORT}:${HTTP_PORT} \
		--name ${APP_NAME}-dev \
		${APP_NAME}-dev \
		modd -f ./cmd/server/dev_modd.conf

build-debug: welcome .env
	@if [ -z ${HAS_DEBUG_IMAGE} ]; then \
  		docker build \
  		--target debug \
  		--tag ${APP_NAME}-debug \
  		. ; \
  	fi

build-dev: welcome .env
	@if [ -z ${HAS_DEV_IMAGE} ]; then \
  		docker build \
  		--target dev \
  		--tag ${APP_NAME}-dev \
  		. ; \
  	fi

test: welcome vendor ## Run tests
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

lint: welcome .env ## Run linter
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

vendor: ## Run vendor
	@go mod vendor

setup: welcome .env vendor ## Install dependencies

clean: ## Clean vendor and temp files
	@-rm -rf vendor* _vendor* coverage.xml

format: ## Run code formatter
	@command -v goimports > /dev/null 2>&1 || go install golang.org/x/tools/cmd/goimports@latest
	@goimports -l -w -d ${PROJECT_FILES}
	@gofmt -l -s -w ${PROJECT_FILES}

vet: ## Report suspicious constructs
	@go vet ./...

help: welcome
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | grep ^help -v | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-22s\033[0m %s\n", $$1, $$2}'
