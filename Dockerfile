FROM golang:1.18-alpine3.16 AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

# RUN CGO_ENABLED=0 go build -o app ./cmd/app-http/.
RUN go build -o app ./cmd/app-http/.

# ========
FROM alpine:3.16

WORKDIR /

COPY --from=builder /app/files/etc/app_config/config.dev.yml ./files/etc/app_config/config.dev.yml
COPY --from=builder /app/app .

EXPOSE 80
ENTRYPOINT ["/app"]
