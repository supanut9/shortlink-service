package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mssola/user_agent"
	"github.com/supanut9/shortlink-service/internal/service"
)

type RedirectHandler struct {
	linkService       service.LinkService
	clickEventService service.ClickEventService
}

func NewRedirectHandler(linkService service.LinkService,
	clickEventService service.ClickEventService) *RedirectHandler {
	return &RedirectHandler{linkService: linkService,
		clickEventService: clickEventService}
}

func (h *RedirectHandler) RegisterRedirectRoutes(router fiber.Router) {
	router.Get("/:slug", h.redirectLink)
}

func (h *RedirectHandler) redirectLink(c *fiber.Ctx) error {
	slug := c.Params("slug")

	link, err := h.linkService.GetLinkBySlug(slug)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to resolve short link",
		})
	}
	if link == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Short link not found",
		})
	}

	ipAddress := c.IP()
	userAgent := c.Get("User-Agent")
	referrer := c.Get("Referer")

	go func() {
		_ = h.linkService.AddClick(link)
		meta := service.ClickMeta{
			LinkID:    link.ID,
			IPAddress: ipAddress,
			UserAgent: userAgent,
			Referrer:  referrer,
			Platform:  parsePlatform(userAgent),
			Browser:   parseBrowser(userAgent),
		}
		_ = h.clickEventService.Record(meta)
	}()

	return c.Redirect(link.URL, fiber.StatusTemporaryRedirect)
}

func parsePlatform(uaString string) string {
	ua := user_agent.New(uaString)
	return ua.Platform()
}

func parseBrowser(uaString string) string {
	ua := user_agent.New(uaString)
	name, _ := ua.Browser()
	return name
}
