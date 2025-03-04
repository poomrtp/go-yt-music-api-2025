package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/poomrtp/go-yt-music/pkg/middleware"
	"github.com/poomrtp/go-yt-music/pkg/services"
)

type YTMusicHandler struct {
	ytMusicService services.YTMusicService
}

func NewYTMusicHandler(ytMusicService services.YTMusicService) *YTMusicHandler {
	return &YTMusicHandler{ytMusicService: ytMusicService}
}

func (h *YTMusicHandler) SetupRoutes(router fiber.Router) {
	ytMusic := router.Group("/yt-music")
	ytMusic.Use(middleware.Protected())
	ytMusic.Get("/", h.ytMusicService.Search)
}
