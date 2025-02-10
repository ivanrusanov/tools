# Stage 1: Build stage
FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o gooli ./cmd/gooli/main.go

# Stage 2: Run stage
FROM alpine:edge

RUN apk update && apk add bash

WORKDIR /app

COPY --from=builder /app/gooli .

COPY run.sh /app/run.sh
RUN chmod +x /app/run.sh

# Set the entrypoint
ENTRYPOINT ["/app/run.sh"]