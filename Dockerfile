# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install git and necessary tools
RUN apk add --no-cache git

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy all source files
COPY . .

# Build the binary
RUN go build -o parkping ./cmd/server/main.go

# Run stage
FROM alpine:latest

WORKDIR /app

# Copy the binary
COPY --from=builder /app/parkping .

# Copy .env file if needed
COPY .env ./

# Expose port
EXPOSE 8080

# Command to run
CMD ["./parkping"]
