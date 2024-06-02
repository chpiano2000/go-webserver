package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-webserver/internal/interfaces/recipe"
	"github.com/go-webserver/internal/models"
)

type RecipeController struct {
	service recipe.RecipeUseCase
}

func NewRecipeController(useCase recipe.RecipeUseCase) RecipeController {
	return RecipeController{
		service: useCase,
	}
}

func (rc RecipeController) CreateRecipe(c *gin.Context) {
	var recipeRequest models.RecipeRequest
	err := c.ShouldBindJSON(&recipeRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	recipe, err := rc.service.Create(&recipeRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"data": recipe})
}
