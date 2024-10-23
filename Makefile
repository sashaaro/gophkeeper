LOCAL_BIN:=$(CURDIR)/bin
DC = docker compose

help: ## Show help message
	@cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: dependencies

test:
	go test ./...

coverage: ## Run tests
	go test -race -cover -coverprofile=coverage.out ./...
	sed -i "/\(\/cmd\/\|\/internal\/application\/\|\/internal\/config\/\)/d" coverage.out
	go tool cover -func=coverage.out