package main

import (
	"log"
	"os"

	"github.com/dZev1/character-gallery/database"
	"github.com/dZev1/character-gallery/models"
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

	kaladin := &models.Character{
		Name: "Kaladin",
		BodyType: models.TypeA,
		Species: models.Aasimar,
		Class: models.Paladin,
		Stats: &models.Stats{
			Strength: 10,
			Dexterity: 10,
			Constitution: 10,
			Intelligence: 10,
			Wisdom: 10,
			Charisma: 10,
		},
		Customization: &models.Customization{
			Hair: 1,
			Face: 2,
			Shirt: 0,
			Pants: 5,
			Shoes: 9,
		},
	}

	database.CreateCharacter(kaladin)
}
