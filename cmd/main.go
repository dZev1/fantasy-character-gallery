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

	gallery, err := database.NewCharacterGallery(connStr)
	if err != nil {
		panic(err)
	}

	if g, ok := gallery.(*database.PostgresCharacterGallery); ok {
		defer g.Close()
	}

	chars, _ := gallery.GetAll()
	fmt.Println(chars)
}
