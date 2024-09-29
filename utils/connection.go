package utils

import (
	"context"
	"fmt"
	"github.com/saleh-ghazimoradi/GoCommerce/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ConnectionMongoDB() (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=%s",
		config.AppConfig.Databases.MongoDB.User, config.AppConfig.Databases.MongoDB.Pass, config.AppConfig.Databases.MongoDB.Host, config.AppConfig.Databases.MongoDB.Port, config.AppConfig.Databases.MongoDB.AuthSource)

	clientOptions := options.Client().ApplyURI(uri).SetMaxPoolSize(config.AppConfig.Databases.MongoDB.MaxPoolSize)

	ctx, cancel := context.WithTimeout(context.Background(), config.AppConfig.Databases.MongoDB.Timeout)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")

	return client, nil
}
