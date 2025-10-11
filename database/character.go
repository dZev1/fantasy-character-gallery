package database

import (
	"fmt"

	"github.com/dZev1/character-gallery/models"
	"github.com/jmoiron/sqlx"
)

func InsertCharacter(character *models.Character) error {
	query := `
		INSERT INTO characters (name, body_type, species, class)
		VALUES (:name, :body_type, :species, :class) RETURNING id
	`

	stmt, err := db.PrepareNamed(query)
	if err != nil {
		return fmt.Errorf("could not prepare query: %v", err)
	}
	defer func(stmt *sqlx.NamedStmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	err = stmt.Get(&character.ID, character)
	if err != nil {
		return fmt.Errorf("could not insert character: %v", err)
	}

	character.Stats.ID = character.ID
	character.Customization.ID = character.ID

	query = `
		INSERT INTO stats (id, strength, dexterity, constitution, intelligence, wisdom, charisma)
		VALUES(:id, :strength, :dexterity, :constitution, :intelligence, :wisdom, :charisma)
	`
	_, err = db.NamedExec(query, character.Stats)
	if err != nil {
		return fmt.Errorf("could not insert stats: %v", err)
	}

	query = `
		INSERT INTO customizations (id, hair, face, shirt, pants, shoes)
		VALUES(:id, :hair, :face, :shirt, :pants, :shoes)
	`
	_, err = db.NamedExec(query, character.Customization)
	if err != nil {
		return fmt.Errorf("could not insert customization: %v", err)
	}

	return nil
}

func GetCharacterByID(id models.ID) (*models.Character, error) {
	character := &models.Character{}

	query := `
		SELECT * FROM characters
		WHERE id=$1
	`

	err := db.Get(character, query, id)
	if err != nil {
		return nil, fmt.Errorf("could not get character: %v", err)
	}

	character.Stats = &models.Stats{}
	query = `
		SELECT * FROM stats
		WHERE id=$1
	`
	err = db.Get(character.Stats, query, id)
	if err != nil {
		return nil, fmt.Errorf("could not get character: %v", err)
	}

	character.Customization = &models.Customization{}
	query = `
		SELECT * FROM customizations
		WHERE id=$1
	`
	err = db.Get(character.Customization, query, id)
	if err != nil {
		return nil, fmt.Errorf("could not get character: %v", err)
	}

	return character, nil
}
