# syntax=docker/dockerfile:1

# Step 1: Build the Go binary
FROM golang:1.23.2-alpine AS builder
WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy code
COPY . .

# Build app binary (adjust path if needed)
RUN go build -o server ./cmd/app

# Step 2: Create lightweight image
FROM alpine:latest
WORKDIR /app

# Copy binary
COPY --from=builder /app/server .

# Copy env file
COPY .env .

# Expose port
EXPOSE 8080

# Run the app
CMD ["./server"]
