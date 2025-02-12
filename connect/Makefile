.DEFAULT_GOAL := help

## show help
.PHONY: help
help:
	@echo "usage: make <target>"
	@echo
	@echo "---- target list ----"
	@cat Makefile \
		| awk -F: '/^## /{desc=substr($$0,4)} /^[a-zA-Z_-]+:/&&desc{printf "% 10s :%s\n",$$1,desc; desc=""}'

## go test
.PHONY: test
test:
	@../scripts/test.sh

## staticcheck
.PHONY: lint
lint:
	golangci-lint run

## go mod tidy
.PHONY: tidy
tidy:
	go mod tidy

## server-serve
.PHONY: serve
serve:
	go run cmd/server/main.go

## client
.PHONY: client
client:
	go run cmd/client/main.go