package handler

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterLinkRoutes(router fiber.Router) {
	router.Get("/", getAllLinks)
	router.Get("/:id", getLinkByID)
	router.Post("/", createLink)
	router.Put("/:id", updateLink)
	router.Delete("/:id", deleteLink)
}

func getAllLinks(c *fiber.Ctx) error {
	// Sample logic
	return c.JSON(fiber.Map{"message": "List of links"})
}

func getLinkByID(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{"message": "Get link by ID", "id": id})
}

func createLink(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Link created"})
}

func updateLink(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{"message": "Link updated", "id": id})
}

func deleteLink(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{"message": "Link deleted", "id": id})
}
