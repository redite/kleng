build:
	@echo "Building Kleng"
	@go build -x -v -race -o kleng

run:
	@echo "Running kleng with HTTP mode"
	@make build
	./kleng -M http

populate:
	@echo "Running kleng with Populate mode"
	@make build
	./kleng -M populate

clean:
	@rm -f kleng
	@go clean all

.PHONY: deps build
