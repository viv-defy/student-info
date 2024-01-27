package config

import (
	"fmt"
	"os"
	"path/filepath"
	"student-info/models"

	"github.com/lpernett/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Port string
var rootDir = "/Users/vishnu/Projects/Go/student-info"
var Db *gorm.DB

func Init() {
	envPath := filepath.Join(rootDir, ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	Port = os.Getenv("PORT")

	dsn := os.Getenv("DB_URI")
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database")
	}
	Db.AutoMigrate(&models.Student{})
}
