package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Routes(app *fiber.App, server *Server) {
	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	app.Post("/user", server.createUser)
	app.Get("/user/:email", server.getUser)
	app.Get("/user/id/:id", server.getUserById)
	app.Get("/users", server.getUsers)
	// router.POST("/category", server.createCategory)

}
