package handlers

import (
	"context"
	"net/http"
	"testing"

	"messenger/interactorsDb"
)

func TestHandlerGetUserChats_ServeHTTP(t *testing.T) {
	var conn *pgx.Conn

	ctx := context.Background()

	conn, err := interactorsDb.InitConnect(ctx, "127.0.0.1", "5431")
	if err != nil {
		t.Errorf("Please, run docker container with test_db, err = %v\n", err)
	}

	handler_input := &HandlerGetUserChats{
		Ctx:    ctx,
		connDB: conn,
	}
	cases := []struct {
		name   string
		method string
		target string
		json   string
		input  http.Handler
		want   string
	}{
		{
			name:   "Ok",
			method: http.MethodPost,
			target: "/chats/get",
			json:   `{"user": 3}`,
			input:  handler_input,
			want:   "The user's chat_id: 2\n",
		},
		{
			name:   "Ok",
			method: http.MethodPost,
			target: "/chats/get",
			json:   `{"user": 0}`,
			input:  handler_input,
			want:   "User with id = 0 does not exist\n",
		},
	}

	CheckTestsRange(HndlInput(cases), t)
}
