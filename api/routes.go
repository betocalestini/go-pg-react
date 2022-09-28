package api

import (
	db "github.com/betocalestini/go-pg-react/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

func NewServerRoutes(store *db.SQLStore) *Server {
	server := &Server{store: store}
	router := fiber.New()

	router.Post("/user", server.createUser)
	router.Get("/user/:email", server.getUser)
	router.Get("/user/id/:id", server.getUserById)
	router.Get("/users", server.getUsers)

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

	server.router = router
	return server
}
