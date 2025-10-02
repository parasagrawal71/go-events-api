package config

import (
	"fmt"
	"go-events-api/cmd/api/models"
	"go-events-api/internal/env"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	SSLMode    string
}

func LoadDbConfigFromEnv() DBConfig {
	return DBConfig{
		DBHost:     env.GetEnvString("DB_HOST", "localhost"),
		DBUser:     env.GetEnvString("DB_USER", "postgres"),
		DBPassword: env.GetEnvString("DB_PASSWORD", ""),
		DBName:     env.GetEnvString("DB_NAME", "go_events_api"),
		DBPort:     env.GetEnvString("DB_PORT", "5432"),
		SSLMode:    env.GetEnvString("SSL_MODE", "disable"),
	}
}

var DB *gorm.DB

func ConnectDatabase() {
	dbConfig := LoadDbConfigFromEnv()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbConfig.DBHost,
		dbConfig.DBUser,
		dbConfig.DBPassword,
		dbConfig.DBName,
		fmt.Sprint(dbConfig.DBPort),
		dbConfig.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// Auto-migrate all models
	db.AutoMigrate(&models.Event{})
	db.AutoMigrate(&models.User{})

	DB = db
	log.Println("Database connected!")
}
