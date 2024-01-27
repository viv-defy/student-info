package config

import (
	"log"
	"os"
	"path/filepath"
	"student-info/models"

	"github.com/lpernett/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Port string
var RootDir string
var Db *gorm.DB

func Init() {
	RootDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting the current directory")
	}

	envPath := filepath.Join(RootDir, ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Port = os.Getenv("PORT")

	dsn := os.Getenv("DB_URI")
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	Db.AutoMigrate(&models.Student{})
}
