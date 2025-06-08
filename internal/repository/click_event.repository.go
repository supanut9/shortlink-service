package repository

import (
	"github.com/supanut9/shortlink-service/internal/entity"
	"gorm.io/gorm"
)

type ClickEventRepository interface {
	Create(event *entity.ClickEvent) error
}

type clickEventRepository struct {
	db *gorm.DB
}

func NewClickEventRepository(db *gorm.DB) ClickEventRepository {
	return &clickEventRepository{db: db}
}

func (r *clickEventRepository) Create(event *entity.ClickEvent) error {
	return r.db.Create(event).Error
}
