package controllers

import (
	"api/models"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type CreateCocktailInput struct {
	Name     string                    `json:"name" binding:"required"`
	History  string                    `json:"history"`
	Recipe   string                    `json:"recipe" binding:"required"`
	Contents []AddCocktailContentInput `json:"contents"`
}

type UpdateCocktailInput struct {
	Name    string `json:"name"`
	History string `json:"history"`
	Recipe  string `json:"recipe"`
}

type AddCocktailContentInput struct {
	IngredientID *uint    `json:"ingredientID" binding:"required"`
	Amount       *float64 `json:"amount" binding:"required"`
	Measurement  string   `json:"measurement" binding:"required"`
}

// GetCocktails
//
// Method: GET
//
// Route: /cocktails
//
// Get all cocktails from database
func (ctrl *PublicController) GetCocktails(ctx *gin.Context) {
	var cocktails []models.Cocktail

	cur, err := ctrl.CocktailsCollection.Find(context.TODO(), primitive.M{})

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if err = cur.All(context.TODO(), &cocktails); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"data": cocktails})
}

//// GetCocktailById
////
//// Method: GET
////
//// Route: /cocktails/:id
////
//// Find a cocktail by id
//func GetCocktailById(c *gin.Context) {
//	var cocktail models.Cocktail
//
//	err := models.DB.Model(&models.Cocktail{}).
//		Preload("Contents.Ingredient").
//		Where("id = ?", c.Param("id")).
//		First(&cocktail).Error
//
//	//err := models.DB.Where("id =?", c.Param("id")).First(&cocktail).Error
//
//	if err != nil {
//		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
//		return
//	}
//
//	c.IndentedJSON(http.StatusOK, gin.H{"data": cocktail})
//}
//
//// PostCocktail
////
//// Method: POST
////
//// Route: /cocktails
////
//// Create a new cocktail
//func PostCocktail(c *gin.Context) {
//	// Validate input
//	var input CreateCocktailInput
//	if err := c.ShouldBindJSON(&input); err != nil {
//		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	var contents []models.CocktailContent
//	for _, content := range input.Contents {
//		var ingredient models.Ingredient
//		if err := models.DB.Where("id =?", content.IngredientID).First(&ingredient).Error; err != nil {
//			c.IndentedJSON(http.StatusBadRequest,
//				gin.H{"error": fmt.Sprintf("Ingredient %v not found", content.IngredientID)})
//			return
//		}
//		contents = append(contents, models.CocktailContent{
//			IngredientID: *content.IngredientID,
//			Ingredient:   ingredient,
//			Amount:       *content.Amount,
//			Measurement:  content.Measurement,
//		})
//	}
//
//	// create ingredient
//	cocktail := models.Cocktail{
//		Name:     input.Name,
//		History:  input.History,
//		Recipe:   input.Recipe,
//		Contents: contents,
//	}
//	models.DB.Create(&cocktail)
//
//	c.IndentedJSON(http.StatusOK, gin.H{"data": cocktail})
//}
//
//// UpdateCocktailById
////
//// Method: PATCH
////
//// Route: /cocktails/:id
////
//// Update a cocktail
//func UpdateCocktailById(c *gin.Context) {
//	// get model if exists
//	var cocktail models.Cocktail
//	if err := models.DB.Where("id = ?", c.Param("id")).First(&cocktail).Error; err != nil {
//		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
//		return
//	}
//
//	// validate input
//	var input UpdateCocktailInput
//	if err := c.ShouldBindJSON(&input); err != nil {
//		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	// patch object
//	models.DB.Model(&cocktail).Updates(input)
//
//	c.IndentedJSON(http.StatusOK, gin.H{"data": cocktail})
//}
//
//// DeleteCocktailById
////
//// Method: DELETE
////
//// Route: /cocktails/:id
////
//// Delete a cocktail
//func DeleteCocktailById(c *gin.Context) {
//	// get model if exists
//	var cocktail models.Cocktail
//	if err := models.DB.Where("id = ?", c.Param("id")).First(&cocktail).Error; err != nil {
//		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
//		return
//	}
//
//	models.DB.Delete(&cocktail)
//	c.Status(http.StatusNoContent)
//}
//
//// AddCocktailContent
////
//// Method: POST
////
//// Route: /cocktails/:id
////
//// Add an ingredient to a cocktail
//func AddCocktailContent(c *gin.Context) {
//	// get model if exists
//	var cocktail models.Cocktail
//	if err := models.DB.Where("id = ?", c.Param("id")).First(&cocktail).Error; err != nil {
//		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Record not found"})
//		return
//	}
//
//	// validate input
//	var input AddCocktailContentInput
//	if err := c.ShouldBindJSON(&input); err != nil {
//		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//	}
//
//	var ingredient models.Ingredient
//	if err := models.DB.Where("id = ?", input.IngredientID).First(&ingredient).Error; err != nil {
//		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Ingredient %v not found", *input.IngredientID)})
//		return
//	}
//
//	content := models.CocktailContent{
//		CocktailID:   cocktail.ID,
//		IngredientID: ingredient.ID,
//		Ingredient:   ingredient,
//		Amount:       *input.Amount,
//		Measurement:  input.Measurement,
//	}
//
//	models.DB.Create(&content)
//
//	c.IndentedJSON(http.StatusOK, gin.H{"data": content})
//}
