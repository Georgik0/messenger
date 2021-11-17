package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"io/ioutil"
	"net/http"
	"strconv"
)

/* Добавить нового пользователя
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"username": "user_1"}' \
  http://localhost:9000/users/add
*/

type User struct {
	Username string `json:"username"`
}

func Add_users() {
	http.HandleFunc("/users/add", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		p := &User{}
		err = json.Unmarshal(body, p)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		conn, err := pgx.Connect(context.Background(), "postgres://db_user:db_user_pass@myapp_db:5432/app_db")
		defer conn.Close(context.Background())
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		var id int = 0
		err = conn.QueryRow(context.Background(), "insert into my_user (username) values ($1) returning id", p.Username).Scan(&id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		} else {
			fmt.Fprintln(w, "Id созданного пльзователя:", id)
		}
	})
}

/*------------------------------------------------------------------------*/

/* Создать новый чат между пользователями
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name": "chat_1", "users": [1, 2]}' \
  http://localhost:9000/chats/add
*/

type Chat struct {
	Name       string   `json:"name"`
	Users      []int 	`json:"users"`
}

func fillChatUser(users []int, chat_id int, ctx context.Context, conn *pgx.Conn) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, user_id := range users {
		_, err = tx.Exec(ctx,"insert into chat_user (user_id, chat_id) values ($1, $2)", user_id, chat_id)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func Create_chat() {
	http.HandleFunc("/chats/add", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		p := &Chat{}
		err = json.Unmarshal(body, p)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		conn, err := pgx.Connect(context.Background(), "postgres://db_user:db_user_pass@myapp_db:5432/app_db")
		defer conn.Close(context.Background())
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if err = CheckUsers(p.Users, context.Background(), conn); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		var id int
		err = conn.QueryRow(context.Background(), "insert into chat (name) values ($1) returning id", p.Name).Scan(&id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		} else {
			fmt.Fprintln(w, "Id созданного чата", id)
		}
		if err = fillChatUser(p.Users, id, context.Background(), conn); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	})
}

/*------------------------------------------------------------------------*/

/* Отправить сообщение в чат от лица пользователя
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"chat": 1, "author": 1, "text": "hi"}' \
  http://localhost:9000/messages/add
*/

type Message struct {
	Chat_id    int	  `json:"chat"`   // ссылка на идентификатор чата, в который было отправлено сообщение
	Author_id  int	  `json:"author"` // ссылка на идентификатор отправителя сообщения, отношение многие-к-одному
	Text       string `json:"text"`   // текст отправленного сообщения
}

func Message_add() {
	http.HandleFunc("/messages/add", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		p := &Message{}
		err = json.Unmarshal(body, p)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		conn, err := pgx.Connect(context.Background(), "postgres://db_user:db_user_pass@myapp_db:5432/app_db")
		defer conn.Close(context.Background())
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if err = CheckChats([]int{p.Chat_id}, context.Background(), conn); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		if err = CheckUsers([]int{p.Author_id}, context.Background(), conn); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		if err = CheckUserInChat([]int{p.Author_id}, p.Chat_id, context.Background(), conn); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		var id int
		if err = conn.QueryRow(context.Background(), "insert into message (chat_id, author_id, text) values ($1, $2, $3) returning id",
			p.Chat_id, p.Author_id, p.Text).Scan(&id); err != nil {
			http.Error(w, err.Error(), 500)
			return
		} else {
			fmt.Fprintln(w, "Id созданного сообщения", id)
		}
	})
}

/*------------------------------------------------------------------------*/

/* Получить список чатов конкретного пользователя
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"user": 1}' \
  http://localhost:9000/chats/get
*/

type User_chats struct {
	User int `json:"user"`
}

func Get_chats() {
	http.HandleFunc("/chats/get", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		p := &User_chats{}
		err = json.Unmarshal(body, p)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		conn, err := pgx.Connect(context.Background(), "postgres://db_user:db_user_pass@myapp_db:5432/app_db")
		defer conn.Close(context.Background())
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if err = CheckUsers([]int{p.User}, context.Background(), conn); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if row, err := conn.Query(context.Background(), "SELECT cu.chat_id FROM chat_user cu LEFT JOIN chat c ON c.id=cu.chat_id WHERE cu.user_id = ($1) ORDER BY c.created_at", p.User); err != nil {
			http.Error(w, err.Error(), 500)
			return
		} else {
			defer row.Close()
			fmt.Fprintf(w, "The user's chat_id:")
			var chat_id int
			for row.Next() {
				row.Scan(&chat_id)
				fmt.Fprintf(w, " " + strconv.Itoa(chat_id))
			}
			fmt.Fprintln(w)
		}
	})
}

/*------------------------------------------------------------------------*/

/* Получить список сообщений в конкретном чате
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"chat": 1}' \
  http://localhost:9000/messages/get
*/

type Chat_messages struct {
	Chat int `json:"chat"`
}

func viewMessages(row *pgx.Rows, w http.ResponseWriter) {
	var text string
	if ((*row).Next() == false) {
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

func Get_messages() {
	http.HandleFunc("/messages/get", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		p := &Chat_messages{}
		err = json.Unmarshal(body, p)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		conn, err := pgx.Connect(context.Background(), "postgres://db_user:db_user_pass@myapp_db:5432/app_db")
		defer conn.Close(context.Background())
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if err = CheckChats([]int{p.Chat}, context.Background(), conn); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if row, err := conn.Query(context.Background(), "SELECT text FROM message WHERE chat_id = ($1)", p.Chat); err != nil {
			http.Error(w, err.Error(), 500)
			return
		} else {
			defer row.Close()
			viewMessages(&row, w)
		}
	})
}

func main() {
	Add_users()
	Create_chat()
	Message_add()
	Get_chats()
	Get_messages()
	http.ListenAndServe(":9000", nil)
}
