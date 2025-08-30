package models

type Customization struct {
	Hair  uint8 `db:"hair"`
	Face  uint8 `db:"face"`
	Shirt uint8 `db:"shirt"`
	Pants uint8 `db:"pants"`
	Shoes uint8 `db:"shoes"`
}