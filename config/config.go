package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	if os.Getenv("VERCEL") != "1" {
		loadError := godotenv.Load()
		if loadError != nil {
			log.Println("⚠️ Warning: No .env file found (using system environment variables instead)")
		}
	}

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	databaseInstance, connectionError := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if connectionError != nil {
		log.Fatal("❌ Failed to connect to database:", connectionError)
	}

	DB = databaseInstance
	fmt.Println("✅ Database connected successfully")
}
