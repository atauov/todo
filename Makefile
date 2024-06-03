IMAGE_NAME=postgres
CONTAINER_NAME=todo-container
PORT=5432
MIGRATION_DIR=./schema
DATABASE_URL='postgres://postgres:password@localhost:5432/todo-db?sslmode=disable'

.PHONY: run docker-build docker-run migrate-up migrate-down
run:
	go run cmd/todo/main.go

docker-build:
	docker build -t $(IMAGE_NAME) .

docker-run:
	docker run -d --name $(CONTAINER_NAME) -p $(PORT):5432 $(IMAGE_NAME)

migrate-up:
	migrate -path $(MIGRATION_DIR) -database $(DATABASE_URL) up

migrate-down:
	migrate -path $(MIGRATION_DIR) -database $(DATABASE_URL) down