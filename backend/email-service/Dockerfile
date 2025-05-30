FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum
COPY email-service/go.mod email-service/go.sum ./

# Download dependencies
RUN go mod download

# Copy the service's source code directories
COPY email-service/cmd/ ./cmd/
COPY email-service/internal/ ./internal/

# Build the service
RUN go build -o email-service ./cmd/main.go

# Create minimal production image
FROM alpine:latest

WORKDIR /app

# Copy binary from builder stage
COPY --from=0 /app/email-service .

# Set environment variables for NATS (the rest will be provided via docker-compose)
ENV NATS_URL=nats://nats:4222

# Run the service
CMD ["./email-service"]
