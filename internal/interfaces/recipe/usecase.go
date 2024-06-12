package recipe

import "github.com/go-webserver/internal/models"

type RecipeUseCase interface {
	Create(request *models.RecipeRequest) (*models.Recipe, error)
	List(opts *models.RecipeFilter) ([]*models.Recipe, error)
	Get(id string) (*models.Recipe, error)
	Delete(id string) error
	Update(request *models.RecipeUpdateRequest) (*models.Recipe, error)
}
