package database

import (
	"errors"
	"fmt"
	"log"

	"github.com/dZev1/character-gallery/internal/models"
	"github.com/jmoiron/sqlx"
)

// Errors
var (
	ErrCouldNotInsert = errors.New(`could not insert character`)
	ErrCouldNotGet    = errors.New(`could not get character`)
)

func InsertCharacter(character *models.Character) error {
	query := `
		INSERT INTO characters (name, body_type, species, class)
		VALUES (:name, :body_type, :species, :class) RETURNING id
	`

	stmt, err := db.PrepareNamed(query)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCouldNotInsert, err)
	}

	defer func(stmt *sqlx.NamedStmt) {
		err = stmt.Close()
		if err != nil {
			log.Fatalf("%s: %v", ErrCouldNotInsert.Error(), err)
		}
	}(stmt)

	err = stmt.Get(&character.ID, character)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCouldNotInsert, err)
	}

	character.Stats.ID = character.ID

	query = `
		INSERT INTO stats (id, strength, dexterity, constitution, intelligence, wisdom, charisma)
		VALUES(:id, :strength, :dexterity, :constitution, :intelligence, :wisdom, :charisma)
	`
	_, err = db.NamedExec(query, character.Stats)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCouldNotInsert, err)
	}

	character.Customization.ID = character.ID
	query = `
		INSERT INTO customizations (id, hair, face, shirt, pants, shoes)
		VALUES(:id, :hair, :face, :shirt, :pants, :shoes)
	`
	_, err = db.NamedExec(query, character.Customization)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCouldNotInsert, err)
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
		return nil, fmt.Errorf("%w: %v", ErrCouldNotGet, err)
	}

	character.Stats, err = GetStatsByID(id)
	if err != nil {
		return nil, err
	}

	character.Customization, err = GetCustomizationByID(id)
	if err != nil {
		return nil, err
	}

	return character, nil
}

func GetCustomizationByID(id models.ID) (*models.Customization, error) {
	customization := &models.Customization{}
	query := `
			SELECT * FROM customizations
			WHERE id = $1
		`
	err := db.Get(customization, query, id)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCouldNotGet, err)
	}
	return customization, nil
}

func GetStatsByID(id models.ID) (*models.Stats, error) {
	stats := &models.Stats{}
	query := `
			SELECT * FROM stats
			WHERE id = $1
		`
	err := db.Get(stats, query, id)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCouldNotGet, err)
	}
	return stats, nil
}

func GetCharacters() ([]models.Character, error) {
	var characters []models.Character
	query := `
		SELECT * FROM characters
	`

	err := db.Select(&characters, query)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCouldNotGet, err)
	}

	for i := range characters {
		currID := characters[i].ID
		characters[i].Stats, err = GetStatsByID(currID)
		if err != nil {
			return nil, err
		}
		characters[i].Customization, err = GetCustomizationByID(currID)
		if err != nil {
			return nil, err
		}
	}

	return characters, nil
}

func EditCharacter(character *models.Character) error {
	query := `
		UPDATE characters
		SET name = :name,
			body_type = :body_type,
			species = :species,
			class = :class
		WHERE id = :id
	`
	_, err := db.NamedExec(query, character)
	if err != nil {
		return fmt.Errorf("character not found: %v", err)
	}
	return nil
}

func EditCustomization(customization *models.Customization) error {
	query := `
		UPDATE customizations
		SET hair = :hair,
			face = :face,
			shirt = :shirt,
			pants = :pants,
			shoes = :shoes
		WHERE id = :id
	`

	_, err := db.NamedExec(query, customization)
	if err != nil {
		return fmt.Errorf("character not found: %v", err)
	}
	return nil
}

func EditStats(stats *models.Stats) error {
	query := `
		UPDATE stats
		SET strength = :strength,
			dexterity = :dexterity,
			constitution = :constitution,
			intelligence = :intelligence,
			wisdom = :wisdom,
			charisma = :charisma
		WHERE id = :id
	`

	_, err := db.NamedExec(query, stats)
	if err != nil {
		return fmt.Errorf("character not found: %v", err)
	}
	return nil
}
