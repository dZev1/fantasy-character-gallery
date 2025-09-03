package main

import (
	"fmt"

	"github.com/dZev1/character-gallery/database"
)

func main() {
	err := database.InitDB()
	if err != nil {
		panic(err)
	}
	defer database.CloseDB()

	fmt.Println("KEK")
}
