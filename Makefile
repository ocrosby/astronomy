.PHONY: install lint test clean

clean:
	@echo "Cleaning up..."
	rm -f junit.xml

install:
	@echo "Installing dependencies..."
	go mod tidy
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/jstemmer/go-junit-report@latest

lint:
	@echo "Running golangci-lint..."
	golangci-lint run

test:
	@echo "Running unit tests..."
	go test -v ./... | go-junit-report > junit.xml

