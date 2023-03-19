package handlers

import (
	"context"
	"github.com/jackc/pgx/v4"
	"messenger/interactorsDb"
	"net/http"
	"testing"
)

func TestHandlerAddUser_ServeHTTP(t *testing.T) {
	/*cmdStart := "sudo docker-compose up --build -d test_db"
	cmdEnd := "sudo docker-compose stop test_db"
	cmdDeleteData := "sudo rm -rf ../test_data"
	c := exec.Command("/bin/sh", "-c", cmdStart)
	defer exec.Command("/bin/sh", "-c", cmdDeleteData).Run()
	defer exec.Command("/bin/sh", "-c", cmdEnd).Run()
	err := c.Run()
	if err != nil {
		t.Error(err)
	}*/

	var conn *pgx.Conn
	ctx := context.Background()
	if err := interactorsDb.Init_connect_for_test(&conn, &ctx, "127.0.0.1", "5431"); err != nil {
		t.Errorf("Please, run docker container with test_db, err = %v\n", err)
	}

	handler_input := &HandlerAddUser{
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
			target: "/users/add",
			json:   `{"username": "user_1"}`,
			input:  handler_input,
			want:   "User id: 4\n",
		},
		{
			name:   "Ok",
			method: http.MethodPost,
			target: "/users/add",
			json:   `{"username": "user_1"}`,
			input:  handler_input,
			want:   "ERROR: duplicate key value violates unique constraint \"my_user_username_key\" (SQLSTATE 23505)\n",
		},
	}

	CheckTestsRange(HndlInput(cases), t)
}
