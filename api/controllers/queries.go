package controllers

import (
	"api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IngredientIDListInput struct {
	IngredientIDs []uint `json:"ingredientIDs" binding:"required"`
}

//type FindMissingInput struct {
//	CocktailID    uint   `json:"cocktailID" binding:"required"`
//	IngredientIDs []uint `json:"ingredientIDs" binding:"required"`
//}

type CocktailResult struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type MissingResult struct {
	ID       uint   `json:"id"`
	Category string `json:"category"`
	Type     string `json:"type"`
	Subtype  string `json:"subtype"`
}

func PostFindMissing(c *gin.Context) {
	var input IngredientIDListInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cID, err := (strconv.ParseUint(c.Param("id"), 10, 32))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ingredients, err := WhatsMissing(uint(cID), input.IngredientIDs)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": ingredients})
}

func PostDiscoverCocktails(c *gin.Context) {
	var input IngredientIDListInput
	if err := c.ShouldBind(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	drinks, err := FindPossibleCocktails(input.IngredientIDs, 0)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": drinks})
}

// WhatsMissing
//
// Return a list of ingredients that are missing for a cocktail based on ingredient ids provided
func WhatsMissing(cocktailId uint, ingredients []uint) ([]MissingResult, error) {
	iIdList := "("
	for i, n := range ingredients {
		if i != 0 {
			iIdList += fmt.Sprintf(", %v", n)
		} else {
			iIdList += fmt.Sprintf("%v", n)
		}
	}
	iIdList += ")"

	var result []MissingResult

	err := models.DB.Raw(`
		SELECT i.id, i.category, i.type, i.subtype
		FROM cocktail_contents cc inner join ingredients i on i.id = cc.ingredient_id
		WHERE cc.cocktail_id = ? AND
		      cc.ingredient_id NOT IN `+iIdList, cocktailId, ingredients).
		Scan(&result).Error

	return result, err
}

// FindPossibleCocktails
//
// Return a list of Cocktails that can be made with subset of given ingredients.
// Ignores basic ingredients.
// Will return cocktails missing <tolerance> amount of ingredients.
func FindPossibleCocktails(ingredients []uint, tolerance uint) ([]CocktailResult, error) {
	iIdList := "("
	for i, n := range ingredients {
		if i != 0 {
			iIdList += fmt.Sprintf(", %v", n)
		} else {
			iIdList += fmt.Sprintf("%v", n)
		}
	}
	iIdList += ")"

	q := fmt.Sprintf(`
		SELECT c.name, c.id
		FROM cocktails c
		WHERE
		    (SELECT count(*)
		     FROM cocktail_contents cc inner join ingredients i on i.id = cc.ingredient_id
		     WHERE
		         cc.cocktail_id = c.id AND
		         i.category != 'basic')
		        -
		    (SELECT count(*)
		     FROM cocktail_contents cc inner join ingredients i on i.id = cc.ingredient_id
		     WHERE
		         cc.cocktail_id = c.id AND
		         i.category != 'basic' AND
		         cc.ingredient_id in %v)
		    <= %v`,
		iIdList, tolerance)

	var result []CocktailResult

	err := models.DB.Raw(q).Scan(&result).Error

	return result, err
}
