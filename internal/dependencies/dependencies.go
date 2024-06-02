package dependencies

import (
	"github.com/go-webserver/internal/interfaces/recipe"
)

type Dep struct {
	RecipeService recipe.RecipeUseCase
}
