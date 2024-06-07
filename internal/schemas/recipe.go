package schemas

type RecipeSchemaRequest struct {
	Name         string   `json:"name" binding:"required"`
	Prep         string   `json:"prep"`
	Cook         string   `json:"cook"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
}
