# syntax=docker/dockerfile:1

FROM golang:1.25-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git tzdata

# 复制当前目录（不再使用 ${PROJECT}_rpc）
COPY ./ ./

# 构建 Go 可执行文件
RUN go mod tidy && go build -o plusins_rpc plusins_rpc.go

FROM alpine:3.19
WORKDIR /app

# 拷贝执行文件和配置
COPY --from=builder /app/plusins_rpc ./
COPY --from=builder /app/etc ./etc

ENV TZ=Asia/Manila
EXPOSE 9103

CMD ["./plusins_rpc", "-f", "etc/plusins.yaml"]
