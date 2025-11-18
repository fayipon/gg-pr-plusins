# syntax=docker/dockerfile:1

FROM golang:1.25-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN apk add --no-cache git tzdata

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o users_rpc users.go

FROM alpine:3.19

WORKDIR /app

RUN apk add --no-cache tzdata ca-certificates

COPY --from=builder /app/users_rpc ./
COPY --from=builder /app/etc ./etc

EXPOSE 9107

ENTRYPOINT ["./users_rpc", "-f", "etc/users.yaml"]
