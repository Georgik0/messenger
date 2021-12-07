package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"messenger/handlers"
	"messenger/interactorsDb"
	"net/http"
)

func main() {
	var err error
	var conn *pgx.Conn
	ctx := context.Background()
	if err = interactorsDb.Init_connect_for_test(&conn, &ctx, "myapp_db", "5432"); err != nil {
		fmt.Errorf("Please, run docker container with db, err = %v\n", err)
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
