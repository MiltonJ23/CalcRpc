# --- Stage 1: Builder ---
FROM golang:1.22-alpine AS builder

# Install system dependencies required for building (if any)
RUN apk add --no-cache git

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
# This ensures we don't re-download modules unless dependencies change
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
# -o main: names the output binary "main"
# ./cmd/server/main.go: points to your entry file
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server/main.go

# --- Stage 2: Runner ---
# Use a distroless static image for production
FROM gcr.io/distroless/static-debian11

WORKDIR /

# Copy the binary from the builder stage
COPY --from=builder /app/main /main

# Expose the ports defined in your architecture
# [cite_start]50051 for gRPC[cite: 119], 8080 for HTTP/Gin
EXPOSE 50051
EXPOSE 8080

# Run the binary
CMD ["/main"]