BINARY_NAME = genealogy
DOCKER_IMAGE = genealogy-image

all: build

build:
	docker build -t $(DOCKER_IMAGE) .

clean:
	docker-compose down
	docker image rm -f $(DOCKER_IMAGE)

test:
	docker-compose run --rm go test ./...

run: build
	docker-compose up -d
#	docker-compose run --rm go run main.go

env:
	docker run --rm $(DOCKER_IMAGE) go version && docker run --rm $(DOCKER_IMAGE) go env

deps:
	docker run --rm $(DOCKER_IMAGE) go mod tidy
	
migrate:
	docker-compose run --rm go run cmd/migrate/main.go

migrate-up:
	docker-compose run --rm go run cmd/migrate/main.go up

migrate-down:
	docker-compose run --rm go run cmd/migrate/main.go down

migrate-create:
	docker-compose run --rm go run cmd/migrate/main.go create $(name)

.PHONY: all build clean test run env deps
