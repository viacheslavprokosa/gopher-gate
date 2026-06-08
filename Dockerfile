FROM golang:1.26.4-alpine AS builder
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o gopher-gate ./cmd/gateway/main.go

FROM alpine:latest

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /src/gopher-gate .

RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 8080

CMD ["./gopher-gate"]