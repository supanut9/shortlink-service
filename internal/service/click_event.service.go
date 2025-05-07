package service

import (
	"github.com/supanut9/shortlink-service/internal/entity"
	"github.com/supanut9/shortlink-service/internal/repository"
)

type ClickMeta struct {
	LinkID    uint
	IPAddress string
	UserAgent string
	Referrer  string
	Platform  string
	Browser   string
	Country   string // optional (e.g. via GeoIP)
}

type ClickEventService interface {
	Record(meta ClickMeta) error
}

type clickEventService struct {
	repo repository.ClickEventRepository
}

func NewClickEventService(repo repository.ClickEventRepository) ClickEventService {
	return &clickEventService{repo: repo}
}

func (s *clickEventService) Record(meta ClickMeta) error {
	event := &entity.ClickEvent{
		LinkID:    meta.LinkID,
		IPAddress: meta.IPAddress,
		UserAgent: meta.UserAgent,
		Referrer:  meta.Referrer,
		Platform:  meta.Platform,
		Browser:   meta.Browser,
		Country:   meta.Country,
	}
	return s.repo.Create(event)
}
