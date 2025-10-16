package database

import (
	"errors"
	"fmt"

	"github.com/dZev1/character-gallery/models"
)

// Errors
var (
	ErrCouldNotInsert              = errors.New(`could not insert character`)
	ErrCouldNotGet                 = errors.New(`could not get character`)
	ErrCouldNotFind                = errors.New(`could not find character`)
	ErrFailedInitializeTransaction = errors.New(`failed to initialize transaction`)
	ErrFailCommitTransaction       = errors.New(`failed to commit transaction`)
)

func CreateCharacter(character *models.Character) error {
	tx, err := db.Beginx()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrFailedInitializeTransaction, err)
	}
	defer tx.Rollback()

	query := `
		INSERT INTO characters (name, body_type, species, class)
		VALUES (:name, :body_type, :species, :class) RETURNING id
	`

	stmt, err := db.PrepareNamed(query)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCouldNotInsert, err)
	}
	defer stmt.Close()

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

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("%w: %w", ErrFailCommitTransaction, err)
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
		SELECT
			c.id, c.name, c.body_type, c.species, c.class,

			COALESCE(s.strength, 0) AS "stats.strength",
			COALESCE(s.dexterity, 0) AS "stats.dexterity",
			COALESCE(s.constitution, 0) AS "stats.constitution",
            COALESCE(s.intelligence, 0) AS "stats.intelligence",
            COALESCE(s.wisdom, 0)       AS "stats.wisdom",
            COALESCE(s.charisma, 0)     AS "stats.charisma",

			COALESCE(cust.hair, 0)  AS "customization.hair",
            COALESCE(cust.face, 0)  AS "customization.face",
            COALESCE(cust.shirt, 0) AS "customization.shirt",
            COALESCE(cust.pants, 0) AS "customization.pants",
            COALESCE(cust.shoes, 0) AS "customization.shoes"

			FROM
            	characters c
        	LEFT JOIN
            	stats s ON c.id = s.id
        	LEFT JOIN
            	customizations cust ON c.id = cust.id
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
	tx, err := db.Beginx()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrFailedInitializeTransaction, err)
	}
	defer tx.Rollback()

	err = UpdateBaseCharacter(character)
	if err != nil {
		return err
	}

	err = UpdateCustomization(character.Customization)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("%w: %w", ErrFailCommitTransaction, err)
	}

	return nil
}

func UpdateBaseCharacter(character *models.Character) error {
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
		return fmt.Errorf("%w: %v", ErrCouldNotFind, err)
	}
	return nil
}

func UpdateCustomization(customization *models.Customization) error {

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
		return fmt.Errorf("%w: %v", ErrCouldNotFind, err)
	}

	return nil
}

func UpdateStats(stats *models.Stats) error {
	tx, err := db.Beginx()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrFailedInitializeTransaction, err)
	}
	defer tx.Rollback()

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

	_, err = db.NamedExec(query, stats)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCouldNotFind, err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("%w: %w", ErrFailCommitTransaction, err)
	}

	return nil
}

func RemoveCharacterByID(id models.ID) error {
	tx, err := db.Beginx()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrFailedInitializeTransaction, err)
	}
	defer tx.Rollback()

	query := `
		DELETE FROM characters
		WHERE ID=$1
	`

	_, err = db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCouldNotFind, err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("%w: %w", ErrFailCommitTransaction, err)
	}

	return nil
}
