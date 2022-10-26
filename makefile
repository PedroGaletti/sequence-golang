install:
	@echo "Downloading dependecies..."
	@go get
	@echo "Validating dependecies..."
	@go mod tidy
	@echo "Creating vendor..."
	@go mod vendor
	@echo "Installation completed successfully."

build:
	@echo "Building project..."
	@go build
	@echo "Build completed successfully."

coverage:
	@echo "Running project coverage..."
	@go test ./... -coverprofile fmtcoverage.html
	@go test ./... -coverprofile cover.out
	@go tool cover -html=cover.out
	@go tool cover -html=cover.out -o cover.html
	@echo "Coverage completed successfully."

run:
	@echo "Running application..."
	@docker-compose up -d
	@go run main.go

clean:
	@echo "Cleaning up project..."
	@rm -rf ./vendor
	@rm -rf ./go.sum
	@rm -rf ./challenge
	@echo "Project cleaned successfully."