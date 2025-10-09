FROM golang:alpine3.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /http ./cmd/api


FROM alpine:latest

WORKDIR /root/

COPY --from=builder /http .

CMD ['./http']




