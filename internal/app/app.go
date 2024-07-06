package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"tracking-service-go/internal/config"
	"tracking-service-go/internal/routes"
)

func Run() {
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
