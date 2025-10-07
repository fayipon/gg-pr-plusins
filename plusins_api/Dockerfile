# syntax=docker/dockerfile:1

# ---------- Build 阶段 ----------
FROM golang:1.25-alpine AS builder
WORKDIR /app

# 安装依赖工具（git、tzdata）
RUN apk add --no-cache git tzdata

# 复制当前模块代码（不再用 ${PROJECT}_api）
COPY ./ ./

# 整理依赖并编译
RUN go mod tidy && go build -o plusins_api plusins_api.go

# ---------- Runtime 阶段 ----------
FROM alpine:3.19
WORKDIR /app

# 拷贝编译结果和配置文件
COPY --from=builder /app/plusins_api ./
COPY --from=builder /app/etc ./etc

ENV TZ=Asia/Manila

EXPOSE 8081

# 启动服务
CMD ["./plusins_api", "-f", "etc/plusins-api.yaml"]
