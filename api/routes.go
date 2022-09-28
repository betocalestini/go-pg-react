package api

import (
	db "github.com/betocalestini/go-pg-react/db/sqlc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func NewServerRoutes(store *db.SQLStore) *Server {
	server := &Server{store: store}
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}${time} | ${white}${pid} | ${white}${latency} |${red}${status} ${red}${error} | ${blue}${method} ${white}${path} \n",
		TimeFormat: "02/01/2006 03:04:05",
		TimeZone:   "America/Sao_Paulo",
	}))

	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	app.Post("/user", server.createUser)
	app.Get("/user/:email", server.getUser)
	app.Get("/user/id/:id", server.getUserById)
	app.Get("/users", server.getUsers)

	// router.POST("/category", server.createCategory)
	// router.GET("/category/id/:id", server.getCategory)
	// router.GET("/category", server.getCategories)
	// router.DELETE("/category/:id", server.deleteCategory)
	// router.PUT("/category/:id", server.updateCategory)

	// router.POST("/account", server.createAccount)
	// router.GET("/account/id/:id", server.getAccount)
	// router.GET("/account", server.getAccounts)
	// router.GET("/account/graph/:user_id/:type", server.getAccountGraph)
	// router.GET("/account/reports/:user_id/:type", server.getAccountReports)
	// router.DELETE("/account/:id", server.deleteAccount)
	// router.PUT("/account/:id", server.updateAccount)

	// router.POST("/login", server.login)

	server.router = app
	return server
}
