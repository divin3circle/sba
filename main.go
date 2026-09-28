package main

import (
	"database/sql"
	"log"

	"github.com/divin3circle/sba/api"
	db "github.com/divin3circle/sba/db/sqlc"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	dbDriver      = "pgx"
	sourceString  = "postgresql://postgres:postgres@localhost:5432/sba_db?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, sourceString)
	if err != nil {
		log.Fatal("Failed to connect to DB", err)
	}

	sbaStore := db.NewStore(conn)
	sbaServer := api.NewServer(sbaStore)

	err = sbaServer.Start(serverAddress)
	if err != nil {
		log.Fatal("Failed to start Server", err)
	}
}
