package main

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/mccune1224/betrayal-api/handler"
	"github.com/mccune1224/betrayal-api/routes"
)

func main() {
	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	app := echo.New()
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	} else {
		// Live environment (aka toggle prod middleware)
		app.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
		app.Use(middleware.CORS())
	}

	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} | ${uri} | ${status} | ${latency_human} | ${error}\n",
	}))

	app.Use(middleware.RemoveTrailingSlash())
	app.Use(middleware.Recover())

	foo := func(c echo.Context) error { return c.JSON(200, echo.Map{"Hello": "World"}) }
	app.GET("/", foo)

	handler := handler.NewHandler(db)
	routes.AddRoutes(app, handler)
	log.Fatal(app.Start(":" + port))
}
