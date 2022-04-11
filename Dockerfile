# With this file you can build the docker image
# We load the language golang alpine from docker hub load, add the version, create the workdir, build the application
# After that we create a scratch image, copy the files from the build image & start the application
FROM golang:alpine as build

RUN apk add --no-cache git

COPY . /go/src/app
WORKDIR /go/src/app

RUN go get ./
RUN go build -o app

FROM scratch

COPY --from=build /go/src/app /go/src/app

CMD ./app