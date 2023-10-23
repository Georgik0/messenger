package handlers

import (
	"context"
	"net/http"
	"testing"

	"messenger/interactorsDb"
)

func TestHandlerAddMessage_ServeHTTP(t *testing.T) {
	var conn *pgx.Conn

	ctx := context.Background()

	conn, err := interactorsDb.InitConnect(ctx, "127.0.0.1", "5431")
	if err != nil {
		t.Errorf("Please, run docker container with test_db, err = %v\n", err)
	}

	handler_input := &HandlerAddMessage{
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
			target: "/messages/add",
			json:   `{"chat": 1, "author": 1, "text": "hi"}`,
			input:  handler_input,
			want:   "Message id: 3\n",
		},
		{
			name:   "Ok",
			method: http.MethodPost,
			target: "/messages/add",
			json:   `{"chat": 0, "author": 1, "text": "hi"}`,
			input:  handler_input,
			want:   "Chat with id = 0 does not exist\n",
		},
		{
			name:   "Ok",
			method: http.MethodPost,
			target: "/messages/add",
			json:   `{"chat": 1, "author": 321, "text": "Привет"}`,
			input:  handler_input,
			want:   "User with id = 321 does not exist\n",
		},
	}

	CheckTestsRange(HndlInput(cases), t)
}
