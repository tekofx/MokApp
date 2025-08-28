# --- BUILD STAGE ---
# Use the official Golang image to build the application
FROM golang:alpine AS builder

# Set the working directory
WORKDIR /docker-go

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application (output binary named "app")
RUN apk add --no-cache gcc musl-dev
RUN CGO_ENABLED=1 go build -o /app main.go

# --- FINAL STAGE ---
# Use a minimal Alpine image to reduce size
FROM alpine:latest

# Set a working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app /app


ENV APP_ENV=prod
EXPOSE 8080

# Command to run the binary
CMD ["./app"]
