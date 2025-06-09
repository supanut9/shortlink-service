package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/supanut9/shortlink-service/db"
	"github.com/supanut9/shortlink-service/internal/config"
	"github.com/supanut9/shortlink-service/internal/route"
)

func main() {
	cfg := config.Load()
	app := fiber.New()

	db.InitDB(&cfg.DB)

	route.Setup(app, cfg)

	fmt.Println("App running on port:", cfg.Port)
	app.Listen(":" + cfg.Port)
}
