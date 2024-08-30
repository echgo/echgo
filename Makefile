# Here you can reformat, check, test or publish the software.
.PHONY: test coverage doc version build docker-buildx

BINARY_NAME=echgo
APP_PATH=cmd/${BINARY_NAME}/main.go
GIT_TAG=$(shell git describe --tags --abbrev=0)
VERSION=$(if $(GIT_TAG),$(GIT_TAG),unavailible)
DOCKER_HUB_USERNAME=echgo

test:
	@go test ./...

doc:
	@godoc -play=true -goroot=/usr/local/go -http=:6060

version:
	@echo "version: ${VERSION}"

build:
	go mod download
	go build -o ${BINARY_NAME} ${APP_PATH}

docker-buildx:
	docker buildx build \
		--platform linux/amd64,linux/arm64,linux/arm/v7 \
		--push \
		-t ${DOCKER_HUB_USERNAME}/${BINARY_NAME}:${VERSION} -t ${DOCKER_HUB_USERNAME}/${BINARY_NAME}:latest .