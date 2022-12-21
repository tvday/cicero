package models

import "github.com/jinzhu/gorm"

type Cocktail struct {
	Name    string `json:"name"`
	History string `json:"history"`
	Recipe  string `json:"recipe"`
	gorm.Model
	Contents []CocktailContent
}

type CocktailContent struct {
	CocktailID   uint `gorm:"foreignKey:CocktailID;primaryKey"`
	IngredientID uint `gorm:"foreignKey:IngredientID;primaryKey"`
	Ingredient   Ingredient
	Amount       float64 `json:"amount"`
	Measurement  string  `json:"measurement"`
}
