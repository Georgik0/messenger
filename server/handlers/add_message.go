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

/* Отправить сообщение в чат от лица пользователя
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"chat": 1, "author": 1, "text": "hi"}' \
  http://localhost:9000/messages/add
*/

type HandlerAddMessage struct {
	Chat_id   int              `json:"chat"`   // ссылка на идентификатор чата, в который было отправлено сообщение
	Author_id int              `json:"author"` // ссылка на идентификатор отправителя сообщения, отношение многие-к-одному
	Text      string           `json:"text"`   // текст отправленного сообщения
	Ctx       *context.Context `json:"-"`
	Conn_db   *pgx.Conn        `json:"-"`
}

func (message *HandlerAddMessage) InitHandler(ctx *context.Context, conn_db *pgx.Conn) {
	message.Ctx = ctx
	message.Conn_db = conn_db
}

func (message *HandlerAddMessage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	err = json.Unmarshal(body, message)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err = interactorsDb.CheckChats([]int{message.Chat_id}, *message.Ctx, message.Conn_db); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err = interactorsDb.CheckUsers([]int{message.Author_id}, *message.Ctx, message.Conn_db); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err = interactorsDb.CheckUserInChat([]int{message.Author_id}, message.Chat_id, *message.Ctx, message.Conn_db); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var id int
	if err = message.Conn_db.QueryRow(context.Background(), "insert into message (chat_id, author_id, text) values ($1, $2, $3) returning id",
		message.Chat_id, message.Author_id, message.Text).Scan(&id); err != nil {
		http.Error(w, err.Error(), 500)
		return
	} else {
		fmt.Fprintln(w, "Message id:", id)
	}
}
