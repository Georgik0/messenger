package main

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"strconv"
)

func CheckUsers(users []int, ctx context.Context, conn *pgx.Conn) error {
	var err error
	var check bool = true
	for _, user := range users {
		err = conn.QueryRow(ctx, "select exists(select 1 from my_user where id = ($1))", user).Scan(&check)
		if check == false {
			err = errors.New("User with id = " + strconv.Itoa(user) + " does not exist")
			return err
		}
	}
	return nil
}

func CheckChats(chats []int, ctx context.Context, conn *pgx.Conn) error {
	var err error
	var check bool = true
	for _, chat_id := range chats {
		err = conn.QueryRow(ctx, "select exists(select 1 from chat where id = ($1))", chat_id).Scan(&check)
		if check == false {
			err = errors.New("Chat with id = " + strconv.Itoa(chat_id) + " does not exist")
			return err
		}
	}
	return nil
}

func CheckUserInChat(users []int, chat_id int, ctx context.Context, conn *pgx.Conn) error {
	var err error
	var check bool
	for _, user_id := range users {
		err = conn.QueryRow(ctx, "SELECT exists(SELECT 1 FROM chat_user WHERE chat_user.user_id = ($1) and chat_user.chat_id = ($2))",
			user_id, chat_id).Scan(&check)
		if err != nil {
			return err
		}
		if check == false {
			err = errors.New("The user with the user_id=" + strconv.Itoa(user_id) +
				" is not a member of chat with chat_id=" + strconv.Itoa(chat_id))
			return err
		}
	}
	return nil
}
