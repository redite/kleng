deps:
	@echo "Installing dependencies"
	@glide install

build:
	@echo "Building Kleng"
	@go build -x -v -race -o kleng

run:
	@echo "Running kleng with HTTP mode"
	./kleng -M http

populate:
	@echo "Running kleng with Populate mode"
	./kleng -M populate

clean:
	@rm -f kleng
	@go clean all

.PHONY: deps build populate clean run
