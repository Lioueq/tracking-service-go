package main

import (
	"log"
	_ "tracking-service-go/docs"
	"tracking-service-go/internal/app"
)

// @title Echo Swagger Example API
// @version 1.0
// @description This is a sample server for an Echo application with Swagger documentation.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	app, err := app.NewApp()
	if err != nil {
		log.Fatalf("Error initializing application: %v", err)
	}

	app.Run(":8080")
}
