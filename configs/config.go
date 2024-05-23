package configs

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	mongodbURI := os.Getenv("MONGODB_URI")

	clientOptions := options.Client().ApplyURI(mongodbURI)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	Client, err = mongo.Connect(ctx, clientOptions)

	if err != nil {
		fmt.Println("Error connecting to MongoDB", err)
	}

	err = Client.Ping(ctx, nil)

	if err != nil {
		fmt.Println("Error pinging MongoDB", err)
	} else {
		fmt.Println("Successfully connected to MongoDB")
	}
}
