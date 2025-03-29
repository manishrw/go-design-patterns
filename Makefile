.PHONY: build test test-v clean deps test-coverage test-race test-race-v test-focus test-until-fail

# Build the application
build:
	go build ./...

# Define common Ginkgo test command options
GINKGO_CMD = go test -v -ginkgo.v

# Run all Ginkgo tests
test:
	$(GINKGO_CMD) ./...

# Clean build artifacts
clean:
	go clean
	rm -f *.out

# Install dependencies
deps:
	go mod tidy
	go install github.com/onsi/ginkgo/v2/ginkgo@latest

# Run Ginkgo tests with coverage
test-coverage:
	$(GINKGO_CMD) --cover ./...
	go tool cover -html=coverage.out -o coverage.html

# Run Ginkgo tests with race detector
test-race:
	$(GINKGO_CMD) -race ./...

# Run Ginkgo tests with race detector and verbose output
test-race-v:
	$(GINKGO_CMD) -race ./... 