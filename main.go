package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kirinyaofficial/simple-bank/api"
	db "github.com/kirinyaofficial/simple-bank/db/sqlc"
	"github.com/kirinyaofficial/simple-bank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := pgxpool.New(context.Background(), config.DBSoruce)
	if err != nil {
		log.Fatal("Can not create connection pool", err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		log.Fatal("cannot ping database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewSerevr(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
