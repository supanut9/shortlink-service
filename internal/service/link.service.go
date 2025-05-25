package service

import (
	"errors"

	"github.com/supanut9/shortlink-service/internal/entity"
	"github.com/supanut9/shortlink-service/internal/repository"
	"github.com/supanut9/shortlink-service/internal/utils"
)

type LinkService interface {
	CreateLink(url string) (string, error)
	GetLinkBySlug(hash string) (*entity.Link, error)
	AddClick(link *entity.Link) error
}

type linkService struct {
	repo repository.LinkRepository
}

func NewLinkService(repo repository.LinkRepository) LinkService {
	return &linkService{repo: repo}
}

func (s *linkService) CreateLink(url string) (string, error) {
	const maxAttempts = 5

	for i := 0; i < maxAttempts; i++ {
		slug := utils.GenerateSlug(8)

		// Check if the slug already exists
		existingLink, err := s.repo.FindBySlug(slug)
		if err != nil {
			return "", err
		}
		if existingLink != nil {
			// Slug exists → retry
			continue
		}

		// Slug is unique → create new link
		link := &entity.Link{
			Slug: slug,
			URL:  url,
		}
		if err := s.repo.Create(link); err != nil {
			return "", err
		}
		return slug, nil
	}

	return "", errors.New("failed to generate unique slug after multiple attempts")
}

func (s *linkService) GetLinkBySlug(slug string) (*entity.Link, error) {
	return s.repo.FindBySlug(slug)
}

func (s *linkService) AddClick(link *entity.Link) error {
	return s.repo.IncrementClicks(link)
}
