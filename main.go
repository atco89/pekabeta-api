package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"pekabeta/internal/database"
	"pekabeta/internal/fixtures"
	"pekabeta/internal/interfaces/rest/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db := database.InitDb()
	database.RunMigrations(db)

	err = fixtures.Fixtures(db)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	apiGroup := e.Group("/api/1.0")
	api.RegisterHandlers(apiGroup, CreateApi(db))

	e.Logger.Fatal(e.Start(":3000"))
}
