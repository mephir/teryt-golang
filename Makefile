default: build

build-downloader:
	@echo "Building downloader..."
	go build -o bin/downloader cmd/downloader/main.go

build-conv:
	@echo "Building converter..."
	go build -o bin/conv cmd/conv/main.go

build: build-downloader build-conv

test:
	@echo "Running tests..."
	go test -v .\internal\dataset\

clean:
	@echo "Cleaning..."
	rm -rf bin
