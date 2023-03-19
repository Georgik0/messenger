package handlerbuilder

import (
	"context"
	"github.com/jackc/pgx/v4"
	headeriniter "messenger/init-services/header-initer"
	"net/http"
)

func New(ctx context.Context, connDB *pgx.Conn) *Builder {
	return &Builder{
		ctx:    ctx,
		connDB: connDB,
	}
}

type Builder struct {
	ctx    context.Context
	connDB *pgx.Conn
}

func (b *Builder) InitHandler(h headeriniter.Provider, pattern string) {
	h.InitHandler(b.ctx, b.connDB)
	http.Handle(pattern, h)
}
