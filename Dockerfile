FROM golang:1.18-alpine
LABEL maintainer="mochammad.lukman@gmail.com"

RUN apk update && \
    apk add bash git && \
    apk add gcc && \
    apk add musl-dev && \
    apk add curl && \
    apk add --update make

COPY . /home/golang/lib
WORKDIR /home/golang/lib