package repositories

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/go-webserver/internal/domains"
	"github.com/go-webserver/internal/interfaces/recipe"
	"github.com/go-webserver/internal/models"
)

type mongoRecipeRepo struct {
	db *mongo.Database
}

func NewMongoRecipeRepo(db *mongo.Database) recipe.RecipeRepo {
	return &mongoRecipeRepo{db: db}
}

func (m *mongoRecipeRepo) Create(request *models.RecipeRequest) domains.Result[string] {
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
		log.Errorf("mongoRecipeRepo::Create failed: %v", err)
		// Infrastructure error - convert to domain error
		return domains.Failure[string](domains.ErrDatabaseConnection)
	}
	oid := result.InsertedID.(primitive.ObjectID)
	oidStr := oid.Hex()
	return domains.Success(oidStr)
}

func (m *mongoRecipeRepo) List(opts *models.RecipeFilter) domains.Result[[]*models.Recipe] {
	offset := opts.Offset
	size := opts.Size
	if size == 0 {
		size = 10
	}
	mongoOpts := options.Find().SetSkip(offset).SetLimit(size)
	cur, err := m.db.Collection("recipes").Find(context.TODO(), bson.D{}, mongoOpts)
	if err != nil {
		log.Errorf("mongoRecipeRepo::List query failed: %v", err)
		return domains.Failure[[]*models.Recipe](domains.ErrDatabaseConnection)
	}

	var recipes []*models.Recipe
	err = cur.All(context.TODO(), &recipes)
	if err != nil {
		log.Errorf("mongoRecipeRepo::List decode failed: %v", err)
		return domains.Failure[[]*models.Recipe](domains.ErrDatabaseConnection)
	}
	return domains.Success(recipes)
}

func (m *mongoRecipeRepo) Get(id string) domains.Result[*models.Recipe] {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Infof("mongoRecipeRepo::Get invalid id: %s", id)
		return domains.Failure[*models.Recipe](domains.RecipeError.NotFound())
	}

	var recipeInDB models.Recipe
	err = m.db.Collection("recipes").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&recipeInDB)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Expected error - entity not found
			return domains.Failure[*models.Recipe](domains.RecipeError.NotFound())
		}
		// Unexpected infrastructure error
		log.Errorf("mongoRecipeRepo::Get database error: %v", err)
		return domains.Failure[*models.Recipe](domains.ErrDatabaseConnection)
	}
	return domains.Success(&recipeInDB)
}

func (m *mongoRecipeRepo) Delete(id string) domains.Result[bool] {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domains.Failure[bool](domains.RecipeError.NotFound())
	}

	getResult := m.Get(id)
	if getResult.IsFailure() {
		return domains.Failure[bool](*getResult.Error())
	}

	_, err = m.db.Collection("recipes").DeleteOne(context.TODO(), bson.M{"_id": oid})
	if err != nil {
		log.Errorf("mongoRecipeRepo::Get database error: %v", err)
		return domains.Failure[bool](domains.ErrDatabaseConnection)
	}
	return domains.Success(true)
}

func (m *mongoRecipeRepo) Update(
	Id string,
	name, prep, cook *string,
	ingredients, instructions *[]string,
) domains.Result[bool] {
	oid, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return domains.Failure[bool](domains.RecipeError.NotFound())
	}

	getResult := m.Get(Id)
	if err != nil {
		log.Infof("mongoRecipeRepo::Update::Get %v", err)
		return domains.Failure[bool](*getResult.Error())
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
		log.Errorf("mongoRecipeRepo::Get database error: %v", err)
		return domains.Failure[bool](domains.ErrDatabaseConnection)
	}

	return domains.Success(true)
}
