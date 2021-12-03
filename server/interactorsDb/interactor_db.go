package interactorsDb

import (
	"context"
	"github.com/jackc/pgx/v4"
)

func FillChatUser(users []int, chat_id int, ctx *context.Context, conn *pgx.Conn) error {
	tx, err := conn.Begin(*ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(*ctx)

	for _, user_id := range users {
		_, err = tx.Exec(*ctx, "insert into chat_user (user_id, chat_id) values ($1, $2)", user_id, chat_id)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(*ctx)
	if err != nil {
		return err
	}

	return nil
}
