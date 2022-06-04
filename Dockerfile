FROM golang:1.18.3

WORKDIR /usr/src/app

COPY . .

RUN go mod download

RUN go build -o /http_hello_world

EXPOSE 3000

ENTRYPOINT ["/http_hello_world"]
