package api

import (
	db "github.com/betocalestini/go-pg-react/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	store  *db.SQLStore
	router *fiber.App
}

func (server *Server) Start(address string) error {
	return server.router.Listen(address)
}
