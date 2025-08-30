package models

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

type BodyType uint8

const (
	TypeA BodyType = iota
	TypeB
)

func (bt BodyType) String() string {
	return [...]string{"Type A", "Type B"}[bt]
}