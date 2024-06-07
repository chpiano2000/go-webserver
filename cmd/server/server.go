package server

import (
	"fmt"

	"github.com/go-webserver/config"
	docs "github.com/go-webserver/docs"
	"github.com/go-webserver/internal/api"
	"github.com/go-webserver/internal/lib"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	docs.SwaggerInfo.BasePath = group
	route.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	healthRouter := api.NewHealthRouter(route)
	recipeRouter := api.NewRecipeRouter(*dep.RecipeController, route)
	routes := api.NewRoutes(healthRouter, recipeRouter)
	routes.Setup(group)

	route.Gin.Run(":5000")
}
