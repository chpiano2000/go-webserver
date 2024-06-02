package api

import (
	"github.com/go-webserver/config"
	"github.com/go-webserver/internal/controllers"
	"github.com/go-webserver/internal/databases"
	"github.com/go-webserver/internal/repositories"
	"github.com/go-webserver/internal/services"
)

type Dependences struct {
	RecipeController *controllers.RecipeController
}

func InitDependences(cfg config.Config) Dependences {
	mongoClient := databases.NewMongoDB(cfg)
	// Recipe Dependencies
	recipeRepo := repositories.NewMongoRecipeRepo(mongoClient)
	recipeService := services.NewService(recipeRepo)
	recipeControllers := controllers.NewRecipeController(recipeService)

	dep := Dependences{
		RecipeController: &recipeControllers,
	}
	return dep
}
