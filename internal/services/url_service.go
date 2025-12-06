package services

import (
	"errors"

	"gorm.io/gorm"

	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/config"
	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/models"
	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/repositories"
	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/utils"
)

type URLService struct {
	repo *repositories.URLRepository
	cfg  *config.Config
}

func NewURLService(repo *repositories.URLRepository, cfg *config.Config) *URLService {
	return &URLService{repo: repo, cfg: cfg}
}

func (s *URLService) Shorten(originalURL string) (string, string, error) {
	// If exists, reuse
	if existing, err := s.repo.FindByOriginalURL(originalURL); err == nil {
		return s.cfg.BaseShortURL + "/" + existing.ShortCode, existing.ShortCode, nil
	}

	// Generate unique code
	var shortCode string
	var err error
	for i := 0; i < 5; i++ {
		shortCode, err = utils.GenerateShortCode(7)
		if err != nil {
			return "", "", err
		}

		_, errFind := s.repo.FindByShortCode(shortCode)
		if errors.Is(errFind, gorm.ErrRecordNotFound) {
			break
		}
	}

	if shortCode == "" {
		return "", "", errors.New("failed to generate code")
	}

	urlModel := &models.URL{
		OriginalURL: originalURL,
		ShortCode:   shortCode,
	}

	if err := s.repo.Create(urlModel); err != nil {
		return "", "", err
	}

	return s.cfg.BaseShortURL + "/" + shortCode, shortCode, nil
}

func (s *URLService) ResolveShortCode(code string) (string, error) {
	urlModel, err := s.repo.FindByShortCode(code)
	if err != nil {
		return "", err
	}

	_ = s.repo.IncrementVisits(urlModel.ID)

	return urlModel.OriginalURL, nil
}
