FROM golang:latest AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# COMPILACIÓN ESTÁTICA (clave)
ENV CGO_ENABLED=0
RUN GOOS=linux GOARCH=amd64 go build -o main /app/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]
