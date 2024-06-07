package api

import (
	"github.com/go-webserver/internal/controllers"
	"github.com/go-webserver/internal/lib"
)

type RecipeRouter struct {
	recipeController controllers.RecipeController
	handler          lib.RequestHandler
}

func (rr RecipeRouter) Setup(group string) {
	api := rr.handler.Gin.Group(group)
	{
		api.POST("/recipe", rr.recipeController.CreateRecipe)
		api.GET("/recipes", rr.recipeController.ListRecipes)
		api.GET("/recipe/:Id", rr.recipeController.GetRecipe)
		api.DELETE("/recipe/:Id", rr.recipeController.DeleteRecipe)
	}
}

func NewRecipeRouter(svc controllers.RecipeController, handler lib.RequestHandler) RecipeRouter {
	recipeRouter := RecipeRouter{
		recipeController: svc,
		handler:          handler,
	}
	return recipeRouter
}
