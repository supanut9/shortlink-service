package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/supanut9/shortlink-service/internal/service"
)

type LinkHandler struct {
	service service.LinkService
}

func NewLinkHandler(s service.LinkService) *LinkHandler {
	return &LinkHandler{service: s}
}

func (h *LinkHandler) RegisterLinkRoutes(router fiber.Router) {
	// router.Get("/", h.GetLinkByHash)
	router.Get(":slug", h.GetLinkBySlug)
	// router.Post("/", h.CreateLink)
	// router.Put("/:id", h.UpdateLink)
	// router.Delete("/:id", h.DeleteLink)
}

func (h *LinkHandler) GetLinkBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")

	links, err := h.service.GetLinkBySlug(slug)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch links"})
	}
	return c.JSON(links)
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
