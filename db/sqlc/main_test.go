package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/divin3circle/sba/util"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	testQueries *Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Failed to load test config", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.SourceString)
	if err != nil {
		log.Fatal("connection to test db failed!", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
