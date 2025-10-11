FROM golang:alpine3.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /http ./cmd/http


FROM alpine:latest

WORKDIR /root/

COPY --from=builder /http /http

ENTRYPOINT /http




