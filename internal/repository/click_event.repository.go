package repository

import (
	"github.com/supanut9/shortlink-service/db"
	"github.com/supanut9/shortlink-service/internal/entity"
)

type ClickEventRepository interface {
	Create(event *entity.ClickEvent) error
}

type clickEventRepository struct{}

func NewClickEventRepository() ClickEventRepository {
	return &clickEventRepository{}
}

func (r *clickEventRepository) Create(event *entity.ClickEvent) error {
	return db.DB.Create(event).Error
}
