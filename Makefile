.DEFAULT_GOAL = check

.PHONY: test
test:
	@echo 🧪 Run unit tests...
	@go test ./... -race -count=1

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: check
check: test lint
