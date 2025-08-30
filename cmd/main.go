package main

import (
	"fmt"
	"log"

	"github.com/dZev1/character-gallery/database"
	"github.com/dZev1/character-gallery/models"
)

func main() {
	database.InitDB()
	defer database.CloseDB()
	char := &models.Character{
		Name: "Juan",
		BodyType: models.TypeA,
		Species: models.Halfling,
		Class: models.Monk,
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
			Face: 1,
			Shirt: 2,
			Pants: 3,
			Shoes: 0,
		},
	}

	fmt.Println(char)

	err := database.DeleteCharacter(char.Name)
	if err != nil  {
		log.Fatal(err)
	}
}