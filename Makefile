BINARY_NAME=kevin-chromik-minihackathon-2024-challange-3
MODULE_PATH=github.com/matthiashermsen/${BINARY_NAME}

APP_VERSION=$(shell git tag --points-at HEAD)

COMMON_FLAGS=-ldflags="-X '$(MODULE_PATH)/app.Version=$(APP_VERSION)'"

BUILD_DIR=./build

analyze:
	@go vet ./...

tidy:
	@go mod tidy -v

test:
	@go test -cover ./...

coverage:
	@mkdir -p ./coverage
	@go test -coverprofile=./coverage/cover.out ./...
	@go tool cover -html=./coverage/cover.out -o ./coverage/cover.html

build:
	@go build $(COMMON_FLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)

build-all: build-darwin-amd64 build-darwin-arm64 build-linux-amd64 build-windows-amd64

build-darwin-amd64:
	@GOOS=darwin GOARCH=amd64 go build $(COMMON_FLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64

build-darwin-arm64:
	@GOOS=darwin GOARCH=arm64 go build $(COMMON_FLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64

build-linux-amd64:
	@GOOS=linux GOARCH=amd64 go build $(COMMON_FLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64

build-windows-amd64:
	@GOOS=windows GOARCH=amd64 go build $(COMMON_FLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe

.PHONY: analyze \
		tidy \
		test \
		coverage \
		build \
		build-all \
		build-darwin-amd64 \
		build-darwin-arm64 \
		build-linux-amd64 \
		build-windows-amd64
