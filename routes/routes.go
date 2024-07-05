package routes

import (
	"github.com/labstack/echo/v4"
	"tracking-service-go/controllers"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/index", controllers.HelloWorld)
}
