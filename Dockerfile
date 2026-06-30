# syntax=docker/dockerfile:1.7

FROM golang:1.22-alpine AS build

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .

ARG VERSION=dev
ARG COMMIT=none
ARG BUILT_AT=unknown

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-s -w -X github.com/happysnaker/go-service-starter/internal/buildinfo.Version=${VERSION} -X github.com/happysnaker/go-service-starter/internal/buildinfo.Commit=${COMMIT} -X github.com/happysnaker/go-service-starter/internal/buildinfo.BuiltAt=${BUILT_AT}" \
    -o /out/api ./cmd/api

FROM alpine:3.20

RUN addgroup -S app && adduser -S -G app -u 10001 app

WORKDIR /app

COPY --from=build /out/api /app/api

EXPOSE 8080

USER app

ENTRYPOINT ["/app/api"]
