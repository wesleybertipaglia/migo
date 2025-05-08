BINARY_NAME=migo
BUILD_DIR=bin

build:
	@echo "Building $(BINARY_NAME)..."
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go
	@echo "Build complete. Executable is located at $(BUILD_DIR)/$(BINARY_NAME)"

install: build
	@echo "Installing $(BINARY_NAME)..."
	sudo cp -r $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/
	@echo "$(BINARY_NAME) installed successfully."

clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)
	@echo "Clean complete."

dev:
	@echo "Running in development mode..."
	go run main.go	
