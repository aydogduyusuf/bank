package main

import (
	"database/sql"
	"log"

	"github.com/aydogduyusuf/bank/api"
	db "github.com/aydogduyusuf/bank/db/sqlc"
	"github.com/aydogduyusuf/bank/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect do db", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("can not create server")
	}
	
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}