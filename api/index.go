package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/poomrtp/go-yt-music/pkg/handlers"
	"github.com/poomrtp/go-yt-music/pkg/services"
)

// Handler - handles serverless function requests
func Handler(w http.ResponseWriter, r *http.Request) {
	app := fiber.New()

	// Setup CORS
	app.Use(cors.New())

	// Initialize services
	ytMusicService := services.NewYTMusicService()

	// Setup routes
	ytMusicHandler := handlers.NewYTMusicHandler(ytMusicService)
	ytMusicHandler.SetupRoutes(app)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
