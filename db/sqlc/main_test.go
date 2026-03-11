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
	var err error

	testDB, err = pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("Can not create connection pool", err)
	}

	if err := testDB.Ping(context.Background()); err != nil {
		log.Fatal("cannot ping database:", err)
	}

	testQueries = New(testDB)

	code := m.Run()

	testDB.Close()
	os.Exit(code)
}
