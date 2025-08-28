V="0.0.1"

.PHONY: all fa fmt lint test vet

all: fa fmt lint test vet

fa:
	@fieldalignment -fix ./...

fmt:
	@goimports -w -local github.com/pixel365/goreydenx .
	@gofmt -w .
	@golines -w .

lint:
	@golangci-lint run

test:
	@go test ./...

vet:
	@go vet ./...
