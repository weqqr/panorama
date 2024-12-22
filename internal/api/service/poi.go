package service

import (
	"context"

	"github.com/lord-server/panorama/internal/domain"
)

type (
	POIService struct {
		repo POIRepo
	}

	POIRepo interface {
		ListPOIs(ctx context.Context) ([]domain.POI, error)
	}
)

func NewPOIService(repo POIRepo) *POIService {
	return &POIService{
		repo: repo,
	}
}

type (
	ListPOIsRequest struct{}

	ListPOIsResponse struct {
		POIs []domain.POI
	}
)

func (s *POIService) ListPOIs(ctx context.Context, r ListPOIsRequest) (ListPOIsResponse, error) {
	pois, err := s.repo.ListPOIs(ctx)
	if err != nil {
		return ListPOIsResponse{}, err
	}

	return ListPOIsResponse{
		POIs: pois,
	}, nil
}
