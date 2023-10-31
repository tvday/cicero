package controllers

import (
	"api/db"
	"api/models"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type PostIngredientInput struct {
	Category   string   `json:"category" binding:"required"`
	Type       string   `json:"type" binding:"required"`
	Subtype    string   `json:"subtype"`
	Brand      string   `json:"brand"`
	Name       string   `json:"name"`
	Abv        float64  `json:"abv"`
	FlavorTags []string `json:"flavorTags"`
}

type PatchIngredientInput struct {
	Category   string   `json:"category"`
	Type       string   `json:"type"`
	Subtype    string   `json:"subtype"`
	Brand      string   `json:"brand"`
	Name       string   `json:"name"`
	Abv        float64  `json:"abv"`
	FlavorTags []string `json:"flavorTags"`
}

// GetIngredients
//
// Method: GET
//
// Route: /ingredients
//
// Get all ingredients from database
func (ctrl *PublicController) GetIngredients(ctx *gin.Context) {
	var ingredients []models.Ingredient

	cur, err := ctrl.IngredientsCollection.Find(context.TODO(), primitive.M{})

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if err = cur.All(context.TODO(), &ingredients); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"data": ingredients})
}

// GetIngredientById
//
// Method: GET
//
// Route: /ingredients/:id
//
// Find an ingredient by id
func (ctrl *PublicController) GetIngredientById(ctx *gin.Context) {
	var ingredient models.Ingredient

	if err := db.CollectionFindByID(ctrl.IngredientsCollection, ctx.Param("id"), &ingredient); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"data": ingredient})
}

// PostIngredient
//
// Method: POST
//
// Route: /ingredients
//
// Create a new ingredient
func (ctrl *PublicController) PostIngredient(ctx *gin.Context) {
	// validate input
	var input models.Ingredient

	if err := ctx.BindJSON(&input); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inputError := InputError{}
	if input.Category == "" {
		inputError.AddMissingInput("category")
	} else {
		inputError.InvalidCategory = !models.IsValidCategory(input.Category)
	}
	if input.Type == "" {
		inputError.AddMissingInput("type")
	}
	if inputError.IsError() {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": inputError.Error()})
		return
	}

	// insert into db
	rslt, err := ctrl.IngredientsCollection.InsertOne(context.TODO(), &input)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get the id, return inserted object
	input.ID = rslt.InsertedID.(primitive.ObjectID)
	ctx.IndentedJSON(http.StatusOK, gin.H{"data": input})
}

// PatchIngredientById
//
// Method: PATCH
//
// Route: /ingredients/:id
//
// Update an ingredient
func (ctrl *PublicController) PatchIngredientById(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)

	//// get model if exists
	//var ingredient models.Ingredient
	//if err := db.CollectionFindByID(ctrl.IngredientsCollection, ctx.Param("id"), ingredient); err != nil {
	//	ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//}
	//
	//var input models.Ingredient
	//if err := ctx.BindJSON(&input); err != nil {
	//	ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//ctrl.IngredientsCollection.UpdateByID()
	//
	//ctx.IndentedJSON(http.StatusOK, gin.H{"data": ingredient})
}

// DeleteIngredientById
//
// Method: DELETE
//
// Route: /ingredients/:id
//
// Delete an ingredient
func (ctrl *PublicController) DeleteIngredientById(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := ctrl.IngredientsCollection.DeleteOne(context.TODO(), primitive.M{"_id": id})
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if model doesn't exist
	if result.DeletedCount == 0 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	ctx.Status(http.StatusNoContent)
}
