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
	Missing         []string
	InvalidCategory bool
}

func (e *InputError) Error() string {
	out := ""
	for field := range e.Missing {
		out += fmt.Sprintf("Missing required field %v.\n", field)
	}
	if e.InvalidCategory {
		out += fmt.Sprintf("Invalid category provided.")
	}
	return out
}

func (e *InputError) AddMissingInput(s string) {
	e.Missing = append(e.Missing, s)
}

func (e *InputError) IsError() bool {
	return len(e.Missing) != 0 || e.InvalidCategory
}
