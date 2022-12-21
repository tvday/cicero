package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cocktail struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	History  string             `json:"history" bson:"history,omitempty"`
	Recipe   string             `json:"recipe" bson:"recipe,omitempty"`
	Contents []CocktailContent  `json:"contents" bson:"contents,omitempty"`
}

type CocktailContent struct {
	Amount      float64 `json:"amount" bson:"amount,omitempty"`
	Measurement string  `json:"measurement" bson:"measurement,omitempty"`
	Category    string  `json:"category" bson:"category,omitempty"`
	Type        string  `json:"type" bson:"type,omitempty"`
	Subtype     string  `json:"subtype" bson:"subtype,omitempty"`
}
