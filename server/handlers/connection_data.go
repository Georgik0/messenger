package handlers

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type ConnectionDataI interface {
	InitHandler(ctx *context.Context, conn_db *pgx.Conn)
}

type ConnectionData struct {
	Ctx     *context.Context `json:"-"`
	Conn_db *pgx.Conn        `json:"-"`
}

func (data *ConnectionData) InitHandler(ctx *context.Context, conn_db *pgx.Conn) {
	data.Ctx = ctx
	data.Conn_db = conn_db
}

func InitHandler(connectionDataI ConnectionDataI) {
	//connectionDataI.InitHandler()
}
