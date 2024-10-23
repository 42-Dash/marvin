FROM golang:1.22.5

# All dependencies required to build the Go project
RUN apt-get update && apt-get install -y curl build-essential sudo

# Add 'student' user for running tests without permissions (default user in containers is root)
# RUN useradd -m student
RUN chmod 777 /root

# Switch to 'student' user
# USER student

WORKDIR /app

COPY /go.sum go.sum
COPY /go.mod go.mod
COPY /maps.json maps.json

COPY /internals internals
COPY /maps maps
COPY /pkg pkg
COPY /main.go main.go

RUN go build -o tester .
