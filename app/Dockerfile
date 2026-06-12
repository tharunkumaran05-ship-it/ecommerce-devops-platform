# Build Stage
FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod .
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ecommerce-app .

# Runtime Stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/ecommerce-app .

EXPOSE 8080

CMD ["./ecommerce-app"]