package controllers

import (
	"api_sql/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostIngredientInput struct {
	Category string `json:"category" binding:"required"`
	Type     string `json:"type" binding:"required"`
	Subtype  string `json:"subtype"`
}

type PatchIngredientInput struct {
	Category string `json:"category"`
	Type     string `json:"type"`
	Subtype  string `json:"subtype"`
}

// GetIngredients
//
// Method: GET
//
// Route: /ingredients
//
// Get all ingredients from database
func GetIngredients(c *gin.Context) {
	var ingredients []models.Ingredient
	models.DB.Find(&ingredients)

	c.IndentedJSON(http.StatusOK, gin.H{"data": ingredients})
}

// GetIngredientById
//
// Method: GET
//
// Route: /ingredients/:id
//
// Find an ingredient by id
func GetIngredientById(c *gin.Context) {
	var ingredient models.Ingredient

	if err := models.DB.Where("id =?", c.Param("id")).First(&ingredient).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": ingredient})
}

// PostIngredient
//
// Method: POST
//
// Route: /ingredients
//
// Create a new ingredient
func PostIngredient(c *gin.Context) {
	// Validate input
	var input PostIngredientInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create ingredient
	ingredient := models.Ingredient{
		Category: input.Category,
		Type:     input.Type,
		Subtype:  input.Subtype,
	}
	models.DB.Create(&ingredient)

	c.IndentedJSON(http.StatusOK, gin.H{"data": ingredient})
}

// PatchIngredientById
//
// Method: PATCH
//
// Route: /ingredients/:id
//
// Update an ingredient
func PatchIngredientById(c *gin.Context) {
	// get model if exists
	var ingredient models.Ingredient
	if err := models.DB.Where("id = ?", c.Param("id")).First(&ingredient).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// validate input
	var input PatchIngredientInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// patch object
	models.DB.Model(&ingredient).Updates(input)

	c.IndentedJSON(http.StatusOK, gin.H{"data": ingredient})
}

// DeleteIngredientById
//
// Method: DELETE
//
// Route: /ingredients/:id
//
// Delete an ingredient
func DeleteIngredientById(c *gin.Context) {
	// get model if exists
	var ingredient models.Ingredient
	if err := models.DB.Where("id = ?", c.Param("id")).First(&ingredient).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&ingredient)
	c.Status(http.StatusNoContent)
}
