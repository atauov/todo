IMAGE_NAME=postgres
CONTAINER_NAME=todo-container
PORT=5432

.PHONY: run docker-build docker-run
run:
	go run cmd/todo/main.go

docker-build:
	docker build -t $(IMAGE_NAME) .

docker-run:
	docker run -d --name $(CONTAINER_NAME) -p $(PORT):5432 $(IMAGE_NAME)


