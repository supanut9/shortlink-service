package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/supanut9/shortlink-service/internal/handler"
	"github.com/supanut9/shortlink-service/internal/repository"
	"github.com/supanut9/shortlink-service/internal/service"
)

func Setup(app *fiber.App) {
	// Repository
	linkRepo := repository.NewLinkRepository()
	clickEventRepo := repository.NewClickEventRepository()

	// Service
	linkService := service.NewLinkService(linkRepo)
	clickEventService := service.NewClickEventService(clickEventRepo)

	// Handler
	linkHandler := handler.NewLinkHandler(linkService)
	redirectHandler := handler.NewRedirectHandler(linkService, clickEventService)

	// Public Controller
	redirectHandler.RegisterRedirectRoutes(app)

	// Private Controller
	api := app.Group("/api/v1")
	linkHandler.RegisterLinkRoutes(api.Group("/links"))

}
