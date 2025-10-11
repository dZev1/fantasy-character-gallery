package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDB(connStr string) error {
	var err error
	db, err = sqlx.Connect("pgx", connStr)
	if err != nil {
		return fmt.Errorf("could not establish connection to database: %v", err)
	}

	log.Println("Database connection established")

	schema, err := os.ReadFile("./schema.sql")
	if err != nil {
		return fmt.Errorf("could not load schema: %v", err)
	}

	db.MustExec(string(schema))

	return nil
}

func CloseDB() {
	log.Println("Database connection terminated")
	err := db.Close()
	if err != nil {
		panic(err)
	}
}
