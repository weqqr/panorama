package server

import (
	"context"
	"net/http"

	"github.com/lord-server/panorama/internal/api/service"
	"github.com/lord-server/panorama/internal/db/datarepo"
)

type POIHandler struct {
	service *service.POIService
}

func buildPOIService(mux *http.ServeMux, dataRepo *datarepo.DataRepo) {
	poiService := service.NewPOIService(dataRepo)

	poiHandler := &POIHandler{
		service: poiService,
	}

	mux.Handle("GET /api/v1/ListPOIs", handle(poiHandler.ListPOIs))
}

type POI struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type (
	ListPOIsRequest struct{}

	ListPOIsResponse struct {
		POIs []POI `json:"pois"`
	}
)

func (h *POIHandler) ListPOIs(ctx context.Context, request ListPOIsRequest) (ListPOIsResponse, error) {
	_, err := h.service.ListPOIs(ctx, service.ListPOIsRequest{})
	if err != nil {
		return ListPOIsResponse{}, err
	}

	// FIXME:
	return ListPOIsResponse{}, nil
}
