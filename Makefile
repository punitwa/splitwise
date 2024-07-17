BINARY_NAME=myservice

# Define the main package
MAIN_PACKAGE=./cmd/myservice

# Define the build output directory
BUILD_DIR=bin

.PHONY: all build clean

all: build

# Build the Go binary
build:
	@echo "Building the binary..."
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)

# Run the Go service
run: build
	@echo "Running the server..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Clean up the build output
clean:
	@echo "Cleaning up..."
	go clean
	rm -rf $(BUILD_DIR)