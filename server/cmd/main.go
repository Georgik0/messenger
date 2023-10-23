package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	initservices "messenger/init-services"
	"messenger/interactorsDb"
)

func main() {
	ctx := context.Background()

	conn, err := interactorsDb.InitConnect(ctx, "myapp_db", "5432")
	if err != nil {
		fmt.Errorf("Please, run docker container with db, err = %v\n", err)
		os.Exit(1)
	}

	initservices.Run(ctx, conn)

	http.ListenAndServe(":9000", nil)
}
