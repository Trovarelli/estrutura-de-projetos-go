package config

import (
	"log"
	"myapi/internal/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar com o BD: %v", err)
	}
	DB = db

	if err = DB.AutoMigrate(&models.Item{}); err != nil {
		log.Fatalf("Erro durante a mgiration")
	}
}
