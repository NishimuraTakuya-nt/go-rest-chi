package piyographqls

import (
	"context"

	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/domain/model"
)

type SampleClient interface {
	BaseClient
	Sample(ctx context.Context, id string) (*model.Sample, error)
	ListSample(ctx context.Context, offset, limit *int) ([]*model.Sample, error)
	CreateSample(ctx context.Context, m *model.Sample) (*model.Sample, error)
}
