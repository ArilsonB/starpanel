BINARY_NAME=starpanel

MAIN_PATH=./api/main.go

build:
	@echo "Building..."
	@go build -o $(BINARY_NAME) $(MAIN_PATH)
	@echo "Build complete."

run:
	@echo "Running..."
	@go run $(MAIN_PATH)
	@echo "Run complete."

clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@echo "Clean complete."

lint:
	@echo "Linting..."
	@golangci-lint run ./...
	@echo "Linting complete."

test:
	@echo "Running tests..."
	@go test ./...
	@echo "Tests complete."

start: build
	@echo "Starting..."
	@./$(BINARY_NAME)
	@echo "Start complete."