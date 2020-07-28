package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"pekabeta/internal/database"
	"pekabeta/internal/interfaces/rest/api"
)

func main() {
	db := database.InitDb()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	apiGroup := e.Group("/api/1.0")
	api.RegisterHandlers(apiGroup, CreateApi(db))

	e.Logger.Fatal(e.Start(":3000"))
}
