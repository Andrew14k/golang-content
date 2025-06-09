# Stage 1: Build the Go app
FROM golang:1.24 AS builder

WORKDIR /app

# Copy go.mod first to cache deps
COPY go.mod go.sum ./   
RUN go mod download

# Copy the rest of the source
COPY . ./

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp ./cmd/data-access

# Stage 2: Minimal runtime image
FROM alpine:latest

WORKDIR /root/

# Copy binary from builder stage
COPY --from=builder /app/myapp ./

# Run the binary
CMD ["./myapp"]
