# Build Stage
FROM golang:1.25.4-alpine3.21 AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o ledger-service .

# Final Stage
FROM alpine:3.21

WORKDIR /app

# Create a non-root user with explicit UID 1000
RUN addgroup -S -g 1000 ledger && adduser -S -u 1000 -G ledger ledger

# Ensure the working directory is owned by the non-root user and accessible
RUN chown -R ledger:ledger /app && chmod 755 /app

# Copy the binary with correct ownership
COPY --from=builder --chown=ledger:ledger /app/ledger-service /app/ledger-service

# Ensure execution permissions explicitly
RUN chmod 755 /app/ledger-service

# Switch to non-root user
USER ledger

# Expose port 8080
EXPOSE 8080

# Command to run the executable (Absolute path)
CMD ["/app/ledger-service"]
