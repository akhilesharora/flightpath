.PHONY: test* run build docker* benchmarks

PACKAGE_NAME := github.com/akhilesharora/flightpath

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: clean ## Compile server
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/flightpath cmd/main.go

run: build ## Run server
	go run cmd/main.go

test: ## Run all tests
	go test ./... -v count=1

test-race:
	go test -race -short ./...


test-clean-testcache:
	go clean -testcache && go test -v ./...

benchmarks:
	go test -bench=. ./...

coverage: clean ## Run tests and generate coverage files per package
	mkdir .coverage 2> /dev/null || true
	rm -rf .coverage/*.out || true
	go test -race ./... -coverprofile=coverage.out -covermode=atomic

clean:
	rm -rf .coverage/ build/

docker: ## Compile docker image
	docker build -t akhilesharora/flightpath .

docker-run: docker ## Run docker container
	docker run -p 8080:8080 --rm akhilesharora/flightpath
