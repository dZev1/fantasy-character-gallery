package main

import (
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
}
