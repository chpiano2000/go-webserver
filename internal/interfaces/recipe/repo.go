package recipe

import "github.com/go-webserver/internal/models"

type RecipeRepo interface {
	Create(recipe *models.RecipeRequest) (string, error)
	List(opts *models.RecipeFilter) ([]*models.Recipe, error)
	Get(id string) (*models.Recipe, error)
	Delete(id string) error
	Update(Id string, name *string, prep *string, cook *string, ingredients *[]string, instructions *[]string) error
}
