APP_NAME := server
GO := go
DOCKER := docker
DOCKER_COMPOSE := docker compose
ENV_FILE := .env

# Load .env if present
ifneq (,$(wildcard $(ENV_FILE)))
	include $(ENV_FILE)
	export
else
	$(warning $(ENV_FILE) not found)
endif

.PHONY: all build run test clean fmt vet \
        docker-build docker-up docker-down \
        docker-dev docker-dev-down \
        migrate-up migrate-down

all: build

## --------------------
## Go targets
## --------------------

build:
	$(GO) build -o bin/$(APP_NAME) ./cmd/server/main.go

run:
	APP_ENV=dev $(GO) run ./cmd/server/main.go

test:
	@echo "Running tests with in-memory SQLite..."
	APP_ENV=test $(GO) test -v ./...

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

clean:
	rm -rf bin

swagger:
	swag init -g cmd/server/main.go

## --------------------
## Docker (Production)
## --------------------

docker-build:
	$(DOCKER) build -t $(APP_NAME):latest .

docker-up:
	$(DOCKER_COMPOSE) --env-file $(ENV_FILE) up --build

docker-down:
	$(DOCKER_COMPOSE) down -v

## --------------------
## Docker (Development)
## --------------------

docker-dev:
	$(DOCKER_COMPOSE) -f docker-compose.dev.yml up --build

docker-dev-down:
	$(DOCKER_COMPOSE) -f docker-compose.dev.yml down

## --------------------
## Database migrations
## --------------------

migrate-up:
	@echo "GORM AutoMigrate runs automatically on application startup"

migrate-down:
	@echo "Down migrations are not supported by GORM AutoMigrate"
