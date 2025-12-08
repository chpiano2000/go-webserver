package services

import (
	logger "github.com/sirupsen/logrus"

	"github.com/go-webserver/internal/domains"
	"github.com/go-webserver/internal/interfaces/recipe"
	"github.com/go-webserver/internal/models"
)

type recipeService struct {
	recipeRepo recipe.RecipeRepo
}

func NewService(recipeRepo recipe.RecipeRepo) recipe.RecipeUseCase {
	return &recipeService{recipeRepo: recipeRepo}
}

func (s *recipeService) Create(request *models.RecipeRequest) domains.Result[*models.Recipe] {
	createResult := s.recipeRepo.Create(request)
	if createResult.IsFailure() {
		logger.Errorf("recipeService::Create::Create %v", createResult.Error())
		return domains.Failure[*models.Recipe](*createResult.Error())
	}
	getResult := s.recipeRepo.Get(createResult.Value())
	if getResult.IsFailure() {
		logger.Errorf("recipeService::Create::Get %v", getResult.Error())
		return domains.Failure[*models.Recipe](*createResult.Error())
	}
	return domains.Success(getResult.Value())
}

func (s *recipeService) List(opts *models.RecipeFilter) domains.Result[[]*models.Recipe] {
	listResult := s.recipeRepo.List(opts)
	if listResult.IsFailure() {
		logger.Errorf("recipeService::List - %v", listResult.Error())
		return domains.Failure[[]*models.Recipe](*listResult.Error())
	}
	return domains.Success(listResult.Value())
}

func (s *recipeService) Get(id string) domains.Result[*models.Recipe] {
	getResult := s.recipeRepo.Get(id)
	if getResult.IsFailure() {
		logger.Errorf("recipeService::Get - %v", getResult.Error())
		return domains.Failure[*models.Recipe](*getResult.Error())
	}
	return domains.Success(getResult.Value())
}

func (s *recipeService) Delete(id string) domains.Result[bool] {
	delResult := s.recipeRepo.Delete(id)
	if delResult.IsFailure() {
		logger.Errorf("recipeService::Delete - %v", delResult.Error())
		return domains.Failure[bool](*delResult.Error())
	}
	return domains.Success(true)
}

func (s *recipeService) Update(request *models.RecipeUpdateRequest) domains.Result[*models.Recipe] {
	getResult := s.recipeRepo.Get(request.Id)
	if getResult.IsFailure() {
		logger.Errorf("recipeService::Update::Get - %v", getResult.IsFailure())
		return domains.Failure[*models.Recipe](*getResult.Error())
	}
	updateResult := s.recipeRepo.Update(
		request.Id,
		&request.Name,
		&request.Prep,
		&request.Cook,
		&request.Ingredients,
		&request.Instructions,
	)
	if updateResult.IsFailure() {
		logger.Errorf("recipeService::Update::Update - %v", updateResult.IsFailure())
		return domains.Failure[*models.Recipe](*updateResult.Error())
	}

	updatedResult := s.recipeRepo.Get(request.Id)
	if updatedResult.IsFailure() {
		logger.Errorf("recipeService::Update::Get %v", updatedResult.Error())
		// return nil, err
		return domains.Failure[*models.Recipe](*updatedResult.Error())

	}
	return domains.Success(updatedResult.Value())
}
