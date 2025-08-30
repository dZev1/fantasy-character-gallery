package models

import "fmt"

type ID uint64

type Character struct {
	ID            ID   `db:"id"`
	Name          string   `db:"name"`
	BodyType      BodyType `db:"body_type"`
	Species       Species  `db:"species"`
	Class         Class    `db:"class"`
	Stats         *Stats
	Customization *Customization
}

func (char *Character) String() string {
	return fmt.Sprintf("\nName: %v\nSpecies: %v\nBody Type: %v\nClass: %v\n\nStats\n%v", char.Name, char.Species, char.BodyType, char.Class, char.Stats.String())
}
