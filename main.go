package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kirinyaofficial/simple-bank/api"
	db "github.com/kirinyaofficial/simple-bank/db/sqlc"
)

const (
	dbSource      = "postgresql://root:secrete@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("Can not create connection pool", err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		log.Fatal("cannot ping database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewSerevr(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
