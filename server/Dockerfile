FROM golang:1.19-alpine3.17 AS builder

ENV PROJECT_PATH=/usr/local/go/src/messenger

COPY . $PROJECT_PATH

WORKDIR $PROJECT_PATH

RUN go mod download && \
    CGO_ENABLED=0 go build -o messenger cmd/main.go

FROM alpine:3.17

WORKDIR /

COPY --from=builder /usr/local/go/src/messenger/messenger .

EXPOSE 9000

CMD ["./messenger"]
