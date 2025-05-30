FROM golang:1.24-alpine

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copy go.mod and go.sum
COPY booking-service/go.mod booking-service/go.sum ./

# Copy the proto files
COPY proto/ ./proto/

# Copy the service's source code directories
COPY booking-service/cmd/ ./cmd/
COPY booking-service/internal/ ./internal/
COPY booking-service/cache/ ./cache/
COPY booking-service/config/ ./config/

# Download dependencies
RUN go mod download

# Build the booking service
RUN go build -o ./bin/booking-service ./cmd/main.go

# Stage 2: Create a minimal image to run the service
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=0 /app/bin/booking-service ./

# Copy necessary runtime dependencies if any

EXPOSE 50053

# Command to run the service
ENTRYPOINT ["./booking-service"] 