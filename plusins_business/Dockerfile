# syntax=docker/dockerfile:1

##############################
# 🧱 Build 階段
##############################
FROM golang:1.25-alpine AS builder
WORKDIR /app

# 安裝必要工具
RUN apk add --no-cache git tzdata ca-certificates

# ✅ 預先複製 go.mod / go.sum 以利用 Docker layer cache
COPY go.mod go.sum ./
RUN go mod download

# ✅ 再複製專案代碼
COPY . .

# ✅ 加入環境變數與快取設定
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    GOMODCACHE=/go/pkg/mod \
    GOPROXY=https://proxy.golang.org,direct

# ✅ 編譯可執行檔（strip debug symbols 減少體積）
RUN go build -ldflags="-s -w" -o plusins_business plusins_business.go


##############################
# 🚀 Runtime 階段
##############################
FROM alpine:3.19
WORKDIR /app

# 安裝時區與憑證
RUN apk add --no-cache tzdata ca-certificates && update-ca-certificates

# 拷貝執行檔與配置
COPY --from=builder /app/plusins_business ./plusins_business
COPY --from=builder /app/etc ./etc

# 設定時區
ENV TZ=Asia/Manila

# 開放端口（⚙️ 比 plusins_api 多一號）
EXPOSE 8082

# 啟動服務
CMD ["./plusins_business", "-f", "etc/plusins_business.yaml"]