package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"lexdev_api.com/paws/src/database/infrastructure/config"
)

type MongoConnection struct {
	client *mongo.Client
}

func NewMongoConnection() *MongoConnection {
	return &MongoConnection{}
}

func (c *MongoConnection) Connect(ctx context.Context) error {
	var dbCredentials config.DatabaseCredentials = *config.NewDatabaseCredentials()

	client, err := mongo.Connect(options.Client().ApplyURI(dbCredentials.Uri).SetAppName("footprints"))

	if err != nil {
		return err
	}

	c.client = client

	return nil
}

func (c *MongoConnection) Disconnect(ctx context.Context) error {
	return c.client.Disconnect(ctx)
}

func (c *MongoConnection) Ping(ctx context.Context) error {
	return c.client.Ping(ctx, nil)
}
