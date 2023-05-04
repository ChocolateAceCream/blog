FROM golang:alpine as builder

WORKDIR /go/src/github.com/ChocolateAceCream/blog/server
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="ChocolateAceCream"

WORKDIR /go/src/github.com/ChocolateAceCream/blog/server

COPY --from=0 /go/src/github.com/ChocolateAceCream/blog/server ./

EXPOSE 3000
ENTRYPOINT ./server -c release
