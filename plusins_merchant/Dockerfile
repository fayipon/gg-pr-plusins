# syntax=docker/dockerfile:1

##############################
# 🧱 Build 階段
##############################
FROM golang:1.25-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git tzdata ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    GOMODCACHE=/go/pkg/mod \
    GOPROXY=https://proxy.golang.org,direct

RUN go build -ldflags="-s -w" -o plusins_merchant main.go


##############################
# 🚀 Runtime 階段
##############################
FROM alpine:3.19
WORKDIR /app

RUN apk add --no-cache tzdata ca-certificates && update-ca-certificates

COPY --from=builder /app/plusins_merchant ./plusins_merchant
COPY --from=builder /app/etc ./etc

ENV TZ=Asia/Manila

EXPOSE 8082

CMD ["./plusins_merchant", "-f", "etc/plusins_merchant.yaml"]
