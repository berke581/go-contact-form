FROM golang:1.18.4-alpine3.16

LABEL maintainer="Berke <berke5sekiz1@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/go-contact-form

CMD ["./main"]