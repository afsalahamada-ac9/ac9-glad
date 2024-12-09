package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Collection *mongo.Collection
}

func (*Mongo) Connect() *mongo.Collection {
	connection_url := "mongodb+srv://test_user:test_user@cluster0.gchkr.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	clientOptions := options.Client().ApplyURI(connection_url)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Println("conneciton failed")
	}
	log.Println(client.Ping(context.Background(), nil))
	db_name := "SFSync"
	coll_name := "logs"
	collection := client.Database(db_name).Collection(coll_name)
	return collection
}
