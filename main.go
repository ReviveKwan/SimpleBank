package main

import (
	"database/sql"
	"log"
	"simplebank/api"
	db "simplebank/db/sqlc"
	"simplebank/util"

	_ "github.com/lib/pq"
)

// const (
// 	dbDriver      = "postgres"
// 	dbSource      = "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable"
// 	serverAddress = "localhost:4040"
// )

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(*store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
