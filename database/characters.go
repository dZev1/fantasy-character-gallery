package database

import (
	"fmt"

	"github.com/dZev1/character-gallery/models"
)

const (
	charactersTable    = `characters`
	statsTable         = `stats`
	customizationTable = `customization`
	joinedTables       = `
		characters c
		JOIN stats s ON c.stats_id = s.id
		JOIN customization cu ON c.customization_id = cu.id
	`

	charactersFields    = `name, body_type, species, class`
	statsFields         = `strength, dexterity, constitution, intelligence, wisdom, charisma`
	customizationFields = `hair, face, shirt, pants, shoes`
	joinedFields        = `
	c.id, c.name, c.body_type, c.species, c.class,
    s.strength, s.dexterity, s.constitution, s.intelligence, s.wisdom, s.charisma,
    cu.hair, cu.face, cu.shirt, cu.pants, cu.shoes
	`
)

func InsertCharacter(character *models.Character) error {
	queryStats := fmt.Sprintf(`
		INSERT INTO %s(%s)
		VALUES(:strength, :dexterity, :constitution, :intelligence, :wisdom, :charisma);
	`, statsTable, statsFields)

	result, err := db.NamedExec(queryStats, character.Stats)
	if err != nil {
		return fmt.Errorf("could not insert stats: %v", err)
	}

	statsID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("could not get stats id: %v", err)
	}

	queryCustomization := fmt.Sprintf(`
		INSERT INTO %s(%s)
		VALUES(:hair, :face, :shirt, :pants, :shoes);
	`, customizationTable, customizationFields)

	result, err = db.NamedExec(queryCustomization, character.Customization)
	if err != nil {
		return fmt.Errorf("could not insert customization: %v", err)
	}

	customizationID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("could not get customization id: %v", err)
	}

	queryChar := fmt.Sprintf(`
		INSERT INTO %s(%s, stats_id, customization_id)
		VALUES(:name, :body_type, :species, :class, :stats_id, :customization_id);
	`, charactersTable, charactersFields)

	_, err = db.NamedExec(queryChar, map[string]any{
		"name":             character.Name,
		"body_type":        character.BodyType,
		"species":          character.Species,
		"class":            character.Class,
		"stats_id":         statsID,
		"customization_id": customizationID,
	})

	if err != nil {
		return fmt.Errorf("could not insert character: %v", err)
	}

	return nil
}

func DeleteCharacter(charName string) error {
	if !isCharInDB(charName) {
		return fmt.Errorf("character %s does not exist in the database", charName)
	}

	var charID, statsID, customizationID models.ID
	query := fmt.Sprintf(`
		SELECT id, stats_id, customization_id FROM %s
		WHERE name=?
	`, charactersTable)

	err := db.QueryRow(query, charName).Scan(&charID, &statsID, &customizationID)
	if err != nil {
		return fmt.Errorf("could not get character id: %v", err)
	}

	query = fmt.Sprintf(`
		DELETE FROM %s WHERE id=?
	`, charactersTable)
	_, err = db.Exec(query, charID)
	if err != nil {
		return fmt.Errorf("could not delete character from database: %v", err)
	}

	query = fmt.Sprintf(`
		DELETE FROM %s WHERE id=?
	`, statsTable)
	_, err = db.Exec(query, statsID)
	if err != nil {
		return fmt.Errorf("could not delete character from database: %v", err)
	}

	query = fmt.Sprintf(`
		DELETE FROM %s WHERE id=?
	`, customizationTable)
	_, err = db.Exec(query, customizationID)
	if err != nil {
		return fmt.Errorf("could not delete character from database: %v", err)
	}

	return nil
}

func isCharInDB(charName string) bool {
	var exists bool
	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE name = ?)`, charactersTable)
	err := db.Get(&exists, query, charName)
	if err != nil {
		return false
	}
	return exists
}
