package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"io/ioutil"
	"net/http"
)

/* Добавить нового пользователя
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"username": "user_1"}' \
  http://localhost:9000/users/add
*/

type HandlerAddUser struct {
	Name    string          `json:"username,int"`
	Ctx     context.Context `json:"-"`
	conn_db *pgx.Conn       `json:"-"`
}

func (user *HandlerAddUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	err = json.Unmarshal(body, user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var id int = 0
	err = user.conn_db.QueryRow(user.Ctx, "insert into my_user (username) values ($1) returning id", user.Name).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	} else {
		fmt.Fprintln(w, "User id:", id)
	}
}
