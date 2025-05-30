FROM golang:1.24-alpine

WORKDIR /app

# Copy go.mod and go.sum from the service directory
COPY go.mod go.sum ./

# Copy the generated proto files from the workspace root
COPY ../proto/ ./proto/

# Copy the service's source code directories
COPY cmd/ ./cmd/
COPY internal/ ./internal/
# Copy any other top-level directories specific to this service
# COPY config/ ./config/ # If config is in a top-level dir
# COPY handlers/ ./handlers/ # If handlers are in a top-level dir

# Download dependencies
RUN go mod download

# Build the api-gateway binary
RUN go build -o ./bin/api-gateway ./cmd/main.go

# Stage 2: Create a minimal image to run the service
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=0 /app/bin/api-gateway ./

# Copy necessary runtime dependencies if any

EXPOSE 8080

# Command to run the api-gateway
ENTRYPOINT ["./api-gateway"] 