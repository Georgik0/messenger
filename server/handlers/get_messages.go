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

/* Получить список сообщений в конкретном чате
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"chat": 1}' \
  http://localhost:9000/messages/get
*/

type HandlerGetChatMessages struct {
	Chat    int              `json:"chat"`
	Ctx     *context.Context `json:"-"`
	conn_db *pgx.Conn        `json:"-"`
}

func (chat_messages *HandlerGetChatMessages) InitHandler(ctx *context.Context, conn_db *pgx.Conn) {
	chat_messages.Ctx = ctx
	chat_messages.conn_db = conn_db
}

func (chat_messages *HandlerGetChatMessages) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	err = json.Unmarshal(body, chat_messages)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err = interactorsDb.CheckChats([]int{chat_messages.Chat}, context.Background(), chat_messages.conn_db); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if row, err := chat_messages.conn_db.Query(context.Background(), "SELECT text FROM message WHERE chat_id = ($1)", chat_messages.Chat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	} else {
		defer row.Close()
		viewMessages(&row, w)
	}
}

func viewMessages(row *pgx.Rows, w http.ResponseWriter) {
	var text string
	if (*row).Next() == false {
		fmt.Fprintf(w, "There are no messages in the chat\n")
	} else {
		fmt.Fprintf(w, "The chat's messages:\n")
		(*row).Scan(&text)
		fmt.Fprintln(w, text)
	}
	for (*row).Next() {
		(*row).Scan(&text)
		fmt.Fprintln(w, text)
	}
}
