package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/glebarez/sqlite"
)

var db *sqlx.DB

func InitDB() error {
	var err error
	db, err = sqlx.Open("sqlite", "./database/mydb.sqlite")
	if err != nil {
		return fmt.Errorf("could not establish connection to database: %v", err)
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return fmt.Errorf("could not enable foreign keys: %v", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("could not ping database: %v", err)
	}

	log.Println("Database connection established")
	return nil
}

func CloseDB() {
	log.Println("Database connection terminated")
	db.Close()
}
