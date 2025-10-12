package main

import (
	"log"
	"os"

	"github.com/dZev1/character-gallery/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	connStr := os.Getenv("DATABASE_URL")

	err = database.InitDB(connStr)
	if err != nil {
		panic(err)
	}
	defer database.CloseDB()

}
