package main

import (
	"context"
	"fmt"

	"github.com/fabienbrocklesby/NFC-Check-In/configs"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		collection := configs.Client.Database("NFC-CheckIn").Collection("users")

		res, err := collection.InsertOne(context.Background(), bson.M{"name": "John", "age": 30})
		if err != nil {
			fmt.Println("Error inserting document:", err)
			return c.Status(500).SendString("Error inserting document")
		}

		fmt.Println("Inserted document with ID:", res.InsertedID)
		return c.SendString("Inserted document")
	})

	app.Listen(":3000")
}
