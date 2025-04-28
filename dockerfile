# Build Stage
FROM golang:1.24 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o app

# Final Stage
FROM debian:bullseye-slim

WORKDIR /root/
COPY --from=builder /app/app .
COPY --from=builder /app/.env .
EXPOSE 4000

CMD ["./app"]
