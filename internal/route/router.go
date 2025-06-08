package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/supanut9/shortlink-service/db"
	"github.com/supanut9/shortlink-service/internal/config"
	"github.com/supanut9/shortlink-service/internal/handler"
	httpService "github.com/supanut9/shortlink-service/internal/http-service"
	"github.com/supanut9/shortlink-service/internal/repository"
	"github.com/supanut9/shortlink-service/internal/service"
)

func Setup(app *fiber.App) {
	cfg := config.Load()

	// Database
	dbConnection := db.DB

	// Repository
	linkRepo := repository.NewLinkRepository(dbConnection)
	clickEventRepo := repository.NewClickEventRepository(dbConnection)

	// Service
	linkService := service.NewLinkService(linkRepo)
	clickEventService := service.NewClickEventService(clickEventRepo)
	fileService := httpService.NewFileService(cfg.URL.FileServiceBaseUrl)

	// Handler
	linkHandler := handler.NewLinkHandler(linkService, fileService)
	redirectHandler := handler.NewRedirectHandler(linkService, clickEventService)

	// Public Controller
	redirectHandler.RegisterRedirectRoutes(app)

	// Private Controller
	api := app.Group("/api/v1")
	linkHandler.RegisterLinkRoutes(api.Group("/links"))

}
