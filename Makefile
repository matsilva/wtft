.phony: build-cli
build-cli:
	@echo "Building CLI..."
	@go build -o wtft ./main.go
	@echo "Done! run ./wtft to start the CLI"
