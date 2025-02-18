# Build stage
FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app
COPY ./app .
RUN go build -o main main.go

# Run Stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/index.html .

EXPOSE 8080
CMD ["/app/main"]