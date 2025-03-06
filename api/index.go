package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/poomrtp/go-yt-music/pkg/handlers"
	"github.com/poomrtp/go-yt-music/pkg/services"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// This is needed to set the proper request path in `*fiber.Ctx`
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

// Handler - handles serverless function requests
func handler() http.HandlerFunc {
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

	return adaptor.FiberApp(app)
}
