package api

import (
	db "github.com/betocalestini/go-pg-react/db/sqlc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	store  *db.SQLStore
	router *fiber.App
}

func NewServer(store *db.SQLStore) (*Server, *fiber.App) {
	server := &Server{store: store}
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}${time} | ${white}${pid} | ${white}${latency} |${red}${status} ${red}${error} |${blue}${method} ${white}${host}${white}${path} \n",
		TimeFormat: "02/01/2006 03:04:05",
		TimeZone:   "America/Sao_Paulo",
	}))
	server.router = app
	return server, app
}

func (server *Server) Start(address string) error {
	return server.router.Listen(address)
}

func errorResponse(err error) fiber.Map {
	return fiber.Map{
		"api has error:": err.Error()}
}
