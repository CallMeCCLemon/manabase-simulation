FROM golang:1.22.7-alpine3.20 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /app/main ./cmd/main.go

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8889
CMD ["./main"]
