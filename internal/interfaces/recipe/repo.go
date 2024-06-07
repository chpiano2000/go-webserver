package recipe

import "github.com/go-webserver/internal/models"

type RecipeRepo interface {
	Create(recipe *models.RecipeRequest) (string, error)
	List() ([]*models.Recipe, error)
	Get(id string) (*models.Recipe, error)
	Delete(id string) error
}
