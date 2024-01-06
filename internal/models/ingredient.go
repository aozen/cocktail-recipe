package models

type Ingredient struct {
	Name     string  `json:"name,omitempty" bson:"name,omitempty"`
	Quantity float64 `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Unit     string  `json:"unit,omitempty" bson:"unit,omitempty"`
}
