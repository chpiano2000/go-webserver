package services

import (
	"fmt"

	"github.com/go-webserver/internal/interfaces/recipe"
	"github.com/go-webserver/internal/models"
)

type RecipeService struct {
	recipeRepo recipe.RecipeRepo
}

func NewService(recipeRepo recipe.RecipeRepo) RecipeService {
	return RecipeService{recipeRepo: recipeRepo}
}

func (s RecipeService) Create(request *models.RecipeRequest) (*models.Recipe, error) {
	recipe, err := s.recipeRepo.Create(request)
	if err != nil {
		return nil, fmt.Errorf("failed to create recipe: %w", err)
	}
	return recipe, nil
}
