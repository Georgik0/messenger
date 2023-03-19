package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	initservices "messenger/init-services"
	"messenger/interactorsDb"
	"net/http"
	"os"
)

func main() {
	var conn *pgx.Conn
	ctx := context.Background()

	err := interactorsDb.Init_connect_for_test(&conn, &ctx, "myapp_db", "5432")
	if err != nil {
		fmt.Errorf("Please, run docker container with db, err = %v\n", err)
		os.Exit(1)
	}

	initservices.Run(ctx, conn)

	http.ListenAndServe(":9000", nil)
}
