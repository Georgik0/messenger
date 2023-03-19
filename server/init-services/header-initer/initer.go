package headeriniter

import (
	"context"
	"github.com/jackc/pgx/v4"
	"net/http"
)

type Provider interface {
	http.Handler
	InitHandler(ctx context.Context, connDB *pgx.Conn)
}
