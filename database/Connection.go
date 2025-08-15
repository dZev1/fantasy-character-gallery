package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)
var db *sql.DB

func InitDB(connStr string) error {
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("could not connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("could not verify connection: %v", err)
	}

	fmt.Println("connection to database has been established")
	return nil
}

func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}