package handlers

import (
	"context"
	"github.com/jackc/pgx/v4"
	"messenger/interactorsDb"
	"net/http"
	"testing"
)

func TestHandlerGetChatMessages_ServeHTTP(t *testing.T) {
	var conn *pgx.Conn
	ctx := context.Background()
	if err := interactorsDb.Init_connect_for_test(&conn, &ctx, "127.0.0.1", "5431"); err != nil {
		t.Errorf("Please, run docker container with test_db, err = %v\n", err)
	}

	handler_input := &HandlerGetChatMessages{
		Ctx:    &ctx,
		ConnDB: conn,
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
			target: "/messages/get",
			json:   `{"chat": 1}`,
			input:  handler_input,
			want:   "The chat's messages:\nuser1 write in chat1\nhi\n",
		},
		{
			name:   "Ok",
			method: http.MethodPost,
			target: "/messages/get",
			json:   `{"chat": 112}`,
			input:  handler_input,
			want:   "Chat with id = 112 does not exist\n",
		},
	}

	CheckTestsRange(HndlInput(cases), t)
}
