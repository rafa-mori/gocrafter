# Build stage
FROM golang:1.21-alpine AS builder

# Install git and ca-certificates (needed for go mod download)
RUN apk add --no-cache git ca-certificates

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o gocrafter .

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN adduser -D -s /bin/sh gocrafter

# Set working directory
WORKDIR /home/gocrafter

# Copy binary from builder stage
COPY --from=builder /app/gocrafter /usr/local/bin/gocrafter

# Copy templates directory
COPY --from=builder /app/templates ./templates

# Set proper ownership
RUN chown -R gocrafter:gocrafter /home/gocrafter

# Switch to non-root user
USER gocrafter

# Set entrypoint
ENTRYPOINT ["gocrafter"]

# Default command
CMD ["--help"]
