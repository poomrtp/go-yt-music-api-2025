package services

import (
	"fmt"

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
	fmt.Printf("[INFO] Search query: %v\n", searchQ)
	result, err := searchClient.Next()
	if err != nil {
		fmt.Printf("[ERROR] Search: %v\n", err)
		return utils.BadRequestResponse(c, "Cannot search: ", err)
	}

	return utils.SendResponse(c, fiber.StatusOK, "Search result", result)
}
