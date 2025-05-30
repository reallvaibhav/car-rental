# Stage 1: Build the application
FROM golang:1.24-alpine AS builder

# Install build dependencies and ca-certificates for HTTPS
RUN apk add --no-cache gcc musl-dev ca-certificates git

WORKDIR /app

# Set Go environment variables for better module handling
ENV GOPROXY=https://proxy.golang.org,direct
ENV GOSUMDB=sum.golang.org
ENV CGO_ENABLED=0
ENV GOOS=linux

# Copy go.mod and go.sum first for better Docker layer caching
COPY go.mod go.sum ./

# Copy the proto files and their go.mod
COPY proto/ ./proto/

# Download dependencies with timeout and retry settings
RUN go mod download -x

# Copy the service's source code
COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY cache/ ./cache/
COPY config/ ./config/

# Build the booking service with static linking
RUN go build -ldflags="-w -s" -o ./bin/booking-service ./cmd/main.go

# Stage 2: Create a minimal image to run the service
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/bin/booking-service ./

EXPOSE 50053

# Command to run the service
ENTRYPOINT ["./booking-service"]