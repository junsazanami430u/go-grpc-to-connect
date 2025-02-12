.DEFAULT_GOAL := help

## show help
.PHONY: help
help:
	@echo "usage: make <target>"
	@echo
	@echo "---- target list ----"
	@cat Makefile \
		| awk -F: '/^## /{desc=substr($$0,4)} /^[a-zA-Z_-]+:/&&desc{printf "% 10s :%s\n",$$1,desc; desc=""}'

## install dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@./scripts/install-deps.sh