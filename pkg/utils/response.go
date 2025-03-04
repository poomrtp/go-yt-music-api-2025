package utils

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Error      interface{} `json:"error,omitempty"`
}

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func SendResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	response := Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
	return c.Status(statusCode).JSON(response)
}

func SendErrorResponse(c *fiber.Ctx, statusCode int, message string, err interface{}) error {
	response := ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
		Error:      err,
	}

	return c.Status(statusCode).JSON(response)
}

func BadRequestResponse(c *fiber.Ctx, message string, err interface{}) error {
	return SendErrorResponse(c, fiber.StatusBadRequest, message, err)
}

func UnauthorizedResponse(c *fiber.Ctx, message string, err interface{}) error {
	return SendErrorResponse(c, fiber.StatusUnauthorized, message, err)
}

func NotFoundResponse(c *fiber.Ctx, message string, err interface{}) error {
	return SendErrorResponse(c, fiber.StatusNotFound, message, err)
}

func InternalServerErrorResponse(c *fiber.Ctx, message string, err interface{}) error {
	return SendErrorResponse(c, fiber.StatusInternalServerError, message, err)
}
