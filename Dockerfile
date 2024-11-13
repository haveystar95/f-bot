FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/app

COPY config.yaml /app/cmd/app/

RUN go build -o main .

EXPOSE 8084

CMD ["./main"]