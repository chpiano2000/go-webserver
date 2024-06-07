package services

import (
	"fmt"

	"github.com/go-webserver/internal/interfaces/recipe"
	"github.com/go-webserver/internal/models"

	logger "github.com/sirupsen/logrus"
)

type recipeService struct {
	recipeRepo recipe.RecipeRepo
}

func NewService(recipeRepo recipe.RecipeRepo) recipe.RecipeUseCase {
	return &recipeService{recipeRepo: recipeRepo}
}

func (s *recipeService) Create(request *models.RecipeRequest) (*models.Recipe, error) {
	recipe, err := s.recipeRepo.Create(request)
	if err != nil {
		logger.Errorf("recipeService::Create: %v", err)
		return nil, err
	}
	return recipe, nil
}

func (s *recipeService) List() ([]*models.Recipe, error) {
	recipes, err := s.recipeRepo.List()
	if err != nil {
		return nil, fmt.Errorf("failed to create recipe: %w", err)
	}
	return recipes, nil
}
