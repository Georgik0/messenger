FROM golang:1.17.2-alpine

WORKDIR /app

COPY go.mod ./
RUN go get "github.com/jackc/pgx/v4"
RUN go mod download

COPY chat_avito.go checker.go ./

RUN go build -o /chat_avito

EXPOSE 9000

CMD [ "/chat_avito" ]