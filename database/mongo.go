package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	DB_URI string
	Client *mongo.Client
	Ctx    context.Context
	Cancel context.CancelFunc
}

func NewDataBaseConnection() *MongoDB {
	db_uri, err := NewMongoDBURI()

	if err != nil {
		log.Fatal(err)
	}

	client, ctx, cancel, conn_err := Connect(db_uri)

	if conn_err != nil {
		log.Fatal(conn_err)
	}

	database := MongoDB{
		Client: client,
		Ctx:    ctx,
		Cancel: cancel,
	}

	return &database
}

func Ping(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}
