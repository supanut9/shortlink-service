package service

import (
	"errors"

	"github.com/go-sql-driver/mysql"
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

		link := &entity.Link{
			Slug: slug,
			URL:  url,
		}

		err := s.repo.Create(link)
		if err == nil {
			return slug, nil
		}

		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			continue
		}

		return "", err
	}

	return "", repository.ErrUniqueSlugGenerationFailed
}

func (s *linkService) GetLinkBySlug(slug string) (*entity.Link, error) {
	return s.repo.FindBySlug(slug)
}

func (s *linkService) AddClick(link *entity.Link) error {
	return s.repo.IncrementClicks(link)
}
