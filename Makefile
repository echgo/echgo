# Here you can reformat, check, test or publish the software
DOCKER_IMAGE_NAME=echgo/echgo
DOCKER_IMAGE_VERSION=$(shell git describe --tags --abbrev=0)

reformat:
	go fmt ./...

check:
	golangci-lint run ./...

test-build:
	docker build -t ${DOCKER_IMAGE_NAME}:testing .

test-push:
	docker push ${DOCKER_IMAGE_NAME}:testing

publish:
	docker buildx b \
	--platform linux/amd64,linux/arm64,linux/arm/v7 \
	--push \
	-t ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_VERSION} -t ${DOCKER_IMAGE_NAME}:latest .
