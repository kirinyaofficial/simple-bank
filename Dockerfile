# Build Stage
FROM golang:1.26.1-alpine3.23 AS builder
WORKDIR /app
COPY . .
# Build the binary
RUN CGO_ENABLED=0 go build -o main main.go
RUN apk add --no-cache curl tar
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-amd64.tar.gz | tar xvz

# Run Stage
FROM alpine:3.23
WORKDIR /app

# Copy the binaries and files from builder
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY --from=builder /app/app.env .
COPY --from=builder /app/start.sh .
COPY --from=builder /app/wait-for.sh . 
COPY --from=builder /app/db/migration ./migration

# Ensure the script is executable
RUN chmod +x start.sh
RUN chmod +x wait-for.sh

EXPOSE 8080
# Use the script to run migrations before starting the app
ENTRYPOINT ["/app/start.sh"]
CMD ["/app/main"]