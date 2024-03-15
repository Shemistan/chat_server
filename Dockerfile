FROM golang:1.20.3-alpine AS builder

COPY . github.com/Shemistan/chat_server
WORKDIR github.com/Shemistan/chat_server

RUN go mod download
RUN go build -o ./bin/user_server cmd/chat_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /go/github.com/Shemistan/chat_server/bin/user_server .
COPY --from=builder /go/github.com/Shemistan/chat_server/.env .

CMD ["./user_server", "--config-path", ".env" ]
