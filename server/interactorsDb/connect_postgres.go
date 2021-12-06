package interactorsDb

import (
	"context"
	"github.com/jackc/pgx/v4"
)

func GetConnect(ctx *context.Context, port string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(*ctx, "postgres://db_user:db_user_pass@127.0.0.1:"+port+"/app_db")
	if err != nil {
		return nil, err
	}
	return conn, err
}
