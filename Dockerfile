FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 go build -o main .

FROM debian:buster-slim

WORKDIR /root/

COPY --from=builder /app/main .

COPY --from=builder /app/infrastructure/config /root/infrastructure/config

EXPOSE 8080

CMD ["./main"]


