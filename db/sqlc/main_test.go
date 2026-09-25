package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	dbDriver     = "pgx"
	sourceString = "postgresql://postgres:postgres@localhost:5432/sba_db?sslmode=disable"
)

var (
	testQueries *Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, sourceString)
	if err != nil {
		log.Fatal("connection to test db failed!", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
