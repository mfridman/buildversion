build:
	@go build -ldflags "-X github.com/mfridman/buildversion.Version=v1.2.3" -o bin/example ./cmd/example
	@./bin/example --version

build-no-ldflags:
	@go build -o bin/example ./cmd/example
	@./bin/example --version
