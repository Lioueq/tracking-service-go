package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"tracking-service-go/routes"
)

func main() {
	//cfg, err := config.InitConfig()
	//if err != nil {
	//	fmt.Println(err)
	//}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.InitRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
