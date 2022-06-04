FROM golang:1.18.3 AS builder

WORKDIR /usr/src/app
COPY . .
RUN go mod download
RUN go build -o /http_hello_world

FROM gcr.io/distroless/base-debian11
WORKDIR /
COPY --from=builder /http_hello_world /http_hello_world
ENTRYPOINT ["/http_hello_world"]
