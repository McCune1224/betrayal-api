package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	// DefaultPerPage is the default number of items to return per page
	DefaultLimit = 10
	// DefaultLimit is the maximum number of items to return per page
	MaxLimit = 100
)

func (h *Handler) GetItemByName(c echo.Context) error {
	name := c.QueryParam("name")

	if name == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "name is required"})
	}
	item, err := h.models.ItemModel.GetItemByName(name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "item not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, item)
}

func (h *Handler) GetItemByID(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "id must be an integer"})
	}

	if id == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "id is required"})
	}

	item, err := h.models.ItemModel.GetItem(intID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "item not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, item)
}

func (h *Handler) GetItems(c echo.Context) error {
	offsetQuery := c.QueryParam("offset")
	limitQuery := c.QueryParam("limit")
	var offset, limit int
	var err error
	if offsetQuery == "" {
		offset = 1
	} else {
		offset, err = strconv.Atoi(offsetQuery)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "page must be an integer"})
		}
	}
	if limitQuery == "" {
		limit = DefaultLimit
	} else {
		limit, err = strconv.Atoi(limitQuery)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "limit must be an integer"})
		}
	}

	if limit > MaxLimit {
		limit = MaxLimit
	}
	items, err := h.models.ItemModel.PageItems(offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	if len(items) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "no items found"})
	}

	return c.JSON(http.StatusOK, items)
}
