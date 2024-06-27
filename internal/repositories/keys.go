package repositories

import (
	"context"

	"github.com/go-webserver/internal/interfaces/keys"
	"github.com/go-webserver/pkg/utils"
	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoKeyRepo struct {
	db *mongo.Database
}

func NewMongoKeyRepo(db *mongo.Database) keys.KeyAdapter {
	return &mongoKeyRepo{db: db}
}

func (m *mongoKeyRepo) InsertKeys(accountId string, publicKey string, privateKey string, refreshToken string) error {
	accountOid, _ := primitive.ObjectIDFromHex(accountId)

	filter := bson.M{"account": accountOid}
	opts := options.Update().SetUpsert(true)
	update := bson.M{
		"$set": bson.M{
			"account":          accountOid,
			"publicKey":        publicKey,
			"privateKey":       privateKey,
			"refreshToken":     refreshToken,
			"refreshTokenUsed": make([]int, 0),
		},
	}

	_, err := m.db.Collection("keys").UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		logger.Infof("mongoKeyRepo::InsertKeys %v", err)
		return err
	}
	return nil
}

func (m *mongoKeyRepo) RemoveKeysByID(accountId string) error {
	accountOid, _ := primitive.ObjectIDFromHex(accountId)
	deleteResult, err := m.db.Collection("keys").DeleteOne(context.TODO(), bson.M{"account": accountOid})
	if err != nil {
		return utils.InternalServerError
	}
	if deleteResult.DeletedCount == 0 {
		return utils.LogoutUnsuccessfully
	}

	return nil
}
