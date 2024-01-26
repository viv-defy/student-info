package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lpernett/godotenv"
)

var Port string
var rootDir = "/Users/vishnu/Projects/Go/student-info"

func InitConfig() {
	envPath := filepath.Join(rootDir, ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	Port = os.Getenv("PORT")
}
