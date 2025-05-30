FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy context is the parent directory, so specify full path
COPY nats-service/go.mod .
COPY nats-service/go.sum .
COPY nats-service/main.go .

RUN go mod download && \
    go build -o nats-service main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/nats-service .

ENV NATS_URL=nats://nats:4222

CMD ["./nats-service"]