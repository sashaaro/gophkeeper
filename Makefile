LOCAL_BIN:=$(CURDIR)/bin
DC = docker compose

help: ## Show help message
	@cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test

test:
	go test ./...

coverage: ## Run tests
	go test -race -cover -coverprofile=coverage.out ./...
	sed -i "/\(\/cmd\/\|\/internal\/application\/\|\/internal\/config\/\)/d" coverage.out
	go tool cover -func=coverage.out

protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./internal/contract/gophkeeper.proto

lint:
	golangci-lint run

installDevTools:
	go install -mod mod github.com/golang/mock/mockgen@v1.6.0
