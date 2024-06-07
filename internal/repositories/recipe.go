package repositories

import (
	"context"
	"time"

	"github.com/go-webserver/internal/interfaces/recipe"
	"github.com/go-webserver/internal/models"
	"github.com/go-webserver/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRecipeRepo struct {
	db *mongo.Database
}

func NewMongoRecipeRepo(db *mongo.Database) recipe.RecipeRepo {
	return &mongoRecipeRepo{db: db}
}

func (m *mongoRecipeRepo) Create(request *models.RecipeRequest) (string, error) {
	createdAt := time.Now()
	result, err := m.db.Collection("recipes").InsertOne(context.TODO(), bson.M{
		"name":         request.Name,
		"prep":         request.Prep,
		"cook":         request.Cook,
		"ingredients":  request.Ingredients,
		"instructions": request.Instructions,
		"createdAt":    createdAt,
		"updatedAt":    createdAt,
	})
	if err != nil {
		return "", err
	}

	oid := result.InsertedID.(primitive.ObjectID)
	oidStr := oid.Hex()
	return oidStr, nil
}

func (m *mongoRecipeRepo) List() ([]*models.Recipe, error) {
	cur, err := m.db.Collection("recipes").Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var recipes []*models.Recipe
	err = cur.All(context.TODO(), &recipes)
	if err != nil {
		return nil, err
	}
	return recipes, nil
}

func (m *mongoRecipeRepo) Get(id string) (*models.Recipe, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var recipeInDB models.Recipe
	err = m.db.Collection("recipes").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&recipeInDB)
	if err != nil {
		if _, ok := err.(utils.RecipeMessage); ok {
			return nil, utils.RecipeNotFound
		}
		return nil, err
	}
	return &recipeInDB, nil
}

func (m *mongoRecipeRepo) Delete(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = m.db.Collection("recipes").DeleteOne(context.TODO(), bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}
