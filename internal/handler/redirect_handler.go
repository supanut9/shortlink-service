package handler

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRedirectRoutes(router fiber.Router) {
	router.Get(":code", redirectLink)

}

func redirectLink(c *fiber.Ctx) error {
	// Sample logic
	return c.JSON(fiber.Map{"message": "...Redirecting Link"})
}
