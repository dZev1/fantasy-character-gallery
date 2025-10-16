package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dZev1/character-gallery/database"
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

	chars, _ := database.GetCharacters()
	fmt.Println(chars)
}
