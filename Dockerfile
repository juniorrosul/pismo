FROM golang:1.15

WORKDIR /app/src/pismo-api

ENV GOPATH=/app

COPY . /app/src/pismo-api

RUN go build main.go

ENTRYPOINT ["./main"]

EXPOSE 8080
