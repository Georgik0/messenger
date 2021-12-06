package handlers

import (
	"bytes"
	"context"
	"github.com/jackc/pgx/v4"
	"messenger/interactorsDb"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"testing"
	"time"
)

func TestHandlerAddUser_ServeHTTP(t *testing.T) {
	cmdStart := "sudo docker-compose up --build -d test_db"
	cmdEnd := "sudo docker-compose stop test_db"
	cmdDeleteData := "sudo rm -rf ../test_data"
	c := exec.Command("/bin/sh", "-c", cmdStart)
	defer exec.Command("/bin/sh", "-c", cmdDeleteData).Run()
	defer exec.Command("/bin/sh", "-c", cmdEnd).Run()
	err := c.Run()
	if err != nil {
		t.Error(err)
	}

	var conn *pgx.Conn
	ctx := context.Background()
	timer_connect := time.NewTimer(10 * time.Second)
	conn, err = interactorsDb.GetConnect(&ctx, "5431")
	for err != nil {
		select {
		case <-timer_connect.C:
			if err == nil {
				t.Error(err)
				return
			}
		default:
			conn, err = interactorsDb.GetConnect(&ctx, "5431")
			continue
		}
	}

	cases := []struct {
		name   string
		method string
		target string
		json   string
		input  HandlerAddUser
		want   string
	}{
		{
			name:   "Ok",
			method: http.MethodPost,
			target: "/users/add",
			json:   `{"username": "user_1"}`,
			input: HandlerAddUser{
				Name:    "",
				Ctx:     &ctx,
				Conn_db: conn,
			},
			want: "User id: 1\n",
		},
		{
			name:   "Ok",
			method: http.MethodPost,
			target: "/users/add",
			json:   `{"username": "user_1"}`,
			input: HandlerAddUser{
				Name:    "",
				Ctx:     &ctx,
				Conn_db: conn,
			},
			want: "User id: 2\n",
		},
	}

	for _, current_case := range cases {
		//t.Run(current_case.name, func(t *testing.T) {
		request := httptest.NewRequest(current_case.method, "/users/add", bytes.NewReader([]byte(current_case.json)))
		responseRecoder := httptest.NewRecorder()

		current_case.input.ServeHTTP(responseRecoder, request)
		if responseRecoder.Body.String() != current_case.want {
			t.Errorf("received: %v	expected: %v\n", responseRecoder.Body.String(), current_case.want)
		}
		//})
	}
}
