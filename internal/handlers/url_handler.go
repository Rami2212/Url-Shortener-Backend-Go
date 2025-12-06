package handlers

import (
	"net/url"

	"github.com/gofiber/fiber/v2"

	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/config"
	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/services"
	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/utils"
)

type URLHandler struct {
	service *services.URLService
	cfg     *config.Config
}

func NewURLHandler(service *services.URLService, cfg *config.Config) *URLHandler {
	return &URLHandler{
		service: service,
		cfg:     cfg,
	}
}

type shortenRequest struct {
	URL string `json:"url"`
}

func (h *URLHandler) Shorten(c *fiber.Ctx) error {
	var body shortenRequest

	if err := c.BodyParser(&body); err != nil {
		return utils.JSONError(c, fiber.StatusBadRequest, "Invalid JSON body")
	}

	if _, err := url.ParseRequestURI(body.URL); err != nil {
		return utils.JSONError(c, fiber.StatusBadRequest, "Invalid URL format")
	}

	shortURL, code, err := h.service.Shorten(body.URL)
	if err != nil {
		return utils.JSONError(c, fiber.StatusInternalServerError, "Failed to shorten URL")
	}

	return utils.JSONSuccess(c, fiber.Map{
		"short_url":    shortURL,
		"short_code":   code,
		"original_url": body.URL,
	})
}

func (h *URLHandler) Redirect(c *fiber.Ctx) error {
	code := c.Params("code")

	originalURL, err := h.service.ResolveShortCode(code)
	if err != nil {
		return utils.JSONError(c, fiber.StatusNotFound, "Short URL not found")
	}

	return c.Redirect(originalURL, fiber.StatusTemporaryRedirect)
}
