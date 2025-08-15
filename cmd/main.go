package main

import (
	"fmt"

	"github.com/dZev1/fantasy-character-gallery/models"
)

func main() {
	myChar := &models.Character {
		ID: 10,
		Creator: "Diego",
		Name: "Pepe",
		Species: models.Human,
		Class: models.Monk,
		Stats: &models.Stats{
			Strength: 10,
			Dexterity: 10,
			Intelligence: 10,
			Constitution: 10,
			Wisdom: 10,
			Charisma: 10,
		},
		Customization: &models.Customization{
			Hair: 0,
			Face: 0,
			Shirt: 0,
			Pants: 0,
			Shoes: 0,
		},
	}

	fmt.Println(myChar)
}