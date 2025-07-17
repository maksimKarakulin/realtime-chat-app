package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/pusher/pusher-http-go/v5"
	"log"
	"os"
)

func main() {
	app := fiber.New()
	app.Use(cors.New()) // Enable CORS

	// Load environment variables from .env file if it exists
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}

	log.Println("AppID:", os.Getenv("PUSHER_APP_ID"))
	log.Println("Key:", os.Getenv("PUSHER_APP_KEY"))
	log.Println("Secret:", os.Getenv("PUSHER_APP_SECRET"))
	log.Println("Cluster:", os.Getenv("PUSHER_APP_CLUSTER"))

	// Initialize Pusher client
	pusherClient := pusher.Client{
		AppID:   os.Getenv("PUSHER_APP_ID"),
		Key:     os.Getenv("PUSHER_APP_KEY"),
		Secret:  os.Getenv("PUSHER_APP_SECRET"),
		Cluster: "eu",
		Secure:  true,
	}

	// Handle POST /api/messages
	app.Post("/api/messages", func(c fiber.Ctx) error {
		var data map[string]string

		// Bind JSON body
		if err := c.Bind().JSON(&data); err != nil {
			log.Println("JSON binding failed:", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON body",
			})
		}

		// Validate fields
		if data["username"] == "" || data["message"] == "" {
			log.Println("Missing fields:", data)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Username and message are required",
			})
		}

		// Trigger Pusher event
		if err := pusherClient.Trigger("chat", "message", data); err != nil {
			log.Println("Pusher trigger failed:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to trigger Pusher event",
			})
		}

		return c.JSON(fiber.Map{"status": "ok"})
	})

	app.Get("/", func(c fiber.Ctx) error { // Health check endpoint
		return c.SendString("Server is running!")
	})

	log.Println("Server running on :8000")
	log.Fatal(app.Listen(":8000"))
}
