# Build Stage
FROM golang:1.26.1-alpine3.23 AS builder
WORKDIR /app
COPY . .
# Added CGO_ENABLED=0 to ensure the binary is statically linked
# This prevents "file not found" errors in the final slim image
RUN CGO_ENABLED=0 go build -o main main.go

# Run Stage
FROM alpine:3.23
WORKDIR /app
COPY --from=builder /app/main .
# Copied app.env so the binary can find it
COPY app.env .

EXPOSE 8080
CMD ["/app/main"]