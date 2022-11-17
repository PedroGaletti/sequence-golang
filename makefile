install:
	@echo "Downloading dependecies..."
	@go get
	@echo "Validating dependecies..."
	@go mod tidy
	@echo "Installation completed successfully."

build:
	@echo "Building project..."
	@go build
	@echo "Build completed successfully."

coverage:
	@echo "Running project coverage..."
	@go test ./... -coverprofile=cover.out
	@go tool cover -func=cover.out
	@echo "Coverage completed successfully."

run:
	@echo "Running application..."
	@docker-compose up -d
	@go run main.go

clean:
	@echo "Cleaning up project..."
	@rm -rf ./go.sum
	@rm -rf ./challenge
	@echo "Project cleaned successfully."