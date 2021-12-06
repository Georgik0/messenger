package handlers

import (
	"bytes"
	"context"
	"github.com/jackc/pgx/v4"
	"messenger/interactorsDb"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerAddChat_ServeHTTP(t *testing.T) {
	var conn *pgx.Conn
	ctx := context.Background()
	if err := interactorsDb.Init_connect_for_test(&conn, &ctx); err != nil {
		t.Errorf("Please, run docker container with test_db, err = %v\n", err)
	}

	cases := []struct {
		name   string
		method string
		target string
		json   string
		input  HandlerAddChat
		want   string
	}{
		{
			name:   "Ok",
			method: http.MethodPost,
			target: "/chats/add",
			json:   `{"name": "chat1", "users": [1, 2]}`,
			input: HandlerAddChat{
				Name:    "",
				Users:   []int{},
				Ctx:     &ctx,
				Conn_db: conn,
			},
			want: "ERROR: duplicate key value violates unique constraint \"chat_name_key\" (SQLSTATE 23505)\n",
		},
		{
			name:   "Ok",
			method: http.MethodPost,
			target: "/chats/add",
			json:   `{"name": "chat1", "users": [1, 228]}`,
			input: HandlerAddChat{
				Name:    "",
				Users:   []int{},
				Ctx:     &ctx,
				Conn_db: conn,
			},
			want: "User with id = 228 does not exist\n",
		},
		{
			name:   "Ok",
			method: http.MethodPost,
			target: "/chats/add",
			json:   `{"name": "chat3", "users": [0, 2]}`,
			input: HandlerAddChat{
				Name:    "",
				Users:   []int{},
				Ctx:     &ctx,
				Conn_db: conn,
			},
			want: "User with id = 0 does not exist\n",
		},
		{
			name:   "Ok",
			method: http.MethodPost,
			target: "/chats/add",
			json:   `{"name": "chat3", "users": [1, 2]}`,
			input: HandlerAddChat{
				Name:    "",
				Users:   []int{},
				Ctx:     &ctx,
				Conn_db: conn,
			},
			want: "Chat id 3\n",
		},
	}

	for _, current_case := range cases {
		request := httptest.NewRequest(current_case.method, current_case.target, bytes.NewReader([]byte(current_case.json)))
		responseRecoder := httptest.NewRecorder()

		current_case.input.ServeHTTP(responseRecoder, request)
		if responseRecoder.Body.String() != current_case.want {
			t.Errorf("received: %v	expected: %v\n", responseRecoder.Body.String(), current_case.want)
		}
	}
}
