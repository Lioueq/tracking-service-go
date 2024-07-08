package routes

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
	"tracking-service-go/internal/controllers"
	"tracking-service-go/internal/service"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	var jwtSecret = service.GetEnv("JWT_SECRET", "secret")
	controllers.InitUserRepository(db, jwtSecret)
	controllers.InitOrderRepository(db)

	// Auth routes
	e.POST("/auth/register", controllers.Register)
	e.POST("/auth/login", controllers.Login)

	// Orders routes
	r := e.Group("/orders")
	r.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(jwtSecret),
		ContextKey: "id",
	}))
	r.GET("", controllers.GetOrders)
	r.PUT("/update", controllers.UpdateOrder)
	r.POST("/create", controllers.CreateOrder)

	// Get all users
	e.GET("/users", controllers.GetUsers)

	// Swagger route
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
