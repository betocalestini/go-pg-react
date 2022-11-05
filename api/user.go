package api

import (
	"database/sql"

	db "github.com/betocalestini/go-pg-react/db/sqlc"
	"github.com/betocalestini/go-pg-react/util"
	"github.com/gofiber/fiber/v2"
)

type createUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (server *Server) createUser(c *fiber.Ctx) error {
	var body createUserRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}
	if err := util.Validate.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}
	arg := db.CreateUserParams{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	}

	if user, err := server.store.CreateUser(c.Context(), arg); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
	} else {
		c.Status(fiber.StatusCreated)
		return c.JSON(fiber.Map{
			"user created": user})
	}
}

// type getUserRequest struct {
// 	Email string `params:"email" validate:"required,email"`
// }

func (server *Server) getUser(c *fiber.Ctx) error {
	// var param getUserRequest
	// if err := c.ParamsParser(&param); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	// }
	// if err := util.Validate.Struct(param); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	// }

	email := c.Params("email")
	if err := util.Validate.Var(email, "required,email"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	user, err := server.store.GetUser(c.Context(), email)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(errorResponse(err))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
		}
	}
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user})
	return nil
}

type getUserByIdRequest struct {
	ID int32 `params:"id" validate:"required"`
}

func (server *Server) getUserById(c *fiber.Ctx) error {
	var param getUserByIdRequest
	if err := c.ParamsParser(&param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}
	if err := util.Validate.Struct(param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	user, err := server.store.GetUserById(c.Context(), param.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(errorResponse(err))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
		}
	}
	c.Status(fiber.StatusOK)
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
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"users": users})
}
