package database

import (
	"fmt"
	"log"
	"os"

	"github.com/dZev1/character-gallery/models/characters"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewCharacterGallery(connStr string) (characters.CharacterGallery, error) {
	var err error
	db, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		return nil, fmt.Errorf("could not establish connection to database: %v", err)
	}

	log.Println("Database connection established")

	schema, err := os.ReadFile("./schema.sql")
	if err != nil {
		return nil, fmt.Errorf("could not load schema: %v", err)
	}

	db.MustExec(string(schema))

	return &PostgresCharacterGallery{
		db: db,
	}, nil
}

func (cg *PostgresCharacterGallery) Close() {
	log.Println("Database connection terminated")
	err := cg.db.Close()
	if err != nil {
		log.Printf("error closing database connection: %v\n", err)
	}
}
