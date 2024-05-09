start: fe-build build run

build:
	@echo "Building..."
	@cd src/backend && go build -o ../../main cmd/api/main.go

fe-build:
	@echo "Building..."
	@cd src/frontend && npm install && npm run build

# Run the application
run:
	@export $$(cat .env | xargs) && cd src/backend && go run cmd/api/main.go

# Run the application
fe-run:
	@cd src/frontend && npm run dev


# Create DB container
docker-run:
	@if docker compose up 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@cd src/backend && go test ./tests -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main
	@cd src/backend && gofmt -w .
