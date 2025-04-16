# ----------- Build Stage -------------
FROM golang:1.22-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build main app binary
RUN go build -o app ./cmd/main.go

# Build migration binary
RUN go build -o migrate ./internal/migrate/migrate_up.go

# ----------- Final Stage -------------
FROM alpine:latest

RUN apk add --no-cache ca-certificates postgresql-client

WORKDIR /root/

COPY --from=builder /app/app .
COPY --from=builder /app/migrate .
COPY migrations ./migrations
COPY .env .

COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

EXPOSE 8080

CMD ["./entrypoint.sh"]
