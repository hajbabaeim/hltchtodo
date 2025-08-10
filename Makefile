.PHONY: build run stop clean logs test docker-build

# Build the application locally
build:
	go build -o bin/todo-app main.go

# Run with docker-compose
run:
	docker-compose up --build -d
	@echo "✅ Application starting..."
	@echo "🔗 API: http://localhost:8080"
	@echo "🔗 Health: http://localhost:8080/health"
	@echo "📊 Use 'make logs' to view logs"

# Stop all services
stop:
	docker-compose down

# View application logs
logs:
	docker-compose logs -f todo-app

# Check service status
status:
	docker-compose ps

# Clean up everything
clean:
	docker-compose down -v --rmi all --remove-orphans
	docker system prune -f

# Run tests
test:
	go test -v ./...

# Build Docker image only
docker-build:
	docker build -t todo-app .

# Quick API test
test-api:
	@echo "🧪 Testing API endpoints..."
	@echo "1. Health check:"
	@curl -s http://localhost:8080/health | jq .
	@echo "\n2. Create todo:"
	@curl -s -X POST http://localhost:8080/api/v1/todos \
		-H "Content-Type: application/json" \
		-d '{"title":"Test Todo","description":"Test Description","dueDate":"2024-12-31T23:59:59Z"}' | jq .
	@echo "\n3. List todos:"
	@curl -s http://localhost:8080/api/v1/todos | jq .

help:
	@echo "Available commands:"
	@echo "  run        - Start all services with docker-compose"
	@echo "  stop       - Stop all services"
	@echo "  logs       - View application logs"
	@echo "  status     - Check service status"
	@echo "  clean      - Clean up all containers and volumes"
	@echo "  test       - Run Go tests"
	@echo "  test-api   - Test API endpoints (requires jq)"
	@echo "  build      - Build application locally"
	@echo "  help       - Show this help"