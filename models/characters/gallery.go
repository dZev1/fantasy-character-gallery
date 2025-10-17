package characters

type CharacterGallery interface {
	Create(character *Character) error
	Get(id ID) (*Character, error)
	GetAll() ([]Character, error)
	Edit(character *Character) error
	Remove(id ID) error
}
