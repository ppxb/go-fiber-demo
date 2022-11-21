FROM golang:1.19-alpine AS builder

WORKDIR /build
ENV GOPROXY https://goproxy.cn
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLE=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o courseapi .

FROM alpine

COPY --from=builder ["/build/courseapi","/build/deploy/wait-for.sh","/"]

EXPOSE 8848
