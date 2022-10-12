package main

import (
	"api/controllers"
	"api/models"
	"github.com/gin-gonic/gin"
	"log"
)

const ADDRESS string = ":3000"

func runServer(addr string) {
	r := gin.Default()

	r.GET("/ingredients", controllers.GetIngredients)
	r.POST("/ingredients", controllers.PostIngredient)
	r.GET("/ingredients/:id", controllers.GetIngredientById)
	r.PATCH("/ingredients/:id", controllers.PatchIngredientById)
	r.DELETE("/ingredients/:id", controllers.DeleteIngredientById)

	r.GET("/cocktails", controllers.GetCocktails)
	r.POST("/cocktails", controllers.PostCocktail)
	r.GET("/cocktails/:id", controllers.GetCocktailById)
	r.PATCH("/cocktails/:id", controllers.UpdateCocktailById)
	r.DELETE("/cocktails/:id", controllers.DeleteCocktailById)

	r.POST("/cocktails/:id", controllers.AddCocktailContent)

	r.POST("/query/cocktails", controllers.PostDiscoverCocktails)
	r.POST("/query/cocktails/:id", controllers.PostFindMissing)

	err := r.Run(addr)
	log.Fatal(err)
}

func main() {

	models.ConnectDatabase()

	runServer(ADDRESS)
}
