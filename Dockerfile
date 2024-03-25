# syntax=docker/dockerfile:1

FROM golang:1.22

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /server-bin cmd/server/main.go

EXPOSE 3000

CMD ["/server-bin"]
