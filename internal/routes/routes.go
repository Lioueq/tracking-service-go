package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
	"tracking-service-go/internal/config"
	"tracking-service-go/internal/controllers"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	var jwtSecret = config.GetEnv("JWT_SECRET", "secret")
	controllers.InitUserRepository(db, jwtSecret)

	// Auth routes
	e.POST("/auth/register", controllers.Register)
	e.POST("/auth/login", controllers.Login)

	// Get all users
	e.GET("/users", controllers.GetUsers)

	// Swagger route
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
