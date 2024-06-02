package recipe

import "github.com/go-webserver/internal/models"

type RecipeRepo interface {
	Create(recipe *models.RecipeRequest) (*models.Recipe, error)
	// List() (models.Recipe, error)
}
