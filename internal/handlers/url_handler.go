package handlers

import (
	"net/url"

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
		return utils.JSONError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if body.URL == "" {
		return utils.JSONError(c, fiber.StatusBadRequest, "URL is required")
	}

	// Basic URL validation
	if _, err := url.ParseRequestURI(body.URL); err != nil {
		return utils.JSONError(c, fiber.StatusBadRequest, "Invalid URL format")
	}

	shortURL, code, err := h.service.Shorten(body.URL)
	if err != nil {
		return utils.JSONError(c, fiber.StatusInternalServerError, "Failed to shorten URL")
	}

	resp := fiber.Map{
		"short_url":    shortURL,
		"short_code":   code,
		"original_url": body.URL,
	}

	return utils.JSONSuccess(c, resp)
}

// GET /:code
func (h *URLHandler) Redirect(c *fiber.Ctx) error {
	code := c.Params("code")
	if code == "" {
		return utils.JSONError(c, fiber.StatusBadRequest, "Short code is required")
	}

	originalURL, err := h.service.ResolveShortCode(code)
	if err != nil {
		return utils.JSONError(c, fiber.StatusNotFound, "Short URL not found")
	}

	// 307 keeps the method (POST stays POST), but for simple GET a 302 is also fine
	return c.Redirect(originalURL, fiber.StatusTemporaryRedirect)
}
