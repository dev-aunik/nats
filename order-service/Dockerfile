FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o order-service

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/order-service .

CMD ["./order-service"]
