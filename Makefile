# Run all unit tests in the project.
.PHONY: test
test:
	@echo 🧪 Run unit tests...
	@go test ./... -race -count=1
