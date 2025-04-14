FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -o server .

FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder /app/server /app/server

ENTRYPOINT ["/app/server"]
