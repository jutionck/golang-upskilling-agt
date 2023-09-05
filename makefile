# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOTOOL=$(GOCMD) tool
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOINST=$(GOCMD) install

# Binary name
BINARY_NAME=app

# Build
build:
	@$(GOBUILD) -o $(BINARY_NAME) .
	@echo "📦 Build Done"

# Clean
clean:
	@$(GOCLEAN)
	@rm -f $(BINARY_NAME)
	@rm -f test.out
	@echo "🧹 Program removed"

# Generate the doc
doc:
	@$(GOINST) github.com/swaggo/swag/cmd/swag@latest 
	@swag init --parseDependency=true -g app.go >> output.out
	@rm output.out
	@echo "📓 Docs Generated"

# Run apps from development
dev:
	@$(GOCMD) run .

# Run test without coverage
test:
	@echo "🚀 Running App Test"
	@$(GOTEST) -v ./...

# Run test with coverage
coverage:
	@$(GOTEST) -coverprofile=coverage.out ./...
	@$(GOTOOL) cover -html=coverage.out
	@rm coverage.out
	@echo "🎯 Cover profile generated"

# Build and run
run: doc build
	@echo "🚀 Running App"
	@./$(BINARY_NAME)