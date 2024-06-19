package server

import (
	"fmt"

	"github.com/go-webserver/config"
	docs "github.com/go-webserver/docs"
	"github.com/go-webserver/internal/api"
	"github.com/go-webserver/internal/lib"
	"github.com/go-webserver/internal/middlewares"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start() {
	fmt.Print("Starting server...")
	// Load the application configuration
	route := lib.NewRequestHandler()
	route.Gin.Use(middlewares.PanicHandler())
	cfg, err := config.LoadConfig()
	if err != nil {
		// If an error occurs while loading the configuration, panic with the error.
		panic(err)
	}
	dep := api.InitDependences(cfg)

	group := "/api/v1"
	docs.SwaggerInfo.BasePath = group
	route.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	healthRouter := api.NewHealthRouter(route)
	recipeRouter := api.NewRecipeRouter(*dep.RecipeController, route)
	accountRouter := api.NewAccountRouter(*dep.AccountController, route)
	routes := api.NewRoutes(healthRouter, recipeRouter, accountRouter)
	routes.Setup(group)

	route.Gin.Run(":5000")
}
