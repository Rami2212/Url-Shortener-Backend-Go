package repositories

import (
	"gorm.io/gorm"

	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/models"
)

type URLRepository struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) Create(url *models.URL) error {
	return r.db.Create(url).Error
}

func (r *URLRepository) FindByShortCode(code string) (*models.URL, error) {
	var url models.URL
	err := r.db.Where("short_code = ?", code).First(&url).Error
	return &url, err
}

func (r *URLRepository) FindByOriginalURL(original string) (*models.URL, error) {
	var url models.URL
	err := r.db.Where("original_url = ?", original).First(&url).Error
	return &url, err
}

func (r *URLRepository) IncrementVisits(id uint) error {
	return r.db.Model(&models.URL{}).
		Where("id = ?", id).
		UpdateColumn("visits", gorm.Expr("visits + 1")).
		Error
}
