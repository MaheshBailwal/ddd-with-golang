help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"; printf "\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help

run: linter-golangci swagger ### run
	go mod tidy && go mod download && \
	go run .
.PHONY: run

swagger: ### init swagger
	swag fmt
	swag init --parseDependency --parseInternal --parseDepth 1 --generalInfo ./presentation/api/v1/v1.go --output ./presentation/api/docs --instanceName v1 --generatedTime true
.PHONY: swagger

linter-golangci: ### check by golangci linter
	golangci-lint run ./...
.PHONY: linter-golangci

test: ### run unit-test and integration-test
	go test -v -cover -race ./... && \
    go clean -testcache && \
    go test -v ./test/...
.PHONY: test

install-deps: ### install all dependencies
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/golang/mock/mockgen@latest
	go install golang.org/x/tools/cmd/stringer@latest
.PHONY: install-deps


proto: ### init proto
	protoc --proto_path=application/grpc/proto --go_out=application/grpc/pb --go_opt=paths=source_relative --go-grpc_out=application/grpc/pb --go-grpc_opt=paths=source_relative  application/grpc/proto/*.proto
.PHONY: proto

docker 
docker build -t department-service .
docker run -d -p 8003:8003 -p 8004:8004 --name departmet-container department-service