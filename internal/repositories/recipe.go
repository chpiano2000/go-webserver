package repositories

import (
	"context"
	"time"

	"github.com/go-webserver/internal/models"
	"github.com/go-webserver/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRecipeRepo struct {
	db *mongo.Database
}

func NewMongoRecipeRepo(db *mongo.Database) *MongoRecipeRepo {
	return &MongoRecipeRepo{db: db}
}

func (m *MongoRecipeRepo) Create(request *models.RecipeRequest) (*models.Recipe, error) {
	createdAt := time.Now()
	uuid := utils.GenerateUUID()
	recipe := models.Recipe{
		Id:           uuid,
		Name:         request.Name,
		Prep:         request.Prep,
		Cook:         request.Cook,
		Ingredients:  request.Ingredients,
		Instructions: request.Instructions,
		CreatedAt:    createdAt,
		UpdatedAt:    createdAt,
	}
	_, err := m.db.Collection("recipes").InsertOne(context.TODO(), recipe)
	if err != nil {
		return nil, err
	}
	return &recipe, nil
}

func (m *MongoRecipeRepo) List() ([]*models.Recipe, error) {
	cur, err := m.db.Collection("recipes").Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var recipes []*models.Recipe
	for cur.Next(context.TODO()) {
		var recipe models.Recipe
		if err := cur.Decode(&recipe); err != nil {
			return nil, err
		}
		recipes = append(recipes, &recipe)
	}
	return recipes, nil
}
