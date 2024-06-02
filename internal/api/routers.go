package api

type Route interface {
	Setup(group string)
}

type Routes []Route

func NewRoutes(healthRoute HealthRouter, recipeRoute RecipeRouter) Routes {
	return Routes{healthRoute, recipeRoute}
}

func (r Routes) Setup(group string) {
	for _, route := range r {
		route.Setup(group)
	}
}
