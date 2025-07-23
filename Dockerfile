FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o vaultui vault/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/vaultui .

CMD ["./vaultui"]