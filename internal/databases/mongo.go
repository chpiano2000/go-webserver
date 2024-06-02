package databases

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/go-webserver/config"
)

func NewMongoDB(cfg config.Config) *mongo.Database {
	clientOptions := options.Client().ApplyURI(cfg.App.MONGO_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	return client.Database(cfg.App.DB_NAME)
}
