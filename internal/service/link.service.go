package service

import (
	"github.com/supanut9/shortlink-service/internal/entity"
	"github.com/supanut9/shortlink-service/internal/repository"
)

type LinkService interface {
	CreateLink(link *entity.Link) error
	GetLinkBySlug(hash string) (*entity.Link, error)
	AddClick(link *entity.Link) error
}

type linkService struct {
	repo repository.LinkRepository
}

func NewLinkService(repo repository.LinkRepository) LinkService {
	return &linkService{repo: repo}
}

func (s *linkService) CreateLink(link *entity.Link) error {
	return s.repo.Create(link)
}

func (s *linkService) GetLinkBySlug(slug string) (*entity.Link, error) {
	return s.repo.FindBySlug(slug)
}

func (s *linkService) AddClick(link *entity.Link) error {
	return s.repo.IncrementClicks(link)
}
