FROM golang:alpine3.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /ws ./cmd/ws


FROM alpine:latest

WORKDIR /root/

COPY --from=builder /ws .

CMD ['./ws']




