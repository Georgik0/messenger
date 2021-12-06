package handlers

import (
	"context"
	"github.com/jackc/pgx/v4"
	"messenger/interactorsDb"
	"net/http"
	"testing"
)

func TestHandlerAddMessage_ServeHTTP(t *testing.T) {
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
		input  HandlerAddMessage
		want   string
	}{
		{
			name:   "Ok",
			method: http.MethodPost,
			target: "/chats/add",
			json:   `{"name": "chat1", "users": [1, 2]}`,
			input: HandlerAddMessage{
				Ctx:     &ctx,
				Conn_db: conn,
			},
			want: "ERROR: duplicate key value violates unique constraint \"chat_name_key\" (SQLSTATE 23505)\n",
		},
	}
	t.Log(cases)
}
