package database

import (
	"fmt"

	"github.com/dZev1/fantasy-character-gallery/models"
)

const characterFields string = "creator, name, body_type, species, class, stats, customization" 

func UploadCharacter(char *models.Character) error {
	query := `
		INSERT INTO characters
		($1)
		VALUES ($2, $3, $4, $5, $6, $7, $8)
	`

	_, err := db.Exec(query, characterFields, char.Creator, char.Name, char.BodyType, char.Species, char.Class, char.Stats, char.Customization)
	if err != nil {
		return fmt.Errorf("could not upload character to database: %v", err)
	}

	return nil
}

func GetCharacterID(creator, name string) (uint64,error) {
	query := `
		SELECT id FROM characters
		WHERE creator=$1 AND name=$2
	`

	var charID uint64
	err := db.QueryRow(query, creator, name).Scan(&charID)
	if err != nil {
		return 0, fmt.Errorf("character not in database: %v", err)
	}

	return charID, nil
}
