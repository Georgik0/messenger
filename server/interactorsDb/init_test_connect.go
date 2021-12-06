package interactorsDb

import (
	"context"
	"github.com/jackc/pgx/v4"
	"time"
)

func Init_connect_for_test(conn **pgx.Conn, ctx *context.Context) error {
	timer_connect := time.NewTimer(5 * time.Second)
	var err error
	*conn, err = GetConnect(ctx, "5431")
	for err != nil {
		select {
		case <-timer_connect.C:
			return err
		default:
			*conn, err = GetConnect(ctx, "5431")
			continue
		}
	}
	return nil
}
