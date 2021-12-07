package interactorsDb

import (
	"context"
	"github.com/jackc/pgx/v4"
	"time"
)

func Init_connect_for_test(conn **pgx.Conn, ctx *context.Context, host string, port string) error {
	timer_connect := time.NewTimer(5 * time.Second)
	var err error
	*conn, err = GetConnect(ctx, host, port)
	for err != nil {
		select {
		case <-timer_connect.C:
			return err
		default:
			*conn, err = GetConnect(ctx, host, port)
			continue
		}
	}
	return nil
}
