FROM golang:1.13-alpine

RUN apk add --no-cache bash

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
