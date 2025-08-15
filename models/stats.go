package models

import "fmt"

type Stats struct {
	Strength     uint8 `json:"strength"`
	Dexterity    uint8 `json:"dexterity"`
	Constitution uint8 `json:"constitution"`
	Intelligence uint8 `json:"intelligence"`
	Wisdom       uint8 `json:"wisdom"`
	Charisma     uint8 `json:"charisma"`
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
