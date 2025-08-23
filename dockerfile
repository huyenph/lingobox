# ---------- Build stage ----------
FROM golang:1.24-alpine AS build

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o lingobox ./cmd/main.go

# ---------- Final stage ----------
FROM alpine:latest

WORKDIR /app

COPY --from=build /app/lingobox .

ENV APP_ENV=production

CMD ["./lingobox"]
