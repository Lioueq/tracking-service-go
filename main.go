package main

import (
	_ "tracking-service-go/docs"
	"tracking-service-go/internal/app"
)

// @title Echo Swagger Example API
// @version 1.0
// @description This is a sample server for an Echo application with Swagger documentation.
// @host localhost:8080
// @BasePath /
func main() {
	app.Run()
}
