package models

import "fmt"

type Character struct {
	ID            uint64         `json:"id"`
	Creator       string         `json:"creator"`
	Name          string         `json:"name"`
	BodyType      BodyType       `json:"body_type"`
	Species       Species        `json:"species"`
	Class         Class          `json:"class"`
	Stats         *Stats         `json:"stats"`
	Customization *Customization `json:"customization"`
}

func (char *Character) String() string {
	return fmt.Sprintf("\nName: %v\nSpecies: %v\nBody Type: %v\nClass: %v\n\nStats\n%v", char.Name, char.Species, char.BodyType, char.Class, char.Stats.String())
}
