FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app

RUN go build -o main cmd/app/main_tg_bot.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=builder /app/main .
COPY db ./db
COPY templates /app/templates


ENTRYPOINT ["./main"]