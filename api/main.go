package main

import (
	"api/controllers"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const ADDRESS string = ":8080"

func runServer(addr string, controller *controllers.PublicController) {
	r := gin.Default()

	r.GET("/ingredients", controller.GetIngredients)
	r.POST("/ingredients", controller.PostIngredient)
	r.GET("/ingredients/:id", controller.GetIngredientById)
	r.PATCH("/ingredients/:id", controller.PatchIngredientById)
	r.DELETE("/ingredients/:id", controller.DeleteIngredientById)

	r.GET("/cocktails", controller.GetCocktails)
	//r.POST("/cocktails", controller.PostCocktail)
	//r.GET("/cocktails/:id", controller.GetCocktailById)
	//r.PATCH("/cocktails/:id", controller.UpdateCocktailById)
	//r.DELETE("/cocktails/:id", controller.DeleteCocktailById)

	//r.POST("/cocktails/:id", controller.AddCocktailContent)
	//
	//r.POST("/query/cocktails", controller.PostDiscoverCocktails)
	//r.POST("/query/cocktails/:id", controller.PostFindMissing)

	err := r.Run(addr)
	log.Fatal(err)
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	pCtrl := controllers.PublicController{
		DB:                    client.Database("cicero"),
		IngredientsCollection: client.Database("cicero").Collection("ingredients"),
	}

	runServer(ADDRESS, &pCtrl)
}
