package services

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/config"
	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/models"
	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/repositories"
	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/utils"
)

type URLService struct {
	repo  *repositories.URLRepository
	redis *redis.Client
	cfg   *config.Config
	ctx   context.Context
}

func NewURLService(
	repo *repositories.URLRepository,
	redis *redis.Client,
	cfg *config.Config,
) *URLService {
	return &URLService{
		repo:  repo,
		redis: redis,
		cfg:   cfg,
		ctx:   context.Background(),
	}
}

const redisKeyPrefix = "short:"

// Shorten creates (or reuses) a short link and returns:
//
//	shortURL = BASE_SHORT_URL + "/" + code
func (s *URLService) Shorten(originalURL string) (shortURL string, code string, err error) {
	// 1) If URL already exists, reuse it
	if existing, err := s.repo.FindByOriginalURL(originalURL); err == nil && existing != nil {
		short := s.cfg.BaseShortURL + "/" + existing.ShortCode
		return short, existing.ShortCode, nil
	}

	// 2) Generate unique short code
	var shortCode string
	for i := 0; i < 5; i++ { // few retries just in case of collisions
		shortCode, err = utils.GenerateShortCode(7)
		if err != nil {
			return "", "", err
		}

		_, errFind := s.repo.FindByShortCode(shortCode)
		if errors.Is(errFind, gorm.ErrRecordNotFound) {
			// code is free
			break
		}

		// if no error and found, loop again to get a new code
	}

	if shortCode == "" {
		return "", "", errors.New("failed to generate unique short code")
	}

	// 3) Save to DB
	urlModel := &models.URL{
		OriginalURL: originalURL,
		ShortCode:   shortCode,
	}

	if err := s.repo.Create(urlModel); err != nil {
		return "", "", err
	}

	// 4) Put into Redis cache
	_ = s.redis.Set(s.ctx, redisKeyPrefix+shortCode, originalURL, 24*time.Hour).Err()

	return s.cfg.BaseShortURL + "/" + shortCode, shortCode, nil
}

// ResolveShortCode returns the original URL for a given short code
func (s *URLService) ResolveShortCode(code string) (string, error) {
	key := redisKeyPrefix + code

	// 1) Try Redis
	if val, err := s.redis.Get(s.ctx, key).Result(); err == nil {
		return val, nil
	}

	// 2) Fallback to DB
	urlModel, err := s.repo.FindByShortCode(code)
	if err != nil {
		return "", err
	}

	// 3) Increment visits (ignore error for now)
	_ = s.repo.IncrementVisits(urlModel.ID)

	// 4) Put into Redis
	_ = s.redis.Set(s.ctx, key, urlModel.OriginalURL, 24*time.Hour).Err()

	return urlModel.OriginalURL, nil
}
