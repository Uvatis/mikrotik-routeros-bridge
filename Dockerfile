# ---------- Build Stage ----------
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy module definition files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Static compilation
RUN CGO_ENABLED=0 GOOS=linux go build -o mikrotik-bridge .

# ---------- Runtime Stage ----------
FROM alpine:3.20

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/mikrotik-bridge .

# Expose the port used by the Go service
EXPOSE 8080

# Startup command
ENTRYPOINT ["./mikrotik-bridge"]

