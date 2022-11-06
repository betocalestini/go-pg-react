package api

import (
	"database/sql"
	"fmt"
	"time"

	db "github.com/betocalestini/go-pg-react/db/sqlc"
	"github.com/betocalestini/go-pg-react/util"
	"github.com/gofiber/fiber/v2"
)

type accountRequest struct {
	ID          int32     `json:"id,string,omitempty" validate:"required"`
	UserID      int32     `json:"user_id,string,omitempty" params:"user_id" validate:"required"`
	CategoryID  int32     `json:"category_id,string,omitempty" validate:"required"`
	Title       string    `json:"title,omitempty" validate:"required"`
	Type        string    `json:"type,omitempty" params:"type" validate:"required"`
	Description string    `json:"description,omitempty" validate:"required"`
	Value       string    `json:"value,omitempty" validate:"required"`
	Date        time.Time `json:"date,omitempty" validate:"required"`
}

func (server *Server) createAccount(c *fiber.Ctx) error {
	var body accountRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	if err := util.Validate.StructExcept(body, "ID", "Date"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	arg := db.CreateAccountParams{
		UserID:      body.UserID,
		CategoryID:  body.CategoryID,
		Title:       body.Title,
		Type:        body.Type,
		Description: body.Description,
		Value:       body.Value,
		Date:        time.Time{},
	}

	if account, err := server.store.CreateAccount(c.Context(), arg); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	} else {
		c.Status(fiber.StatusCreated)
		return c.JSON(fiber.Map{
			"account created": account})
	}
}

func (server *Server) deleteAccount(c *fiber.Ctx) error {
	var req accountRequest
	if err := c.ParamsParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	if err := util.Validate.StructPartial(req, "ID"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	err := server.store.DeleteAccount(c.Context(), req.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"accountId": req.ID,
		"message":   "account deleted"})
}

func (server *Server) getAccount(c *fiber.Ctx) error {
	var req accountRequest
	err := c.ParamsParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}
	if err := util.Validate.StructPartial(req, "ID"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	account, err := server.store.GetAccount(c.Context(), req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(errorResponse(err))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
		}
	}
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"account": account})
}

func (server *Server) getAccountsReports(c *fiber.Ctx) error {
	var body accountRequest
	if err := c.ParamsParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	if err := util.Validate.StructPartial(body, "UserID", "Type"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	arg := db.GetAccountsReportsParams{
		UserID: body.UserID,
		Type:   body.Type,
	}

	sumReports, err := server.store.GetAccountsReports(c.Context(), arg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	fmt.Println(sumReports)
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"sumReports": string(sumReports)})
}

func (server *Server) getAccountsGraph(c *fiber.Ctx) error {
	var body accountRequest
	if err := c.ParamsParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	if err := util.Validate.StructPartial(body, "UserID", "Type"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	arg := db.GetAccountsGraphParams{
		UserID: body.UserID,
		Type:   body.Type,
	}

	accountsGraph, err := server.store.GetAccountsGraph(c.Context(), arg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"accountsGraph": accountsGraph})
}

func (server *Server) getAccounts(c *fiber.Ctx) error {
	var body accountRequest
	if err := c.BodyParser(&body); err != nil {
		fmt.Println(body)
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	if err := util.Validate.StructExcept(body, "ID", "Value"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	arg := db.GetAccountsParams{
		UserID:      body.UserID,
		Type:        body.Type,
		CategoryID:  body.CategoryID,
		Title:       body.Title,
		Description: body.Description,
		Date:        body.Date,
	}

	accounts, err := server.store.GetAccounts(c.Context(), arg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"accounts": accounts})
}

func (server *Server) getAllAccounts(c *fiber.Ctx) error {
	accounts, err := server.store.GetAllAccounts(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"accounts": accounts})
}

func (server *Server) getDeletedAccount(c *fiber.Ctx) error {
	var req categoryRequest
	err := c.ParamsParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}
	if err := util.Validate.StructPartial(req, "ID"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	category, err := server.store.GetDeletedAccount(c.Context(), req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(errorResponse(err))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}
	}
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"category": category})
}

func (server *Server) getDeletedAccounts(c *fiber.Ctx) error {

	accounts, err := server.store.GetDeletedAccounts(c.Context())
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(errorResponse(err))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}
	}
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"accounts": accounts})
}

func (server *Server) updateAccount(c *fiber.Ctx) error {
	var body accountRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	if err := util.Validate.StructExcept(body, "UserID", "CategoryID", "Date"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}
	arg := db.UpdateAccountParams{
		ID:          body.ID,
		Title:       body.Title,
		Description: body.Description,
		Value:       body.Value,
	}

	if account, err := server.store.UpdateAccount(c.Context(), arg); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	} else {
		c.Status(fiber.StatusCreated)
		return c.JSON(fiber.Map{
			"account upated": account})
	}
}
