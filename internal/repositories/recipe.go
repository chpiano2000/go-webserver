package repositories

import (
	"context"
	"time"

	"github.com/go-webserver/internal/interfaces/recipe"
	"github.com/go-webserver/internal/models"
	"github.com/go-webserver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (m *mongoRecipeRepo) List(opts *models.RecipeFilter) ([]*models.Recipe, error) {
	offset := opts.Offset
	size := opts.Size
	if size == 0 {
		size = 10
	}
	mongoOpts := options.Find().SetSkip(offset).SetLimit(size)
	cur, err := m.db.Collection("recipes").Find(context.TODO(), bson.D{}, mongoOpts)
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
		log.Infof("mongoRecipeRepo::Get %v", err)
		return nil, utils.RecipeNotFound
	}
	return &recipeInDB, nil
}

func (m *mongoRecipeRepo) Delete(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = m.Get(id)
	if err != nil {
		return err
	}

	_, err = m.db.Collection("recipes").DeleteOne(context.TODO(), bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}

func (m *mongoRecipeRepo) Update(Id string, name, prep, cook *string, ingredients, instructions *[]string) error {
	oid, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return err
	}

	_, err = m.Get(Id)
	if err != nil {
		log.Infof("mongoRecipeRepo::Update::Get %v", err)
		return utils.RecipeNotFound
	}

	updateOpts := bson.M{}
	if name != nil {
		updateOpts["name"] = name
	}
	if prep != nil {
		updateOpts["prep"] = prep
	}
	if cook != nil {
		updateOpts["cook"] = cook
	}
	if ingredients != nil {
		updateOpts["ingredients"] = ingredients
	}
	if instructions != nil {
		updateOpts["instructions"] = ingredients
	}

	searchOpts := bson.M{
		"_id": oid,
	}
	update := bson.M{
		"$set": updateOpts,
	}
	_, err = m.db.Collection("recipes").UpdateOne(context.TODO(), searchOpts, update)
	if err != nil {
		return err
	}

	return nil
}
