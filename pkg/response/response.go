package response

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Success(c *fiber.Ctx, status int16, data any) error {
	return c.JSON(fiber.Map{
		"status": status,
		"data":   data,
	})
}

func Error(c *fiber.Ctx, status int16, message string) error {
	return c.JSON(fiber.Map{
		"status":  status,
		"message": message,
	})
}

func Redirect(c *fiber.Ctx, url string, status int) error {
	if err := c.Redirect(url, http.StatusMovedPermanently); err != nil {
		return Error(c, http.StatusInternalServerError, err.Error())
	}

	return nil
}
