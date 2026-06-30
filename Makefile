.RECIPEPREFIX := >

APP_NAME ?= go-service-starter
ENV_FILE ?= configs/service.env.example

.PHONY: help run fmt tidy docker-build docker-run

help:
> @printf "%s\n" \
> "Available targets:" \
> "  make run          # start the service with $(ENV_FILE)" \
> "  make fmt          # format Go source files" \
> "  make tidy         # tidy go.mod / go.sum" \
> "  make docker-build # build a local container image" \
> "  make docker-run   # run the local container image"

run:
> set -a; [ -f $(ENV_FILE) ] && . $(ENV_FILE); set +a; go run ./cmd/api

fmt:
> gofmt -w ./cmd ./internal

tidy:
> go mod tidy

docker-build:
> docker build -t $(APP_NAME):dev .

docker-run:
> docker run --rm -p 8080:8080 --env-file $(ENV_FILE) $(APP_NAME):dev
