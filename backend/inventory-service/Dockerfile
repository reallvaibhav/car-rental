# Stage 1: Build the application
FROM golang:1.24-alpine AS builder

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Set up the working directory
WORKDIR /src/github.com/Car-Rental/backend

# Copy the proto files first (keeping the directory structure)
COPY proto/ ./proto/

# Copy the inventory-service directory
COPY inventory-service/ ./inventory-service/

# Set the working directory to the service directory for building
WORKDIR /src/github.com/Car-Rental/backend/inventory-service

# Download dependencies
RUN go mod download

# Build the service
RUN go build -o ./bin/inventory-service ./cmd/main.go

# Stage 2: Create a minimal image to run the service
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /src/github.com/Car-Rental/backend/inventory-service/bin/inventory-service ./

# Copy necessary runtime dependencies if any

EXPOSE 50052

# Command to run the service
ENTRYPOINT ["./inventory-service"]