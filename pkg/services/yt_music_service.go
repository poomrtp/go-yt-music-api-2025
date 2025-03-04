package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/poomrtp/go-yt-music/pkg/utils"
	"github.com/raitonoberu/ytmusic"
)

type YTMusicService interface {
	Search(c *fiber.Ctx) error
}

type ytMusicServiceImpl struct{}

func NewYTMusicService() YTMusicService {
	return &ytMusicServiceImpl{}
}

func (s *ytMusicServiceImpl) Search(c *fiber.Ctx) error {
	searchQ := c.Query("search")
	searchClient := ytmusic.Search(searchQ)

	result, err := searchClient.Next()
	if err != nil {
		return err
	}

	return utils.SendResponse(c, fiber.StatusOK, "Search result", result)
}
