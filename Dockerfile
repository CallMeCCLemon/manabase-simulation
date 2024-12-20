FROM golang:1.22.7-alpine3.20 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /app/main ./cmd/main.go

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/main .

ENV GODEBUG="http2debug=1"
EXPOSE 9000
CMD ["./main", "-server-port=9000"]
