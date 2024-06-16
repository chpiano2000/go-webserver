package api

import (
	"github.com/go-webserver/config"
	"github.com/go-webserver/internal/controllers"
	"github.com/go-webserver/internal/databases"
	"github.com/go-webserver/internal/repositories"
	"github.com/go-webserver/internal/services"
)

type Dependences struct {
	RecipeController  *controllers.RecipeController
	AccountController *controllers.AccountController
}

func InitDependences(cfg config.Config) Dependences {
	mongoClient := databases.NewMongoDB(cfg)

	// Init Repositories
	recipeRepo := repositories.NewMongoRecipeRepo(mongoClient)
	accountRepo := repositories.NewMongoAccountRepo(mongoClient)
	authRepo := repositories.NewAuthRepo()
	keysRepo := repositories.NewMongoKeyRepo(mongoClient)

	// Init Services
	recipeService := services.NewService(recipeRepo)
	accountService := services.NewAccountService(accountRepo, authRepo, keysRepo)

	// Init Controllers
	recipeControllers := controllers.NewRecipeController(recipeService)
	accountControllers := controllers.NewAccountController(accountService)

	dep := Dependences{
		RecipeController:  &recipeControllers,
		AccountController: &accountControllers,
	}
	return dep
}
