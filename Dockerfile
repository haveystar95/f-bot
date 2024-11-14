FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/app

COPY config.yaml /app/cmd/app/
COPY .env /app/cmd/app/
COPY db /app/cmd/app/db

RUN go build -o main main.go
RUN go run migration.go
EXPOSE 8084

CMD ["./main"]