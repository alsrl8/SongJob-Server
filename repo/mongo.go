package repo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"sync"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

func getCloudUrl() string {
	return os.Getenv("X-SongUser-MongoCloud-Url")
}

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() { // Thread safe
		if clientInstance == nil { // Ensure the instance is initialized only once
			clientOptions := options.Client().ApplyURI(getCloudUrl())
			client, err := mongo.Connect(context.TODO(), clientOptions)
			if err != nil {
				clientInstanceError = fmt.Errorf("%w: %w", &CannotConnectToMongoCloudError{}, err)
			}
			clientInstance = client
		}
	})

	if clientInstanceError != nil {
		return nil, clientInstanceError
	}

	clientInstanceError = clientInstance.Ping(context.TODO(), nil)
	if clientInstanceError != nil {
		return nil, clientInstanceError
	}

	return clientInstance, nil
}

func CloseMongoClient(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Printf("Failed to disconnect mongo clinet: %+v", err)
	}
}

func CloseMongoCursor(cursor *mongo.Cursor) {
	err := cursor.Close(context.TODO())
	if err != nil {
		log.Printf("Failed to close mongo cursor: %+v", cursor)
	}
}

func Find(client *mongo.Client, dbName string, collName string, filter interface{}) (*mongo.Cursor, error) {
	collection := client.Database(dbName).Collection(collName)
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Printf("Failed to find documents: %+v", err)
		log.Printf("Database Name: %s, Collection Name: %s, Filter: %+v", dbName, collName, filter)
		return nil, err
	}
	return cur, nil
}

func FindOne(client *mongo.Client, dbName string, collName string, filter interface{}) *mongo.SingleResult {
	collection := client.Database(dbName).Collection(collName)
	return collection.FindOne(context.TODO(), filter)
}
