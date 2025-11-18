# syntax=docker/dockerfile:1


FROM golang:1.25-alpine AS builder

WORKDIR /app
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    GOMODCACHE=/go/pkg/mod \
    GOPROXY=https://proxy.golang.org,direct

RUN apk add --no-cache git tzdata

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o plusins_rpc plusins_rpc.go


FROM alpine:3.19

WORKDIR /app
ENV TZ=Asia/Manila
RUN apk add --no-cache tzdata ca-certificates && update-ca-certificates


COPY --from=builder /app/plusins_rpc ./
COPY --from=builder /app/etc ./etc

EXPOSE 9103

ENTRYPOINT ["./users_rpc", "-f", "etc/users.yaml"]
