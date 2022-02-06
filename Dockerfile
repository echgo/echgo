# With this file you can build the docker image
# We load the language golang from docker hub load, build & start the application
FROM golang:latest

COPY . /go/src/app
WORKDIR /go/src/app

RUN go get ./
RUN go build -o app

CMD ./app