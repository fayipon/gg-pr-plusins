# syntax=docker/dockerfile:1

##############################################
# 🧱 Builder 階段：使用 Go 官方鏡像構建二進制
##############################################
FROM golang:1.25-alpine AS builder

WORKDIR /app
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    GOMODCACHE=/go/pkg/mod \
    GOPROXY=https://proxy.golang.org,direct

# 安裝必要工具
RUN apk add --no-cache git tzdata

# 🧩 預先建立 module cache（提升重建速度）
COPY go.mod go.sum ./
RUN go mod download

# 🧩 再複製源碼
COPY . .

# 🧱 建構 users_rpc 二進位檔
RUN go build -ldflags="-s -w" -o users_rpc users.go


##############################################
# 🚀 Runtime 階段：最小化鏡像體積
##############################################
FROM alpine:3.19

WORKDIR /app
ENV TZ=Asia/Manila
RUN apk add --no-cache tzdata ca-certificates && update-ca-certificates

# 拷貝執行檔與配置
COPY --from=builder /app/users_rpc ./users_rpc
COPY --from=builder /app/etc ./etc

EXPOSE 9107

# 🟩 最佳實踐：使用 ENTRYPOINT 避免 shell 問題
ENTRYPOINT ["./users_rpc", "-f", "etc/users.yaml"]
