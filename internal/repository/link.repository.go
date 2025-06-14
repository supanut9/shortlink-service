package repository

import (
	"errors"

	"github.com/supanut9/shortlink-service/internal/entity"
	"gorm.io/gorm"
)

type LinkRepository interface {
	Create(link *entity.Link) error
	FindBySlug(hash string) (*entity.Link, error)
	IncrementClicks(link *entity.Link) error
}

type linkRepository struct {
	db *gorm.DB
}

func NewLinkRepository(db *gorm.DB) LinkRepository {
	return &linkRepository{db: db}
}

func (r *linkRepository) Create(link *entity.Link) error {
	return r.db.Create(link).Error
}

func (r *linkRepository) FindBySlug(hash string) (*entity.Link, error) {
	var link entity.Link
	err := r.db.Where("slug = ? ", hash).First(&link).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &link, err
}

func (r *linkRepository) IncrementClicks(link *entity.Link) error {
	link.Clicks++
	return r.db.Save(link).Error
}
