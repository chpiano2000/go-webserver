package models

import "time"

type Recipe struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Prep         string    `json:"prep"`
	Cook         string    `json:"cook"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type RecipeRequest struct {
	Name         string   `json:"name"`
	Prep         string   `json:"prep"`
	Cook         string   `json:"cook"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
}
