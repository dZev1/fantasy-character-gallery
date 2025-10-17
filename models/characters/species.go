package characters

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
