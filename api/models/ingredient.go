package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ingredient struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Category   string             `json:"category" bson:"category,omitempty"`
	Type       string             `json:"title" bson:"type,omitempty"`
	Subtype    string             `json:"subtype" bson:"subtype,omitempty"`
	Brand      string             `json:"brand,omitempty" bson:"brand,omitempty"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`
	Abv        string             `json:"abv,omitempty" bson:"abv,omitempty"`
	FlavorTags []string           `json:"flavorTags,omitempty" bson:"flavorTags,omitempty"`
}
