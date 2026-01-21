# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go_soap_test .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/go_soap_test .

# Copy JSON files (se existirem, serão criados se não existirem)
COPY --from=builder /app/*.json ./

# Expose port (Railway vai definir a porta via variável de ambiente)
EXPOSE 8000

# Run the application
CMD ["./go_soap_test"]
