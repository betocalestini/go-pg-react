package api

import (
	"database/sql"
	"strconv"

	db "github.com/betocalestini/go-pg-react/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

type createUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (server *Server) createUser(c *fiber.Ctx) error {
	var body createUserRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	arg := db.CreateUserParams{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	}

	if user, err := server.store.CreateUser(c.Context(), arg); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	} else {
		c.Status(fiber.StatusCreated)
		return c.JSON(fiber.Map{
			"user Criado": user})
	}
}

type getUserRequest struct {
	Email string `json:"email" validate:"required"`
}

func (server *Server) getUser(c *fiber.Ctx) error {
	req := getUserRequest{Email: c.Params("email")}

	user, err := server.store.GetUser(c.Context(), req.Email)
	if err != nil {
		if err == sql.ErrNoRows {

			return c.Status(fiber.StatusNotFound).JSON(errorResponse(err))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}
	}
	c.Status(fiber.StatusFound)
	return c.JSON(fiber.Map{
		"user": user})
}

func (server *Server) getUserById(c *fiber.Ctx) error {
	param := c.Params("id")
	paramId, _ := strconv.ParseInt(param, 10, 32)

	user, err := server.store.GetUserById(c.Context(), int32(paramId))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(err)
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}
	}
	c.Status(fiber.StatusFound)
	return c.JSON(fiber.Map{
		"user": user})
}

func (server *Server) getUsers(c *fiber.Ctx) error {

	users, err := server.store.GetUsers(c.Context())
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(err)
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}
	}
	c.Status(fiber.StatusFound)
	return c.JSON(fiber.Map{
		"users": users})
}
