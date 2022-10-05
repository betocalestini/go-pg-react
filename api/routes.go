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

	app.Post("/category", server.createCategory)

	app.Put("/category", server.updateCategory)

	app.Get("/category/id/:id", server.getCategoryById)
	app.Get("/category/deleted", server.getDeletedCategories)
	app.Get("/category/deleted/:id", server.getDeletedCategory)

	app.Post("/categories", server.getCategories)

	app.Delete("/category/id/:id", server.deleteCategory)

	app.Get("/categories/:title", server.getCategoriesByTitle)
	// app.Get("/categories", server.store.GetCategories())

}
