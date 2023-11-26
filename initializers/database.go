package initializers

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database

func ConnectToDatabase() {
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@cluster0.61i6kpu.mongodb.net/?retryWrites=true&w=majority", os.Getenv("MONGO_USERNAME"), os.Getenv("MONGO_PASSWORD")))
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	if err := client.Database(os.Getenv("MONGO_DATABASE_NAME")).RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	MongoDB = client.Database(os.Getenv("MONGO_DATABASE_NAME"))
	fmt.Println("You successfully connected to MongoDB!")
}
