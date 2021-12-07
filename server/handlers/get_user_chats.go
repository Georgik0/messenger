package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"io/ioutil"
	"messenger/interactorsDb"
	"net/http"
	"strconv"
)

/* Получить список чатов конкретного пользователя
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"user": 1}' \
  http://localhost:9000/chats/get
*/

type HandlerGetUserChats struct {
	User    int              `json:"user"`
	Ctx     *context.Context `json:"-"`
	Conn_db *pgx.Conn        `json:"-"`
}

func (user_chats *HandlerGetUserChats) InitHandler(ctx *context.Context, conn_db *pgx.Conn) {
	user_chats.Ctx = ctx
	user_chats.Conn_db = conn_db
}

func (user_chats *HandlerGetUserChats) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	err = json.Unmarshal(body, user_chats)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err = interactorsDb.CheckUsers([]int{user_chats.User}, context.Background(), user_chats.Conn_db); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if row, err := user_chats.Conn_db.Query(context.Background(), "SELECT cu.chat_id FROM chat_user cu LEFT JOIN chat c ON c.id=cu.chat_id WHERE cu.user_id = ($1) ORDER BY c.created_at", user_chats.User); err != nil {
		http.Error(w, err.Error(), 500)
		return
	} else {
		defer row.Close()
		fmt.Fprintf(w, "The user's chat_id:")
		var chat_id int
		for row.Next() {
			row.Scan(&chat_id)
			fmt.Fprintf(w, " "+strconv.Itoa(chat_id))
		}
		fmt.Fprintln(w)
	}
}
