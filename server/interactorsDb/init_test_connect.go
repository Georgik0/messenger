package interactorsDb

import (
	"context"
	"github.com/jackc/pgx/v4"
	"time"
)

func InitConnect(ctx context.Context, host string, port string) (*pgx.Conn, error) {
	timerConnect := time.NewTimer(5 * time.Second)

	conn, err := GetConnect(ctx, host, port)
	for err != nil {
		select {
		case <-timerConnect.C:
			return conn, err
		default:
			conn, err = GetConnect(ctx, host, port)
			continue
		}
	}
	return conn, nil
}
