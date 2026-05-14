package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
}

func NewMongoDB(uri string) (*MongoDB, error) {
	if uri == "" {
		panic("MONGO_URI is not set")
	}

	log.Printf("Connecting to MongoDB at: %s", uri)
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	log.Printf("Successfully connected to MongoDB")
	return &MongoDB{client: client}, nil
}

func (m *MongoDB) Disconnect(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}

func (m *MongoDB) Database(name string) *mongo.Database {
	return m.client.Database(name)
}

func (m *MongoDB) Collection(db *mongo.Database, collectionName string) *MongoDBCollection {
	return &MongoDBCollection{
		collection: m.client.Database(db.Name()).Collection(collectionName),
	}
}

type MongoDBCollection struct {
	collection *mongo.Collection
}

func (c *MongoDBCollection) FindOne(filter bson.D) (data bson.M) {
	var result bson.M

	err := c.collection.FindOne(context.Background(), filter).Decode(&result)

	if err == mongo.ErrNoDocuments {
		log.Printf("No document found for filter: %v", filter)
		return nil
	}

	if err != nil {
		panic(err)
	}

	return result
}

func (c *MongoDBCollection) InsertOne(document interface{}) (interface{}, error) {
	log.Printf("Inserting document into collection '%s': %+v\n", c.collection.Name(), document)
	result, err := c.collection.InsertOne(context.Background(), document)

	if err != nil {
		log.Printf("ERROR: Failed to insert document into '%s': %v\n", c.collection.Name(), err)
		return nil, err
	}
	log.Printf("Successfully inserted document with ID: %v\n", result.InsertedID)
	return result.InsertedID, nil
}
