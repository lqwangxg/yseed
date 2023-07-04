# syntax=docker/dockerfile:1

#FROM golang:1.19
FROM bash

# Set destination for COPY
WORKDIR /go/bin

# Download Go modules
COPY yseed ./
COPY config.yml ./
# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8000

# Run
CMD ["./yseed"]