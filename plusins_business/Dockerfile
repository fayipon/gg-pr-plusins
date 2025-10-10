FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache git tzdata && \
    go mod tidy && \
    go build -o plusins_business_api plusins_business.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/plusins_business_api .
COPY etc ./etc
EXPOSE 8082
ENTRYPOINT ["./plusins_business_api", "-f", "etc/plusins_business.yaml"]
