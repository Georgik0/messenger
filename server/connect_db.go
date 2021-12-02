package main

import "github.com/jackc/pgx"

type Db struct {
	conn_db *pgx.Conn
}
