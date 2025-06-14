
# Go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean

# Docker commands
DOCKER_COMPOSE=docker compose

# Build the application
build:
	$(GOBUILD) -o $(APP_NAME) ./cmd/app

# Run the application
run:
	$(GORUN) ./cmd/app/main.go

