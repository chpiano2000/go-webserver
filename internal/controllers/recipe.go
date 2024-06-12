package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-webserver/internal/interfaces/recipe"
	"github.com/go-webserver/internal/models"
	"github.com/go-webserver/internal/response"
	"github.com/go-webserver/internal/schemas"
	"github.com/go-webserver/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type RecipeController struct {
	service recipe.RecipeUseCase
}

func NewRecipeController(useCase recipe.RecipeUseCase) RecipeController {
	return RecipeController{
		service: useCase,
	}
}

// CreateRecipe godoc
// @Summary Create Recipe
// @Description Create Recipe
// @Tags Recipe
// @Accept json
// @Produce json
// @Param payload body schemas.RecipeSchemaRequest true "Create recipe payload"
// @Success     200         {object}    models.Recipe
// @Failure     400         {object}    response.ErrorResponse
// @Failure     422         {object}    response.ErrorResponse
// @Failure     500         {object}    response.ErrorResponse
// @Router /recipe [post]
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

// ListRecipe godoc
// @Summary List All Recipes
// @Description List All Recipes
// @Tags Recipe
// @Accept json
// @Produce json
// @Param offset query int false "pagination offset"
// @Param size query int false "pagination size"
// @Success     200         {array}    models.Recipe
// @Failure     500         {object}    response.ErrorResponse
// @Router /recipes [get]
func (rc RecipeController) ListRecipes(c *gin.Context) {
	var queryParams models.RecipeFilter
	err := c.ShouldBindQuery(&queryParams)
	if err != nil {
		resp := utils.Serialize(c, utils.UnprocessableEntity)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, resp)
		return
	}
	recipes, err := rc.service.List(&queryParams)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, response.OK(recipes))
}

// GetRecipe godoc
// @Summary Get One Recipe
// @Description Get One Recipe
// @Tags Recipe
// @Accept json
// @Produce json
// @Param recipe_id path string true "Recipe Id"
// @Success     200         {object}     models.Recipe
// @Failure     400         {object}    response.ErrorResponse
// @Failure     500         {object}    response.ErrorResponse
// @Router /recipe/{recipe_id} [get]
func (rc RecipeController) GetRecipe(c *gin.Context) {
	id := c.Param("Id")
	recipe, err := rc.service.Get(id)
	if err != nil {
		log.Info(err)
		panic(err)
	}
	c.JSON(http.StatusOK, response.OK(recipe))
}

// DeleteRecipe godoc
// @Summary Delete Recipe
// @Description Delete Recipe
// @Tags Recipe
// @Accept json
// @Produce json
// @Param recipe_id path string true "Recipe Id"
// @Success     200         {object}    string
// @Failure     400         {object}    response.ErrorResponse
// @Failure     500         {object}    response.ErrorResponse
// @Router /recipe/{recipe_id} [delete]
func (rc RecipeController) DeleteRecipe(c *gin.Context) {
	id := c.Param("Id")
	err := rc.service.Delete(id)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, utils.Serialize(c, utils.DeleteRecipeSuccessfully))
}

// UpdateRecipe godoc
// @Summary Update Recipe
// @Description Update Recipe
// @Tags Recipe
// @Accept json
// @Produce json
// @Param recipe_id path string true "Recipe Id"
// @Param offset query int false "offset"
// @Param size query int false "size"
// @Success     200         {object}     models.Recipe
// @Failure     400         {object}    response.ErrorResponse
// @Failure     500         {object}    response.ErrorResponse
// @Router /recipe/{recipe_id} [patch]
func (rc RecipeController) UpdateRecipe(c *gin.Context) {
	var payload schemas.RecipeSchemaPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		resp := utils.Serialize(c, utils.UnprocessableEntity)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, resp)
		return
	}

	recipe, err := rc.service.Update(&models.RecipeUpdateRequest{
		Id:           c.Param("Id"),
		Name:         payload.Name,
		Prep:         payload.Prep,
		Cook:         payload.Cook,
		Ingredients:  payload.Ingredients,
		Instructions: payload.Instructions,
	})
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, response.OK(recipe))
}
