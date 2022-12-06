ARG GO_VERSION=1.16.6

FROM golang:${GO_VERSION}-alpine AS builder

RUN go env -w GOPROXY=direct
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates && update-ca-certificates

# Changing work directory
WORKDIR /src

# Installing dependencies
COPY ./go.mod ./go.mod ./
RUN go mod download

# Copiying source file folders
COPY events events
COPY repository repository
COPY database database
COPY search search
COPY models models
COPY feed-service feed-service
COPY query-service query-service

RUN go install ./...

# Alpine images are lightweight and secure
FROM alpine:3.11
WORKDIR /usr/bin

COPY --from=builder /go/bin .
