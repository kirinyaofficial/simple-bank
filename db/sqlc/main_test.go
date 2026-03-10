package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const dbSource = "postgresql://root:secrete@localhost:5432/simple_bank?sslmode=disable"

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Fatal("Can not create connection pool", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("Cannot ping db:", err)
	}
	testDB = pool
	testQueries = New(pool)

	code := m.Run()

	testDB.Close()
	os.Exit(code)
}
