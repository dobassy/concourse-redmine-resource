FROM golang:alpine as builder
# for go test execution
RUN apk add --no-cache gcc
ENV CGO_ENABLED 0

RUN mkdir -p /assets
WORKDIR /go/src
COPY . /go/src/github.com/dobassy/concourse-redmine-resource
RUN go build -o /assets/in github.com/dobassy/concourse-redmine-resource/in/cmd/main.go
RUN go build -o /assets/out github.com/dobassy/concourse-redmine-resource/out/cmd/main.go
RUN go build -o /assets/check github.com/dobassy/concourse-redmine-resource/check/cmd/main.go
WORKDIR /go/src/github.com/dobassy/concourse-redmine-resource
RUN set -e; go test -v ./...

FROM alpine:edge AS resource
RUN apk add --no-cache bash tzdata ca-certificates
COPY --from=builder assets/ /opt/resource/
RUN chmod +x /opt/resource/*
