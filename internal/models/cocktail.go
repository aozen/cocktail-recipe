package models

type Cocktail struct {
	ID          string       `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string       `json:"name,omitempty" bson:"name,omitempty"`
	Ingredients []Ingredient `json:"ingredients,omitempty" bson:"ingredients,omitempty"`
	ImageURL    string       `json:"imageUrl,omitempty" bson:"imageUrl,omitempty"`
}
