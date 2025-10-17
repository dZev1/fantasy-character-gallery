package database

import (
	"errors"
	"fmt"

	"github.com/dZev1/character-gallery/models/characters"
	"github.com/jmoiron/sqlx"
)

// Errors
var (
	ErrCouldNotInsert              = errors.New(`could not insert character`)
	ErrCouldNotGet                 = errors.New(`could not get character`)
	ErrCouldNotFind                = errors.New(`could not find character`)
	ErrFailedInitializeTransaction = errors.New(`failed to initialize transaction`)
	ErrFailCommitTransaction       = errors.New(`failed to commit transaction`)
)

type PostgresCharacterGallery struct {
	db *sqlx.DB
}

func (cg *PostgresCharacterGallery) Create(character *characters.Character) error {
	tx, err := cg.db.Beginx()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrFailedInitializeTransaction, err)
	}
	defer tx.Rollback()

	err = cg.insertBaseCharacter(tx, character)
	if err != nil {
		return err
	}

	character.Stats.ID = character.ID
	err = cg.insertStats(tx, character.Stats)
	if err != nil {
		return err
	}

	character.Customization.ID = character.ID
	err = cg.insertCustomization(tx, character.Customization)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("%w: %w", ErrFailCommitTransaction, err)
	}

	return nil
}

func (cg *PostgresCharacterGallery) Get(id characters.ID) (*characters.Character, error) {
	character, err := cg.getBaseCharacter(id)
	if err != nil {
		return nil, err
	}

	character.Stats, err = cg.getStatsByID(id)
	if err != nil {
		return nil, err
	}

	character.Customization, err = cg.getCustomizationByID(id)
	if err != nil {
		return nil, err
	}

	return character, nil
}

func (cg *PostgresCharacterGallery) GetAll() ([]characters.Character, error) {
	var chars []characters.Character
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

	err := cg.db.Select(&chars, query)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCouldNotGet, err)
	}

	return chars, nil
}

func (cg *PostgresCharacterGallery) Edit(character *characters.Character) error {
	tx, err := cg.db.Beginx()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrFailedInitializeTransaction, err)
	}
	defer tx.Rollback()

	err = cg.updateBaseCharacters(tx, character)
	if err != nil {
		return err
	}

	err = cg.updateCustomization(tx, character.Customization)
	if err != nil {
		return err
	}

	err = cg.updateStats(tx, character.Stats)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("%w: %w", ErrFailCommitTransaction, err)
	}

	return nil
}

func (cg *PostgresCharacterGallery) Remove(id characters.ID) error {
	tx, err := cg.db.Beginx()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrFailedInitializeTransaction, err)
	}
	defer tx.Rollback()

	query := `
		DELETE FROM characters
		WHERE ID=$1
	`

	_, err = tx.Exec(query, id)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCouldNotFind, err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("%w: %w", ErrFailCommitTransaction, err)
	}

	return nil
}
