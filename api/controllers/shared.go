package controllers

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type PublicController struct {
	DB                    *mongo.Database
	IngredientsCollection *mongo.Collection
	CocktailsCollection   *mongo.Collection
}

type InputError struct {
	Missing []string
}

func (e *InputError) Error() string {
	out := ""
	for field := range e.Missing {
		out += fmt.Sprintf("Missing required field %v.\n", field)
	}
	return out
}
