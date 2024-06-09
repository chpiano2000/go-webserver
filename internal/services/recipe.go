package services

import (
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
	recipeId, err := s.recipeRepo.Create(request)
	if err != nil {
		logger.Errorf("recipeService::Create::Create %v", err)
		return nil, err
	}
	recipe, err := s.recipeRepo.Get(recipeId)
	if err != nil {
		logger.Errorf("recipeService::Create::Get %v", err)
		return nil, err
	}
	return recipe, nil
}

func (s *recipeService) List() ([]*models.Recipe, error) {
	recipes, err := s.recipeRepo.List()
	if err != nil {
		logger.Errorf("recipeService::List - %v", err)
		return nil, err
	}
	return recipes, nil
}

func (s *recipeService) Get(id string) (*models.Recipe, error) {
	recipe, err := s.recipeRepo.Get(id)
	if err != nil {
		logger.Errorf("recipeService::Get - %v", err)
		return nil, err
	}
	return recipe, nil
}

func (s *recipeService) Delete(id string) error {
	err := s.recipeRepo.Delete(id)
	if err != nil {
		logger.Errorf("recipeService::Delete - %v", err)
		return err
	}
	return nil
}

func (s *recipeService) Update(request *models.RecipeUpdateRequest) (*models.Recipe, error) {
	_, err := s.recipeRepo.Get(request.Id)
	if err != nil {
		logger.Errorf("recipeService::Update::Get - %v", err)
		return nil, err
	}
	err = s.recipeRepo.Update(
		request.Id,
		&request.Name,
		&request.Prep,
		&request.Cook,
		&request.Ingredients,
		&request.Instructions,
	)
	if err != nil {
		logger.Errorf("recipeService::Update::Update - %v", err)
		return nil, err
	}

	recipe, err := s.recipeRepo.Get(request.Id)
	if err != nil {
		logger.Errorf("recipeService::Update::Get %v", err)
		return nil, err
	}
	return recipe, nil
}
