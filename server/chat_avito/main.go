package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"messenger/handlers"
	"net/http"
)

func main() {
	var err error
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://db_user:db_user_pass@myapp_db:5432/app_db")
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

	handlerAddUser := &handlers.HandlerAddUser{}
	handlerAddChat := &handlers.HandlerAddChat{}
	handlerAddMessage := &handlers.HandlerAddMessage{}
	handlerGetUserChats := &handlers.HandlerGetUserChats{}
	handlerGetChatMessages := &handlers.HandlerGetChatMessages{}

	handlerAddUser.InitHandler(&ctx, conn)
	handlerAddChat.InitHandler(&ctx, conn)
	handlerAddMessage.InitHandler(&ctx, conn)
	handlerGetUserChats.InitHandler(&ctx, conn)
	handlerGetChatMessages.InitHandler(&ctx, conn)

	http.Handle("/users/add", handlerAddUser)
	http.Handle("/chats/add", handlerAddChat)
	http.Handle("/messages/add", handlerAddMessage)
	http.Handle("/chats/get", handlerGetUserChats)
	http.Handle("/messages/get", handlerGetChatMessages)

	http.ListenAndServe(":9000", nil)
}
