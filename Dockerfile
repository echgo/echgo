# With this file you can build the docker image
# We load the language golang from docker hub, define the environment variables & load, build & start the application
FROM golang:latest

ENV GOTIFY_DOMAIN=''
ENV GOTIFY_KEY=''

ENV TELEGRAM_API_TOKEN=''

ENV SMTP_HOST=''
ENV SMTP_PORT=''
ENV SMTP_USERNAME=''
ENV SMTP_PASSWORD=''

ENV WEBHOOK_DOMAIN=''

COPY . /go/src/app
WORKDIR /go/src/app

RUN go get ./
RUN go build -o app

CMD ./app