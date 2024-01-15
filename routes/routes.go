package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-api/handler"
)

func AddRoutes(app *echo.Echo, h *handler.Handler) {
	api := app.Group("/api")

	items := api.Group("/items")
	items.GET("/:id", h.GetItemByID)
	items.GET("/page", h.GetItems)
	items.GET("/search", h.GetItems)
}
