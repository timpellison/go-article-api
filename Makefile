# sort of reminiscient of
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")

.PHONY: swagger
swagger:
	echo "Running swagger stuff"
	swag init --dir ./cmd/http-server,./api,./dto -o ./docs

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)

.PHONY: dock
dock:
	$(MAKE) swagger
	$(MAKE) undock
	docker build . -f ./.local/Dockerfile -t articleapi:latest --progress=plain
	docker compose -f ./.local/docker-compose.yaml up -d

.PHONY: undock
undock:
	docker compose -f ./.local/docker-compose.yaml down
	docker image rm articleapi

.PHONY: run
run:
	$(MAKE) swagger
	CGO_ENABLED=1 go run ./cmd/hello-server