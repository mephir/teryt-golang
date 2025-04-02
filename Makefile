default: build

build-downloader:
	@echo "Building downloader..."
	go build -o build/teryt-downloader cmd/downloader/main.go

build-exporter:
	@echo "Building exporter..."
	go build -o build/teryt-exporter cmd/exporter/main.go

build: build-downloader build-exporter

test:
	@echo "Running tests..."
	go test -v .\internal\dataset\

clean:
	@echo "Cleaning..."
	rm -rf build
