package database

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// verify we can load .env file and create mongo url to connect to
func NewMongoDBURI() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", errors.New("error loading.env file")
	}

	DB_URI := fmt.Sprintf("mongodb+srv://%s:%s@%s",
		os.Getenv("MONGO_DB_USER"), os.Getenv("MONGO_DB_PASSWORD"), os.Getenv("MONGO_DB_PATH"))

	return DB_URI, nil
}

func Close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func Connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}
