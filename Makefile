APP_NAME := ride-share
MAIN_FILE := cmd/server/main.go
BINARY_PATH := bin/$(APP_NAME)
MIGRATE_DIR := cmd/migrate/migrations
MIGRATE_MAIN := cmd/migrate/main.go

build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(BINARY_PATH) $(MAIN_FILE)
	@chmod +x $(BINARY_PATH)


run: build
	@echo "Running $(APP_NAME)..."
	@./$(BINARY_PATH)


test:
	@echo "Running tests..."
	@go test -v ./...


migration:
	ifndef name
		$(error Missing migration name. User: make migration name=your_migration_name)
	endif
		@echo "Creating migration $(name)..."
		@migrate create -ext sql -dir $(MIGRATE_DIR) -seq $(name)


migrate-up:
	@echo "Applying migrations..."
	@go run $(MIGRATE_MAIN) up



migrate-down:
	@echo "Rolling back all migrations..."
	@go run $(MIGRATE_MAIN) down


clean:
	@echo "Cleaning build files..."
	@rm -rf $(BINARY_PATH)


help:
	@echo ""
	@echo "🚖 Ride Share Backend Makefile"
	@echo "------------------------------"
	@echo "make build          # Build the Go binary"
	@echo "make run            # Build and run the app"
	@echo "make test           # Run all Go tests"
	@echo "make migration name=create_users_table   # Create a new migration"
	@echo "make migrate-up     # Apply DB migrations"
	@echo "make migrate-down   # Rollback last migration"
	@echo "make clean          # Delete built binaries"
	@echo ""


.PHONY: build run test migration migrate-up migrate-down clean help