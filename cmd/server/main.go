package main

import (
	"log"
	_ "tracking-service-go/docs"
	"tracking-service-go/internal/app"
)

// @title Tracking service go
// @version 1.0
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
