package initservices

import (
	"context"
	"github.com/jackc/pgx/v4"
	"net/http"
)

type handler interface {
	http.Handler
	InitHandler(ctx context.Context, connDB *pgx.Conn)
}
