package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/pusher/pusher-http-go/v5"
	"log"
	"os"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	pusherClient := pusher.Client{ // Initializes a new Pusher client
		AppID:   os.Getenv("PUSHER_APP_ID"),
		Key:     os.Getenv("PUSHER_KEY"),
		Secret:  os.Getenv("PUSHER_SECRET"),
		Cluster: os.Getenv("PUSHER_CLUSTER"),
		Secure:  true,
	}

	app.Post("/api/messages", func(c fiber.Ctx) error { // Handle POST requests to /api/messages
		var data map[string]string // Create a map to hold the incoming data

		if err := c.Bind().Body(&data); err != nil { // Bind the request body to a map
			return err
		}
		pusherClient.Trigger("chat", "message", data) // Trigger a Pusher event with the data received
		return c.JSON([]string{})                     // Respond with an empty JSON array
	})

	log.Fatal(app.Listen(":8000")) // Start the Fiber app on port 8000
}
