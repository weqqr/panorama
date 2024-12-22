package service

import (
	"context"
)

type (
	MetadataService struct{}
)

type (
	GetMetadataRequest struct{}

	GetMetadataResponse struct {
		Link *string
	}
)

func (s *MetadataService) GetMetadata(ctx context.Context, r GetMetadataRequest) (GetMetadataResponse, error) {
	return GetMetadataResponse{
		Link: nil,
	}, nil
}
