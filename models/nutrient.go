package models

type Part string

const (
	A Part = "A"
	B Part = "B"
	C Part = "C"
)

type Nutrient struct {
	NutrientID string `json:"nutrient_id"`
	Part       Part   `json:"part"`
	N          int    `json:"n"`
	P          int    `json:"p"`
	K          int    `json:"k"`
}
