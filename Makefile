## Install library for production
deps:
	@go mod tidy
.PHONY: deps

fmt:
	@go fmt ./...

## Execute unit tests
test:
	@go test -v -count=1 -timeout 300s -short ./...
.PHONY: test

## Execute race tests
test-race:
	@go test -v -count=1 -timeout 300s -short -race ./...
.PHONY: test-race

## Execute integrated tests
test-integration:
	@go test -v -count=1 -timeout 600s ./...
.PHONY: test-integration

## Output coverage of testing
cov:
	@go test -count 1 -coverprofile=cover.out ./...
	@go tool cover -html=cover.out
.PHONY: cov

## Lint
lint:
	golangci-lint run --tests
.PHONY: lint

# Execute this command before you creates a pull request
pr-prep: fmt lint test-race test-integration

## Help
help:
	@make2help --all
.PHONY: help
