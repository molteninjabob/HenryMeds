package config

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/joho/godotenv"
)

var (
	Port       string
	DBHost     string
	DBName     string
	DBUser     string
	DBPassword string
)

func init() {
	cmd := exec.Command("go", "env", "GOMOD")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	path := filepath.Join(filepath.Dir(string(output)), ".env")
	if err := godotenv.Load(path); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	Port = os.Getenv("PORT")
	DBHost = os.Getenv("DB_HOST")
	DBName = os.Getenv("DB_NAME")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
}
