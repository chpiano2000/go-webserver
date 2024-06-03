package server

import (
	"fmt"

	"github.com/go-webserver/config"
	"github.com/go-webserver/internal/api"
	"github.com/go-webserver/internal/lib"
)

func Start() {
	fmt.Print("Starting server...")
	// Load the application configuration
	cfg, err := config.LoadConfig("config")
	if err != nil {
		// If an error occurs while loading the configuration, panic with the error.
		panic(err)
	}
	dep := api.InitDependences(cfg)

	group := "/api/v1"
	route := lib.NewRequestHandler()
	healthRouter := api.NewHealthRouter(route)
	recipeRouter := api.NewRecipeRouter(*dep.RecipeController, route)
	routes := api.NewRoutes(healthRouter, recipeRouter)
	routes.Setup(group)

	route.Gin.Run(":5000")
}
