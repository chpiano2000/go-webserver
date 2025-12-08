package databases

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/go-webserver/config"
)

func NewMongoDB(cfg config.Config) *mongo.Database {
	clientOptions := options.Client().ApplyURI(cfg.MONGO_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to MongoDB")
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to MongoDB")
	}
	log.Println("Connected to MongoDB!")
	return client.Database(cfg.DB_NAME)
}
