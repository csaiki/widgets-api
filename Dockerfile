FROM golang:1.8.3-alpine3.6

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh && \
    mkdir -p /app

WORKDIR /go/src/app

ADD . /go/src/app

RUN go get && go build ./main.go
