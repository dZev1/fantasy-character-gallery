package config

import (
	"fmt"
	"os"

	"github.com/lpernett/godotenv"
)

func ConnStr() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("error loading .env file: %v", err)
	}
	return os.Getenv("DATABASE_URL"), nil
}