.PHONY: all build run dev test clean docs migrate lint

# 变量定义
APP_NAME=irt-exam-system
MAIN_FILE=cmd/server/main.go
BUILD_DIR=build
DOCKER_IMAGE=$(APP_NAME)
DOCKER_TAG=latest

# 默认目标
all: build

# 构建项目
build:
	@echo "Building $(APP_NAME)..."
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

# 运行项目
run:
	@echo "Running $(APP_NAME)..."
	go run $(MAIN_FILE)

# 开发模式（热重载）
dev:
	@echo "Starting development server..."
	air -c .air.toml

# 运行测试
test:
	@echo "Running tests..."
	go test -v ./...

# 运行测试覆盖率
test-coverage:
	@echo "Running tests with coverage..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# 清理构建文件
clean:
	@echo "Cleaning..."
	rm -rf $(BUILD_DIR)
	go clean

# 生成文档
docs:
	@echo "Generating Swagger documentation..."
	swag init -g $(MAIN_FILE)

# 数据库迁移
migrate:
	@echo "Running database migrations..."
	go run scripts/migrate.go up

# 回滚数据库
migrate-down:
	@echo "Rolling back database migrations..."
	go run scripts/migrate.go down

# 代码格式化
fmt:
	@echo "Formatting code..."
	go fmt ./...

# 代码检查
lint:
	@echo "Linting code..."
	golangci-lint run

# 依赖更新
deps:
	@echo "Updating dependencies..."
	go mod tidy
	go mod verify

# Docker构建
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

# Docker运行
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 $(DOCKER_IMAGE):$(DOCKER_TAG)

# Docker Compose
docker-compose-up:
	@echo "Starting services with Docker Compose..."
	docker-compose up -d

docker-compose-down:
	@echo "Stopping services..."
	docker-compose down

# 帮助信息
help:
	@echo "Available commands:"
	@echo "  make build          - Build the application"
	@echo "  make run           - Run the application"
	@echo "  make dev           - Run in development mode with hot reload"
	@echo "  make test          - Run tests"
	@echo "  make test-coverage - Run tests with coverage report"
	@echo "  make clean         - Clean build files"
	@echo "  make docs          - Generate Swagger documentation"
	@echo "  make migrate       - Run database migrations"
	@echo "  make migrate-down  - Rollback database migrations"
	@echo "  make fmt           - Format code"
	@echo "  make lint          - Run linter"
	@echo "  make deps          - Update dependencies"
	@echo "  make docker-build  - Build Docker image"
	@echo "  make docker-run    - Run Docker container"
	@echo "  make docker-compose-up   - Start services with Docker Compose"
	@echo "  make docker-compose-down - Stop services" 