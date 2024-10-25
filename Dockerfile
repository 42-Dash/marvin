FROM golang:1.22.5

# All dependencies required to build the Go project
RUN apt-get update && apt-get install -y curl build-essential sudo

WORKDIR /app

COPY /go.sum go.sum
COPY /go.mod go.mod
COPY /maps.json maps.json

COPY /internals internals
COPY /maps maps
COPY /pkg pkg
COPY /main.go main.go

RUN go build -o tester .
