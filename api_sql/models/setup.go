package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open("sqlite3", "var/cicero.sqlite3")

	if err != nil {
		log.Panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Ingredient{}, &Cocktail{}, &CocktailContent{})

	DB = db
}
