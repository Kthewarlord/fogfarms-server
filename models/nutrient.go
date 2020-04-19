package models

type Part string

const (
	A Part = "A"
	B Part = "B"
	C Part = "C"
)

type Nutrient struct {
	NutrientID int  `json:"nutrient_id"`
	Part       Part `json:"part"`
	Nitrogen   int  `json:"nitrogen"`
	Phosphorus int  `json:"phosphorus"`
	Potassium  int  `json:"potassium"`
}
