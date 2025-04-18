# Stage 1: Build the Go binary
FROM golang:1.23.4-alpine AS builder
WORKDIR /app

# Copy dependency files first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy all source files and directories
COPY . .

# Build the binary (replace "./main.go" if your entry point is different)
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./main.go

# ------------------------------------------------

# Stage 2: Runtime image
FROM alpine:latest
WORKDIR /app

# Copy the compiled binary
COPY --from=builder /app/main .

# Security: Run as non-root user
RUN addgroup -S nonroot && adduser -S nonroot -G nonroot
USER nonroot:nonroot

# Expose port (adjust as needed)
EXPOSE 3000

# Run the application
ENTRYPOINT ["/app/main"]