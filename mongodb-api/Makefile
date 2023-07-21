PWD := $(shell pwd)

DOCKER_SERVICES := "$(shell cd docker && docker-compose ps --services)"
DOCKER_RUNNINGS := "$(shell cd docker && docker-compose ps --services --filter 'status=running')"

all: build run

build:
	@echo "Building mongodb-api image"
	@go build -mod vendor -o $(PWD)/bin/mongodb-api
	@mv $(PWD)/bin/mongodb-api $(PWD)/docker/mongodb-api/mongodb-api
	@cd docker && docker-compose build
	@rm $(PWD)/docker/mongodb-api/mongodb-api -r

run:
	@if [ $(DOCKER_SERVICES) != $(DOCKER_RUNNINGS) ]; then \
		echo "Run docker containers"; \
		cd docker && docker-compose up -d; \
	else \
		echo "Restart docker containers"; \
		cd docker && docker-compose restart; \
	fi
