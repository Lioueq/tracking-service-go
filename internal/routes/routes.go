package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
	"tracking-service-go/internal/controllers"
	"tracking-service-go/internal/service"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	var jwtSecret = service.GetEnv("JWT_SECRET", "secret")
	controllers.InitUserRepository(db, jwtSecret)

	// Auth routes
	e.POST("/auth/register", controllers.Register)
	e.POST("/auth/login", controllers.Login)

	// Get all users
	e.GET("/users", controllers.GetUsers)

	// Swagger route
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
