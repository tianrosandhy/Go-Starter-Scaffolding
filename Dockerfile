FROM golang:1.24.3-alpine AS builder

# Install Git (required for fetching the dependencies)
RUN apk update && apk add --no-cache git

# Set the workdir
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy project files
COPY . .

# Build the executable binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Start a new stage with a minimal alpine image
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .

# Expose port 9009
EXPOSE 9009

# Run the binary
CMD ["./main"]
