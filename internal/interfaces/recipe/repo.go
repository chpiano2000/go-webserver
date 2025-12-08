package recipe

import (
	"github.com/go-webserver/internal/domains"
	"github.com/go-webserver/internal/models"
)

type RecipeRepo interface {
	Create(recipe *models.RecipeRequest) domains.Result[string]
	List(opts *models.RecipeFilter) domains.Result[[]*models.Recipe]
	Get(id string) domains.Result[*models.Recipe]
	Delete(id string) domains.Result[bool]
	Update(
		Id string,
		name *string,
		prep *string,
		cook *string,
		ingredients *[]string,
		instructions *[]string,
	) domains.Result[bool]
}
