package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"io/ioutil"
	"messenger/interactorsDb"
	"net/http"
)

/* Создать новый чат между пользователями
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name": "chat_1", "users": [1, 2]}' \
  http://localhost:9000/chats/add
*/

type HandlerAddChat struct {
	Name    string           `json:"name"`
	Users   []int            `json:"users"`
	Ctx     *context.Context `json:"-"`
	Conn_db *pgx.Conn        `json:"-"`
}

func (chat *HandlerAddChat) InitHandler(ctx *context.Context, conn_db *pgx.Conn) {
	chat.Ctx = ctx
	chat.Conn_db = conn_db
}

func (chat *HandlerAddChat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	err = json.Unmarshal(body, chat)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if chat.Users == nil || chat.Ctx == nil || chat.Conn_db == nil {
		fmt.Printf("Nil\n")
	}

	if err = interactorsDb.CheckUsers(chat.Users, *chat.Ctx, chat.Conn_db); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var id int
	err = chat.Conn_db.QueryRow(*chat.Ctx, "insert into chat (name) values ($1) returning id", chat.Name).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	} else {
		fmt.Fprintln(w, "Chat id", id)
	}
	if err = interactorsDb.FillChatUser(chat.Users, id, chat.Ctx, chat.Conn_db); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
