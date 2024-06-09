package schemas

type RecipeSchemaRequest struct {
	Name         string   `json:"name" binding:"required"`
	Prep         string   `json:"prep"`
	Cook         string   `json:"cook"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
}

type RecipeSchemaPayload struct {
	Name         string   `json:"name" binding:"omitempty"`
	Prep         string   `json:"prep" binding:"omitempty"`
	Cook         string   `json:"cook" binding:"omitempty"`
	Ingredients  []string `json:"ingredients" binding:"omitempty"`
	Instructions []string `json:"instructions" binding:"omitempty"`
}
