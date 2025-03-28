default: build

build-downloader:
	@echo "Building downloader..."
	go build -o bin/downloader cmd/downloader/main.go

build-search:
	@echo "Building search..."

build: build-downloader build-search

test:
	@echo "Running tests..."
	go test -v .\internal\dataset\

clean:
	@echo "Cleaning..."
	rm -rf bin
