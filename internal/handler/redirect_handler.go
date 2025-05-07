package handler

import (
	"fmt"

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

	// Optional: increment click count
	_ = h.linkService.AddClick(link)

	req := c.Request()
	fmt.Println(string(c.Request().Header.RawHeaders()))
	fmt.Printf("Raw: %s %s\n", req.Header.Method(), req.URI().FullURI())

	meta := service.ClickMeta{
		LinkID:    link.ID,
		IPAddress: c.IP(),
		UserAgent: c.Get("User-Agent"),
		Referrer:  c.Get("Referer"),
		Platform:  parsePlatform(c.Get("User-Agent")),
		Browser:   parseBrowser(c.Get("User-Agent")),
		// Country: use GeoIP if available
	}

	_ = h.clickEventService.Record(meta)

	// Redirect to original long URL
	return c.Redirect(link.URL, fiber.StatusTemporaryRedirect) // 307
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
