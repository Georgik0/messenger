package initservices

import (
	"context"
	"github.com/jackc/pgx/v4"
	"messenger/handlers"
	handlerbuilder "messenger/init-services/handler-builder"
)

type provider struct {
	pattern string
	handler handler
}

func Run(ctx context.Context, connDB *pgx.Conn) {
	handlerBuilder := handlerbuilder.New(ctx, connDB)

	providers := []provider{
		{handler: &handlers.HandlerAddUser{}, pattern: "/users/add"},
		{handler: &handlers.HandlerAddChat{}, pattern: "/chats/add"},
		{handler: &handlers.HandlerAddMessage{}, pattern: "/messages/add"},
		{handler: &handlers.HandlerGetUserChats{}, pattern: "/chats/get"},
		{handler: &handlers.HandlerGetChatMessages{}, pattern: "/messages/get"},
	}

	for i := range providers {
		handlerBuilder.InitHandler(
			providers[i].handler,
			providers[i].pattern,
		)
	}
}
