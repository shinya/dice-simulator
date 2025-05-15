FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY main.go .

RUN go build -o dice-simulator main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/dice-simulator .

ENTRYPOINT ["/app/dice-simulator"]
