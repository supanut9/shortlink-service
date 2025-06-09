package handler

import (
	"bytes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
	"github.com/supanut9/shortlink-service/internal/config"
	httpService "github.com/supanut9/shortlink-service/internal/http-service"
	"github.com/supanut9/shortlink-service/internal/service"
)

type LinkHandler struct {
	service     service.LinkService
	fileService httpService.FileService
}

func NewLinkHandler(linkSvc service.LinkService, fileSvc httpService.FileService) *LinkHandler {
	return &LinkHandler{
		service:     linkSvc,
		fileService: fileSvc,
	}
}

func (h *LinkHandler) RegisterLinkRoutes(router fiber.Router) {
	router.Get(":slug", h.GetLinkBySlug)
	router.Post("/", h.CreateLink)
}

func (h *LinkHandler) GetLinkBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")

	links, err := h.service.GetLinkBySlug(slug)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch links"})
	}
	return c.JSON(links)
}

type CreateLinkRequest struct {
	URL    string `json:"url"`
	QRCode bool   `json:"qrcode"`
}

func (h *LinkHandler) CreateLink(c *fiber.Ctx) error {
	var req CreateLinkRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	slug, err := h.service.CreateLink(req.URL)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch links"})
	}

	cfg := config.Load()

	shortLink := fmt.Sprintf("%s/%s", cfg.URL.BaseUrl, slug)

	if req.QRCode {
		data, err := GenerateQRCodeBuffer(shortLink)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to generate QRCode"})
		}
		qrcodeUrl, err := h.fileService.UploadFile(cfg.QRCode.Bucket, "", shortLink, data)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to upload QRCode"})
		}

		return c.JSON(fiber.Map{
			"message":    "Link created",
			"url":        req.URL,
			"shortenUrl": shortLink,
			"qrcodeUrl":  qrcodeUrl,
		})
	}

	return c.JSON(fiber.Map{
		"message":    "Link created",
		"url":        req.URL,
		"shortenUrl": shortLink,
	})
}

func GenerateQRCodeBuffer(content string) (*bytes.Reader, error) {
	pngData, err := qrcode.Encode(content, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(pngData)
	return reader, nil
}
