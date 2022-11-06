package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Routes(app *fiber.App, server *Server) {
	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	//user routes
	app.Post("/user", server.createUser)

	app.Get("/user/:email", server.getUser)
	app.Get("/user/id/:id", server.getUserById)
	app.Get("/users", server.getUsers)

	//category routes
	app.Post("/category", server.createCategory)
	app.Post("/categories", server.getCategories)

	app.Get("/category/:id", server.getCategoryById)
	app.Get("/category/deleted/:id", server.getDeletedCategory)
	app.Get("/categories/all", server.getAllCategories)
	app.Get("/categories/deleted", server.getDeletedCategories)
	app.Get("/categories/:title", server.getCategoriesByTitle)

	app.Put("/category", server.updateCategory)

	app.Delete("/category/:id", server.deleteCategory)

	//account routes
	app.Post("/account", server.createAccount)
	app.Post("/accounts", server.getAccounts)

	app.Get("/accounts/reports/:user_id/:type", server.getAccountsReports)
	app.Get("/accounts/graph/:user_id/:type", server.getAccountsGraph)

	app.Get("/accounts", server.getAllAccounts)
	app.Get("/account/:id", server.getAccount)
	app.Get("/account/deleted/:id", server.getDeletedAccount)
	app.Get("/accounts/deleted/", server.getDeletedAccounts)

	app.Put("/account", server.updateAccount)

	app.Delete("/account/:id", server.deleteAccount)
}
