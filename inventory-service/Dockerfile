FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o inventory-service

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/inventory-service .

CMD ["./inventory-service"]
