package models

import "fmt"

type Stats struct {
	ID           ID    `db:"id" json:"id"`
	Strength     uint8 `db:"strength" json:"strength"`
	Dexterity    uint8 `db:"dexterity" json:"dexterity"`
	Constitution uint8 `db:"constitution" json:"constitution"`
	Intelligence uint8 `db:"intelligence" json:"intelligence"`
	Wisdom       uint8 `db:"wisdom" json:"wisdom"`
	Charisma     uint8 `db:"charisma" json:"charisma"`
}

func (s *Stats) String() string {
	return fmt.Sprintf("STR: %v\nDEX: %v\nCON: %v\nINT: %v\nWIS: %v\nCHA: %v",
		s.Strength,
		s.Dexterity,
		s.Constitution,
		s.Intelligence,
		s.Wisdom,
		s.Charisma,
	)
}
