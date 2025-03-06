package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/poomrtp/go-yt-music/pkg/handlers"
	"github.com/poomrtp/go-yt-music/pkg/services"
)

// Handler - handles serverless function requests
func Handler(w http.ResponseWriter, r *http.Request) {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))
	fmt.Printf("[INFO] called handler")

	app.Use(cors.New())
	api := app.Group("/api")

	apiHandler := handlers.NewAPIHandler()
	apiHandler.SetupRoutes(api)
	// Initialize services
	ytMusicService := services.NewYTMusicService()

	// Setup routes
	ytMusicHandler := handlers.NewYTMusicHandler(ytMusicService)
	ytMusicHandler.SetupRoutes(api)

	app.Handler()
}
