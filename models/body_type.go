package models

type BodyType uint8

const (
	TypeA BodyType = iota
	TypeB
)

func (bt BodyType) String() string {
	return [...]string{"Type A", "Type B"}[bt]
}
