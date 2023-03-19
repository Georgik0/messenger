package handlers

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type ConnectionDataI interface {
	InitHandler(ctx *context.Context, connDB *pgx.Conn)
}

type ConnectionData struct {
	Ctx    *context.Context `json:"-"`
	connDB *pgx.Conn        `json:"-"`
}

func (data *ConnectionData) InitHandler(ctx *context.Context, connDB *pgx.Conn) {
	data.Ctx = ctx
	data.connDB = connDB
}

func InitHandler(connectionDataI ConnectionDataI) {
	//connectionDataI.InitHandler()
}
