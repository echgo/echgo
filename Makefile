# Here you can reformat, check, test or publish the software
BINARY_NAME=echgo
DOCKER_IMAGE_NAME=echgo/echgo
DOCKER_IMAGE_VERSION=$(shell git describe --tags --abbrev=0)

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run ./...

test-build:
	docker build -t ${DOCKER_IMAGE_NAME}:testing .

test-push:
	docker push ${DOCKER_IMAGE_NAME}:testing

test-init:
	docker run --name ${BINARY_NAME}-testing-init -d --rm \
    	-v /etc/${BINARY_NAME}/configuration:/app/files/configuration \
    	${DOCKER_IMAGE_NAME}:testing

test-run:
	docker run --name ${BINARY_NAME}-testing -d --restart always \
        -v /etc/${BINARY_NAME}/configuration:/app/files/configuration \
        -v /var/lib/${BINARY_NAME}/notification:/app/files/notification \
        ${DOCKER_IMAGE_NAME}:testing

publish:
	git checkout master
	docker buildx b \
		--platform linux/amd64,linux/arm64,linux/arm/v7 \
		--push \
		-t ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_VERSION} -t ${DOCKER_IMAGE_NAME}:latest .
	git checkout development