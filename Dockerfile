# Dockerfile for the echgo application to send notifications via txt, json or xml.
# This file details the process for constructing a lightweight and efficient Docker image using a multi-stage build process.
# The chosen base is Alpine Linux for its minimalistic size, while still providing necessary functionalities.

########################################################################################################################

# This phase uses the Alpine-based Go image to compile the source code of the application.
# By parameterizing the Go version, it becomes straightforward to maintain and modify in the future.
ARG GO_VERSION=1.23
FROM golang:${GO_VERSION}-alpine AS build
RUN apk add --no-cache make
WORKDIR /tmp/src
COPY . .
RUN make build

########################################################################################################################

# The final preparation phase for the production-ready image. Essential system packages and set the correct timezone is set.
FROM alpine:latest AS production
LABEL org.opencontainers.image.title="echgo" \
  org.opencontainers.image.description="The easy way to send notifications via txt, json or xml file." \
  org.opencontainers.image.vendor="Jonas Kwiedor" \
  org.opencontainers.image.source="https://github.com/echgo/echgo"
RUN apk add --no-cache tzdata
ENV TZ=Europe/Berlin
WORKDIR /app
COPY --from=build /tmp/src/echgo .
CMD ["/app/echgo"]