package repositories

import (
	"context"
	"time"

	"github.com/go-webserver/internal/interfaces/account"
	"github.com/go-webserver/internal/models"
	"github.com/go-webserver/pkg/utils"
	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoAccountRepo struct {
	db *mongo.Database
}

func NewMongoAccountRepo(db *mongo.Database) account.AccountRepo {
	return &mongoAccountRepo{db: db}
}

func (m *mongoAccountRepo) GetByEmail(email string) (*models.Account, error) {
	var accountInDB models.Account
	err := m.db.Collection("accounts").FindOne(context.TODO(), bson.M{"email": email}).Decode(&accountInDB)
	if err != nil {
		logger.Infof("mongoAccountRepo::Get %v", err)
		return nil, utils.AccountNotFound
	}
	return &accountInDB, nil
}

func (m *mongoAccountRepo) InsertAccount(email string, name string, passwordHashed string) (string, error) {
	createdAt := time.Now()
	result, err := m.db.Collection("accounts").InsertOne(context.TODO(), bson.M{
		"name":           name,
		"email":          email,
		"passwordHashed": passwordHashed,
		"status":         "inactive",
		"verify":         false,
		"createdAt":      createdAt,
		"updatedAt":      createdAt,
	})
	if err != nil {
		return "", err
	}

	oid := result.InsertedID.(primitive.ObjectID)
	oidStr := oid.Hex()
	return oidStr, nil
}
