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

# 🧩 預先建立 module cache（極大幅提升重建速度）
COPY go.mod go.sum ./
RUN go mod download

# 🧩 再複製源碼
COPY . .

# 🧱 建構二進制檔
RUN go build -ldflags="-s -w" -o plusins_rpc plusins_rpc.go


##############################################
# 🚀 Runtime 階段：最小化鏡像體積
##############################################
FROM alpine:3.19

WORKDIR /app
ENV TZ=Asia/Manila
RUN apk add --no-cache tzdata ca-certificates && update-ca-certificates

# 拷貝執行檔與配置
COPY --from=builder /app/plusins_rpc ./
COPY --from=builder /app/etc ./etc

EXPOSE 9103

# ✅ 最佳實踐：不使用 shell，以防 ENTRYPOINT 問題
ENTRYPOINT ["./plusins_rpc", "-f", "etc/plusins.yaml"]
