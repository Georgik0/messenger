FROM golang:1.17.2-alpine

WORKDIR /app

COPY go.mod ./
RUN go get "github.com/jackc/pgx/v4"
RUN go mod download

COPY ./ ./

RUN cd ./chat_avito \
    go build

EXPOSE 9000

CMD [ "./chat_avito" ]