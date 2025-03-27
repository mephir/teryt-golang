default: build

build-download:
	@echo "Building download..."

build-search:
	@ehco "Building search..."

build:
	@echo "Building..."
	build-download
	build-search

test:
	@echo "Running tests..."