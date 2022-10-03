FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod init goappaluno

RUN go build -o math

CMD ["./math"]
