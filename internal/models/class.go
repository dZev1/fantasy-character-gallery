package models

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
