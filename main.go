package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"tracking-service-go/internal/config"
	_ "tracking-service-go/internal/docs"
	"tracking-service-go/internal/routes"
)

// @title Echo Swagger Example API
// @version 1.0
// @description This is a sample server for an Echo application with Swagger documentation.
// @host localhost:8080
// @BasePath /
func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	sqlDB, err := db.DB()
	defer sqlDB.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.InitRoutes(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
