# Build Stage
FROM golang:1.22 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o app

# Final Stage
FROM debian:bullseye-slim

WORKDIR /root/
COPY --from=builder /app/app .
EXPOSE 4000

CMD ["./app"]
