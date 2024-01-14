package routes

import "github.com/labstack/echo/v4"

func AddRoutes(app *echo.Echo) {
	api := app.Group("/api")

	item := api.Group("/item")
	item.GET("", func(c echo.Context) error {
		return c.JSON(200, echo.Map{"item": "hit"})
	})
}
