package handlers

import (
	"context"
	"github.com/jackc/pgx/v4"
	"messenger/interactorsDb"
	"net/http"
	"testing"
)

func TestHandlerAddChat_ServeHTTP(t *testing.T) {
	var conn *pgx.Conn
	ctx := context.Background()
	if err := interactorsDb.Init_connect_for_test(&conn, &ctx, "127.0.0.1", "5431"); err != nil {
		t.Errorf("Please, run docker container with test_db, err = %v\n", err)
	}

	handler_input := &HandlerAddChat{
		Ctx:    ctx,
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
			target: "/chats/add",
			json:   `{"name": "chat1", "users": [1, 2]}`,
			input:  handler_input,
			want:   "ERROR: duplicate key value violates unique constraint \"chat_name_key\" (SQLSTATE 23505)\n",
		},
		{
			name:   "Ok",
			method: http.MethodPost,
			target: "/chats/add",
			json:   `{"name": "chat1", "users": [1, 228]}`,
			input:  handler_input,
			want:   "User with id = 228 does not exist\n",
		},
		{
			name:   "Ok",
			method: http.MethodPost,
			target: "/chats/add",
			json:   `{"name": "chat3", "users": [0, 2]}`,
			input:  handler_input,
			want:   "User with id = 0 does not exist\n",
		},
		{
			name:   "Ok",
			method: http.MethodPost,
			target: "/chats/add",
			json:   `{"name": "chat3", "users": [1, 2]}`,
			input:  handler_input,
			want:   "Chat id 4\n",
		},
	}

	CheckTestsRange(HndlInput(cases), t)
}
