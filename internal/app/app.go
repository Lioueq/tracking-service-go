package app

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"tracking-service-go/internal/config"
	"tracking-service-go/internal/routes"
)

type App struct {
	echo *echo.Echo
	db   *gorm.DB
}

func NewApp() (*App, error) {
	cfg, err := config.InitConfig()
	if err != nil {
		return nil, err
	}

	db, err := config.InitDB(cfg)
	if err != nil {
		return nil, err
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.InitRoutes(e, db)

	return &App{
		echo: e,
		db:   db,
	}, nil
}

func (a *App) Run(address string) {
	sqlDB, err := a.db.DB()
	if err != nil {
		log.Fatalf("Error getting database connection: %v", err)
	}
	defer sqlDB.Close()

	a.echo.Logger.Fatal(a.echo.Start(address))
}
