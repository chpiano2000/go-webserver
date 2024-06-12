package models

import (
	"time"
)

type Recipe struct {
	Id           string    `json:"id" bson:"_id"`
	Name         string    `json:"name"`
	Prep         string    `json:"prep"`
	Cook         string    `json:"cook"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type RecipeFilter struct {
	Offset int64 `form:"offset" json:"offset" binding:"omitempty,gte=0"`
	Size   int64 `form:"size" json:"size" binding:"omitempty,gte=0,lte=100"`
	// Name   *string `json:"name"`
}

type RecipeRequest struct {
	Name         string   `json:"name"`
	Prep         string   `json:"prep"`
	Cook         string   `json:"cook"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
}

type RecipeUpdateRequest struct {
	Id           string   `json:"id" bson:"_id"`
	Name         string   `json:"name"`
	Prep         string   `json:"prep"`
	Cook         string   `json:"cook"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
}
