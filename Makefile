build:
	@go build -ldflags "-s -w -X main.version=v1.2.3" -o bin/example ./cmd/example
	@./bin/example --version

build-no-ldflags:
	@go build -o bin/example ./cmd/example
	@./bin/example --version
