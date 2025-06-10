package handler

import (
	"bytes"
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
	"github.com/supanut9/shortlink-service/internal/config"
	httpService "github.com/supanut9/shortlink-service/internal/http-service"
	"github.com/supanut9/shortlink-service/internal/repository"
	"github.com/supanut9/shortlink-service/internal/service"
)

type LinkHandler struct {
	service     service.LinkService
	fileService httpService.FileService
	cfg         *config.Config
}

func NewLinkHandler(linkSvc service.LinkService, fileSvc httpService.FileService, cfg *config.Config) *LinkHandler {
	return &LinkHandler{
		service:     linkSvc,
		fileService: fileSvc,
		cfg:         cfg,
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
		if errors.Is(err, repository.ErrUniqueSlugGenerationFailed) {
			// Return a 503 Service Unavailable status
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error": "Cannot process request at this time, please try again later.",
			})
		}

		log.Printf("unhandled database error in CreateLink: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "A database error occurred."})
	}

	shortLink := fmt.Sprintf("%s/%s", h.cfg.URL.BaseUrl, slug)

	if req.QRCode {
		data, err := GenerateQRCodeBuffer(shortLink)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to generate QRCode"})
		}
		qrcodeUrl, err := h.fileService.UploadFile(h.cfg.QRCode.Bucket, "shortlink-qrcodes", shortLink, data)

		if err != nil {
			log.Printf("file service error: %v", err)

			if errors.Is(err, repository.ErrInsufficientStorage) {
				return c.Status(fiber.StatusInsufficientStorage).JSON(fiber.Map{
					"error": "Cannot save QR code, insufficient storage on file server.",
				})
			}

			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error": "The QR code could not be saved due to an issue with an upstream service.",
			})
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
