FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app ./cmd/api/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /bin/app .
COPY --from=builder /app/docs ./docs
COPY --from=builder /app/.env .env
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/build/entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

# !!! Match with APP_PORT in .env file
EXPOSE 8083

ENTRYPOINT ["/entrypoint.sh"]