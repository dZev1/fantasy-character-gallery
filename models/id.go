package models

import "fmt"

type ID uint64

func (id ID) String() string {
	return fmt.Sprintf("Nº%d", id)
}
