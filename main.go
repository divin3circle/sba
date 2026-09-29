package main

import (
	"database/sql"
	"log"

	"github.com/divin3circle/sba/api"
	db "github.com/divin3circle/sba/db/sqlc"
	"github.com/divin3circle/sba/util"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load configuration variables", err)
	}
	conn, err := sql.Open(config.DBDriver, config.SourceString)
	if err != nil {
		log.Fatal("Failed to connect to DB", err)
	}

	sbaStore := db.NewStore(conn)
	sbaServer := api.NewServer(sbaStore)

	err = sbaServer.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Failed to start Server", err)
	}
}
