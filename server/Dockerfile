#Multi-stage build

#Stage 1 - Build the Go binary
FROM golang:1.24.1-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN go build -v -o devflow-console main.go #builds the binary inside /build



# Second Stage: Create a lightweight Alpine container for production
FROM alpine

# Create a non-root user
RUN adduser -S -D -H -h /app appuser
USER appuser

# Copy the binary from the builder stage
COPY --from=builder /build/devflow-console /app/

COPY .env /app/.env
WORKDIR /app

# Run the built Go App
CMD ["/app/devflow-console"]

