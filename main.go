package main

import (
	"log"

	"github.com/betocalestini/go-pg-react/api"
	db "github.com/betocalestini/go-pg-react/db/sqlc"
)

func main() {
	serverAddress, conn := api.Connect()

	store := db.NewStore(conn)
	server := api.NewServerRoutes(store)

	err := server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start api: ", err)
	}
}
