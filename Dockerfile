# With this file you can build the docker image. We load the language
# golang:alpine from docker hub load, change the workdir, load the
# modules & build the application. After that we create a scratch image that based
# on alpine:latest, copy the files from the build image & start the application.
FROM golang:alpine AS build

RUN apk --no-cache add \
    make

WORKDIR /tmp/src

COPY go.mod .

COPY . .
RUN make build

FROM alpine:latest AS production

RUN apk --no-cache add \
    tzdata \
    curl

ENV TZ=Europe/Berlin

WORKDIR /app

COPY files/ files/
COPY --from=build /tmp/src/echgo .

ENTRYPOINT ["/app/echgo"]