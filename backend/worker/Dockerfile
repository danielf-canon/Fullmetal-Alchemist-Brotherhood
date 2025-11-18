FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN go build -o worker_bin ./worker

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/worker_bin /app/worker

RUN chmod +x /app/worker

CMD ["/app/worker"]
