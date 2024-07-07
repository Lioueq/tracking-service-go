package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"tracking-service-go/internal/models"
	"tracking-service-go/internal/service"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

func InitConfig() (Config, error) {
	var config Config
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Initialization error: %v", err)
		return config, err
	}
	config.DBHost = "localhost"
	config.DBPort = "5432"
	config.DBUser = service.GetEnv("DB_USER", "")
	config.DBPassword = service.GetEnv("DB_PASSWORD", "")
	config.DBName = service.GetEnv("DB_NAME", "")
	return config, err
}

func InitDB(cfg Config) (*gorm.DB, error) {
	var dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connection to database: %v", err)
		return nil, err
	}
	db.AutoMigrate(&models.User{}, &models.Order{})
	return db, nil
}
