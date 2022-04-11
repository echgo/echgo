# With this file you can build the docker image
# We load the language golang:alpine from docker hub load, add the version, create the workdir, build the application
# After that we create a scratch image that based on alpine:latest, copy the files from the build image & start the application
FROM golang:alpine as build

RUN apk add --no-cache git

COPY . /tmp/go/src/app
WORKDIR /tmp/go/src/app

RUN go get ./
RUN go build -o echgo

FROM alpine:latest as scratch

COPY --from=build /tmp/go/src/app /go/src/app
WORKDIR /go/src/app

CMD ["./echgo"]