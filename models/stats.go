package models

import "fmt"

type Stats struct {
	Strength     uint8 `db:"strength"`
	Dexterity    uint8 `db:"dexterity"`
	Constitution uint8 `db:"constitution"`
	Intelligence uint8 `db:"intelligence"`
	Wisdom       uint8 `db:"wisdom"`
	Charisma     uint8 `db:"charisma"`
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
