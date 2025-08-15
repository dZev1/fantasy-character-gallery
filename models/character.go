package models

import "fmt"

type Species uint8

const (
	Aasimar Species = iota
	Dragonborn
	Dwarf
	Elf
	Gnome
	Goliath
	Halfling
	Human
	Orc
	Tiefling
)

func (s Species) String() string {
	return [...]string{
		"Aasimar", "Dragonborn", "Dwarf", "Elf", "Gnome",
		"Goliath", "Halfling", "Human", "Orc", "Tiefling",
	}[s]
}

type Class uint8

const (
	Barbarian Class = iota
	Bard
	Cleric
	Druid
	Fighter
	Monk
	Paladin
	Ranger
	Rogue
	Sorcerer
	Warlock
	Wizard
)

func (c Class) String() string {
	return [...]string{
		"Barbarian", "Bard", "Cleric", "Druid", "Fighter", "Monk",
		"Paladin", "Ranger", "Rogue", "Sorcerer", "Warlock", "Wizard",
	}[c]
}

type Character struct {
	ID            uint        `json:"id"`
	Creator       string        `json:"creator"`
	Name          string        `json:"name"`
	Species       Species       `json:"species"`
	Class         Class         `json:"class"`
	Stats         *Stats         `json:"stats"`
	Customization *Customization `json:"customization"`
}

func (char *Character) String() string {
	return fmt.Sprintf("Name: %v\nSpecies: %v\nClass: %v\n\nStats\n%v", char.Name, char.Species, char.Class, char.Stats.String())
}

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

type Customization struct {
	Hair  uint8
	Face  uint8
	Shirt uint8
	Pants uint8
	Shoes uint8
}