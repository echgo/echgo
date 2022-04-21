# With this file you can build the docker image
# We load the language golang:alpine from docker hub load, add the version, create the workdir, build the application
# After that we create a scratch image that based on alpine:latest, copy the files from the build image & start the application
FROM golang:alpine AS build

RUN apk --no-cache add git

COPY . /tmp/go/src/app/
WORKDIR /tmp/go/src/app/

RUN go get ./
RUN go build -o echgo

FROM alpine:latest AS scratch

RUN apk --no-cache add tzdata

ENV TZ=Europe/Berlin

COPY --from=build /tmp/go/src/app/echgo /tmp/go/src/app/files/* /go/src/app/
WORKDIR /go/src/app/

CMD ["./echgo"]