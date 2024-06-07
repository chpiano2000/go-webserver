package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-webserver/internal/interfaces/recipe"
	"github.com/go-webserver/internal/models"
	"github.com/go-webserver/internal/response"
	"github.com/go-webserver/internal/schemas"
	"github.com/go-webserver/pkg/utils"
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
	var recipeSchemas schemas.RecipeSchemaRequest
	if err := c.ShouldBindJSON(&recipeSchemas); err != nil {
		resp := utils.Serialize(c, utils.UnprocessableEntity)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, resp)
		return
	}

	recipeRequest := models.RecipeRequest{
		Name:         recipeSchemas.Name,
		Prep:         recipeSchemas.Prep,
		Cook:         recipeSchemas.Cook,
		Ingredients:  recipeSchemas.Ingredients,
		Instructions: recipeSchemas.Instructions,
	}
	recipe, err := rc.service.Create(&recipeRequest)
	if err != nil {
		panic(err)
	}
	successCode := "RecipeCreated"
	successMessage := "Recipe Created Successfully"
	c.JSON(http.StatusCreated, response.Created(successCode, successMessage, recipe))
}

func (rc RecipeController) ListRecipes(c *gin.Context) {
	recipes, err := rc.service.List()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, response.OK(recipes))
}
