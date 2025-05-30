FROM golang:1.24-alpine

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copy go.mod and go.sum
COPY statistics-service/go.mod statistics-service/go.sum ./

# Copy the proto files
COPY proto/ ./proto/

# Copy the service's source code directories
COPY statistics-service/cmd/ ./cmd/
COPY statistics-service/internal/ ./internal/
COPY statistics-service/scripts/ ./scripts/
COPY statistics-service/internal/ ./internal/
# Copy any other top-level directories specific to this service

# Download dependencies
RUN go mod download

# Build the statistics service
RUN go build -o ./bin/statistics-service ./cmd/main.go

# Stage 2: Create a minimal image to run the service
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=0 /app/bin/statistics-service ./

# Run the application
CMD ["./statistics-service"]