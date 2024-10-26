LOCAL_BIN:=$(CURDIR)/bin
DC = docker compose

help: ## Show help message
	@cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

clean: ## Remove temporary files
	rm -rf ./build ./coverage.out

test: ## Run tests
	go test ./...

coverage: ## Run tests with coverage report
	go test -race -cover -coverprofile=coverage.out `go list ./... | grep -v "/cmd/" | grep -v "/tests/mock/"`
	sed -i '/\(\/cmd\/\|\/mock\/\|\.pb\.go\)/d' coverage.out
	go tool cover -func=coverage.out

protoc: ## Generate package from proto files
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pkg/gophkeeper/gophkeeper.proto

lint: ## Static check
	golangci-lint run

installDevTools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1
	go install -mod mod github.com/golang/mock/mockgen@v1.6.0
