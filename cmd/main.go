package main

import (
	"fmt"

	"github.com/dZev1/fantasy-character-gallery/config"
	"github.com/dZev1/fantasy-character-gallery/database"
)

func main() {
	connStr, err := config.ConnStr()
	if err != nil {
		panic(err)
	}
	err = database.InitDB(connStr)
	if err != nil {
		panic(err)
	}
	defer database.CloseDB()

	// kaladin := &models.Character{
	// 	ID:       34,
	// 	Creator:  "Diego",
	// 	Name:     "Kaladin",
	// 	BodyType: models.TypeA,
	// 	Species:  models.Human,
	// 	Class:    models.Paladin,
	// 	Stats: &models.Stats{
	// 		Strength:     8,
	// 		Dexterity:    8,
	// 		Constitution: 8,
	// 		Intelligence: 8,
	// 		Wisdom:       8,
	// 		Charisma:     8,
	// 	},
	// 	Customization: &models.Customization{
	// 		Hair:  1,
	// 		Face:  1,
	// 		Shirt: 1,
	// 		Pants: 1,
	// 		Shoes: 1,
	// 	},
	// }

	err = database.RegisterUser("dezeta1", "3#jASDj", "", "")
	if err != nil {
		fmt.Println(err)
	}
}
