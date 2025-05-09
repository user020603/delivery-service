# Start from the official Golang image as the build environment
FROM golang:1.24.1 as builder

WORKDIR /app

# Install git and add CA certificates for Go modules
RUN apk add --no-cache git ca-certificates

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o delivery-service ./cmd/main.go

# Create a minimal runtime image
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder
COPY --from=builder /app/delivery-service .

# Copy any migration or seed files if needed (uncomment below if required)
# COPY migrations ./migrations

EXPOSE 8080

CMD ["./delivery-service"]