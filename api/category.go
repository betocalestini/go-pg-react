package api

import (
	"database/sql"

	db "github.com/betocalestini/go-pg-react/db/sqlc"
	"github.com/betocalestini/go-pg-react/util"
	"github.com/gofiber/fiber/v2"
)

type categoryRequest struct {
	ID          int32  `json:"id,string,omitempty" validate:"required"`
	UserID      int32  `json:"user_id,string,omitempty" validate:"required"`
	Type        string `json:"type,omitempty" validate:"required"`
	Title       string `json:"title,omitempty" validate:"required"`
	Description string `json:"description,omitempty" validate:"required"`
}

func (server *Server) createCategory(c *fiber.Ctx) error {
	var body categoryRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	if err := util.Validate.StructExcept(body, "ID"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	arg := db.CreateCategoryParams{
		UserID:      body.UserID,
		Title:       body.Title,
		Type:        body.Type,
		Description: body.Description,
	}

	if category, err := server.store.CreateCategory(c.Context(), arg); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	} else {
		c.Status(fiber.StatusCreated)
		return c.JSON(fiber.Map{
			"category created": category})
	}
}

func (server *Server) deleteCategory(c *fiber.Ctx) error {
	var req categoryRequest
	if err := c.ParamsParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	if err := util.Validate.StructPartial(req, "ID"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	err := server.store.DeleteCategory(c.Context(), req.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"categoryId": req.ID,
		"message":    "category deleted"})
}

func (server *Server) getCategoryById(c *fiber.Ctx) error {
	var req categoryRequest
	err := c.ParamsParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}
	if err := util.Validate.StructPartial(req, "ID"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	category, err := server.store.GetCategoryById(c.Context(), req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(errorResponse(err))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
		}
	}
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"category": category})
}

func (server *Server) getCategoriesByTitle(c *fiber.Ctx) error {
	var req categoryRequest
	if err := c.ParamsParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	if err := util.Validate.StructPartial(req, "Title"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}
	categories, err := server.store.GetCategoriesByTitle(c.Context(), req.Title)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(errorResponse(err))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}
	}
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"categories by title": categories})
}

func (server *Server) getDeletedCategory(c *fiber.Ctx) error {
	var req categoryRequest
	err := c.ParamsParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}
	if err := util.Validate.StructPartial(req, "ID"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	category, err := server.store.GetDeletedCategory(c.Context(), req.ID)
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

func (server *Server) getCategories(c *fiber.Ctx) error {
	var body categoryRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	if err := util.Validate.StructExcept(body, "ID"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	arg := db.GetCategoriesParams{
		UserID:      body.UserID,
		Title:       body.Title,
		Type:        body.Type,
		Description: body.Description,
	}

	categories, err := server.store.GetCategories(c.Context(), arg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"categories": categories})
}

func (server *Server) getAllCategories(c *fiber.Ctx) error {
	categories, err := server.store.GetAllCategories(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"categories": categories})
}

func (server *Server) getDeletedCategories(c *fiber.Ctx) error {

	categories, err := server.store.GetDeletedCategories(c.Context())
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(errorResponse(err))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}
	}
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"categories": categories})
}

func (server *Server) updateCategory(c *fiber.Ctx) error {
	var body categoryRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	if err := util.Validate.StructExcept(body, "UserID"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}
	arg := db.UpdateCategoryParams{
		ID:          body.ID,
		Title:       body.Title,
		Type:        body.Type,
		Description: body.Description,
	}

	if category, err := server.store.UpdateCategory(c.Context(), arg); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	} else {
		c.Status(fiber.StatusCreated)
		return c.JSON(fiber.Map{
			"category upated": category})
	}
}
