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
		c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			Status:  http.StatusUnprocessableEntity,
			Code:    "UnprocessableEntity",
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	recipeRequest := models.RecipeRequest{
		Name:         recipeSchemas.Name,
		Prep:         recipeSchemas.Prep,
		Cook:         recipeSchemas.Cook,
		Ingredients:  recipeSchemas.Ingredients,
		Instructions: recipeSchemas.Instructions,
	}
	createResult := rc.service.Create(&recipeRequest)
	if createResult.IsFailure() {
		err := createResult.Error()
		status := utils.MapErrorToStatus(*createResult.Error())
		c.JSON(status, response.ErrorResponse{
			Status:  status,
			Code:    err.Code,
			Message: err.Message,
			Data:    nil,
		})
	}
	successCode := "RecipeCreated"
	successMessage := "Recipe Created Successfully"
	c.JSON(http.StatusCreated, response.Created(successCode, successMessage, createResult.Value()))
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
		c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			Status:  http.StatusUnprocessableEntity,
			Code:    "UnprocessableEntity",
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}
	listResult := rc.service.List(&queryParams)
	if listResult.IsFailure() {
		err := listResult.Error()
		status := utils.MapErrorToStatus(*listResult.Error())
		c.JSON(status, response.ErrorResponse{
			Status:  status,
			Code:    err.Code,
			Message: err.Message,
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, response.OK(listResult.Value()))
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
	result := rc.service.Get(id)
	if result.IsFailure() {
		err := result.Error()
		status := utils.MapErrorToStatus(*result.Error())
		c.JSON(status, response.ErrorResponse{
			Status:  status,
			Code:    err.Code,
			Message: err.Message,
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, response.OK(result.Value()))
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
	result := rc.service.Delete(id)
	if result.IsFailure() {
		err := result.Error()
		status := utils.MapErrorToStatus(*result.Error())
		c.JSON(status, response.ErrorResponse{
			Status:  status,
			Code:    err.Code,
			Message: err.Message,
			Data:    nil,
		})
		return
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
		c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			Status:  http.StatusUnprocessableEntity,
			Code:    "UnprocessableEntity",
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	result := rc.service.Update(&models.RecipeUpdateRequest{
		Id:           c.Param("Id"),
		Name:         payload.Name,
		Prep:         payload.Prep,
		Cook:         payload.Cook,
		Ingredients:  payload.Ingredients,
		Instructions: payload.Instructions,
	})
	if result.IsFailure() {
		err := result.Error()
		status := utils.MapErrorToStatus(*err)
		c.JSON(status, response.ErrorResponse{
			Status:  status,
			Code:    err.Code,
			Message: err.Message,
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, response.OK(result.Value()))
}
