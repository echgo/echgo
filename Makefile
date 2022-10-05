# Here you can reformat, check, test or publish the software.
BINARY_NAME=echgo
GIT_TAG=$(shell git describe --tags --abbrev=0)
VERSION=$(if $(GIT_TAG),$(GIT_TAG),unavailible)
DOCKER_HUB_USERNAME=echgo

fmt:
	@go fmt ./...

vet:
	@go vet ./...

test:
	@go test ./...

lint:
	@golangci-lint run ./...

doc:
	@godoc -play=true -goroot=/usr/local/go -http=:6060

version:
	@echo "version: ${VERSION}"

build:
	go build -o ${BINARY_NAME}

docker-build:
	docker buildx b \
		--platform linux/amd64,linux/arm64,linux/arm/v7 \
		-t ${DOCKER_HUB_USERNAME}/${BINARY_NAME}:${VERSION} -t ${DOCKER_HUB_USERNAME}/${BINARY_NAME}:latest .

docker-push:
	docker push -a ${DOCKER_HUB_USERNAME}/${BINARY_NAME}