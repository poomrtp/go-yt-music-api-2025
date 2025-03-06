package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	"github.com/poomrtp/go-yt-music/pkg/handlers"
	"github.com/poomrtp/go-yt-music/pkg/services"
)

func main() {
	godotenv.Load()
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))
	app.Use(cors.New())
	api := app.Group("/api")

	apiHandler := handlers.NewAPIHandler()
	apiHandler.SetupRoutes(api)

	ytMusicService := services.NewYTMusicService()
	ytMusicHandler := handlers.NewYTMusicHandler(ytMusicService)
	ytMusicHandler.SetupRoutes(api)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}

// Handler is the exported function that Vercel will use as the entry point
func Handler(w http.ResponseWriter, r *http.Request) {
	godotenv.Load()
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))
	app.Use(cors.New())
	api := app.Group("/api")

	apiHandler := handlers.NewAPIHandler()
	apiHandler.SetupRoutes(api)

	ytMusicService := services.NewYTMusicService()
	ytMusicHandler := handlers.NewYTMusicHandler(ytMusicService)
	ytMusicHandler.SetupRoutes(api)

	// Use Fiber's handler to serve the request
	app.Handler()
}
