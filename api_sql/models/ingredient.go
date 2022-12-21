package models

import "github.com/jinzhu/gorm"

type Ingredient struct {
	Category string `json:"category"`
	Type     string `json:"title"`
	Subtype  string `json:"subtype"`
	gorm.Model
}
