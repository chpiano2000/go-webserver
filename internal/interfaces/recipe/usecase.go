package recipe

import (
	"github.com/go-webserver/internal/domains"
	"github.com/go-webserver/internal/models"
)

type RecipeUseCase interface {
	Create(request *models.RecipeRequest) domains.Result[*models.Recipe]
	List(opts *models.RecipeFilter) domains.Result[[]*models.Recipe]
	Get(id string) domains.Result[*models.Recipe]
	Delete(id string) domains.Result[bool]
	Update(request *models.RecipeUpdateRequest) domains.Result[*models.Recipe]
}
