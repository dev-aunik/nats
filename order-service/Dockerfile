# Builder Stage
FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o order-service

# Runtime Stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/order-service .

CMD [ "./order-service" ]
