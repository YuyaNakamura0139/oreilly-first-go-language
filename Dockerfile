# https://hub.docker.com/_/golang
FROM golang:1.21.9-alpine

ENV ROOT=/go/src/work
WORKDIR ${ROOT}

RUN apk update \
    && apk add --no-cache git