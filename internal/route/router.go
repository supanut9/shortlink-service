package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/supanut9/shortlink-service/internal/handler"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")

	// Link management routes
	linkGroup := v1.Group("/links")
	handler.RegisterLinkRoutes(linkGroup)

	// Public redirect
	handler.RegisterRedirectRoutes(app)
}
