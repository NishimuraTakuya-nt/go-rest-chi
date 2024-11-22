package piyographqls

import (
	"context"

	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/core/domain/models"
)

type SampleClient interface {
	BaseClient
	Sample(ctx context.Context, id string) (*models.Sample, error)
	ListSample(ctx context.Context, offset, limit *int) ([]*models.Sample, error)
	CreateSample(ctx context.Context, m *models.Sample) (*models.Sample, error)
}
