package controllers

import (
	"api/db"
	"api/models"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

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

// GetCocktailById
//
// Method: GET
//
// Route: /cocktails/:id
//
// Find a cocktail by id
func (ctrl *PublicController) GetCocktailById(ctx *gin.Context) {
	var cocktail models.Cocktail

	if err := db.CollectionFindByID(ctrl.CocktailsCollection, ctx.Param("id"), &cocktail); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"data": cocktail})
}

// PostCocktail
//
// Method: POST
//
// Route: /cocktails
//
// Create a new cocktail
func (ctrl *PublicController) PostCocktail(ctx *gin.Context) {
	// Validate input
	var input models.Cocktail
	if err := ctx.BindJSON(&input); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inputError := InputError{}
	if input.Name == "" {
		inputError.AddMissingInput("name")
	}
	for _, content := range input.Contents {
		inputError.InvalidCategory = !models.IsValidCategory(content.Category)
		if inputError.InvalidCategory {
			break
		}
	}
	if inputError.IsError() {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": inputError.Error()})
	}

	result, err := ctrl.CocktailsCollection.InsertOne(context.TODO(), &input)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	input.ID = result.InsertedID.(primitive.ObjectID)

	ctx.IndentedJSON(http.StatusOK, gin.H{"data": input})
}

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
