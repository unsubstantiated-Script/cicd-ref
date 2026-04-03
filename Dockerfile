ARG GO_VERSION=1.25
FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o api ./cmd/api

FROM alpine:3.20

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app
COPY --from=builder /app/api /app/api

USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/api"]