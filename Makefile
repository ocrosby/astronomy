.PHONY: install lint test clean

clean:
	@echo "Cleaning up..."
	rm -f junit.xml

deps:
	@echo "Installing dependencies..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install gotest.tools/gotestsum@latest
	go install github.com/jstemmer/go-junit-report@latest

install: deps
	@echo "Installing ..."
	go mod tidy
	go mod download

lint:
	@echo "Running golangci-lint ..."
	golangci-lint run

test:
	@echo "Running unit tests ..."
	go test -v ./... | go-junit-report > junit.xml

