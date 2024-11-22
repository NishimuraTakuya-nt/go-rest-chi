package usecases

import (
	"context"

	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/core/domain/models"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/core/interfaces/external/piyographqls"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/infrastructure/logger"
)

type SampleUsecase interface {
	Get(ctx context.Context, ID string) (*models.Sample, error)
	List(ctx context.Context, offset, limit *int) ([]*models.Sample, error)
	Create(ctx context.Context, sample *models.Sample) (*models.Sample, error)
}

type sampleUsecase struct {
	logger        logger.Logger
	graphqlClient piyographqls.SampleClient
}

func NewSampleUsecase(logger logger.Logger, client piyographqls.SampleClient) SampleUsecase {
	return &sampleUsecase{
		logger:        logger,
		graphqlClient: client,
	}
}

func (u *sampleUsecase) Get(ctx context.Context, ID string) (*models.Sample, error) {
	s, err := u.graphqlClient.Sample(ctx, ID)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (u *sampleUsecase) List(ctx context.Context, offset, limit *int) ([]*models.Sample, error) {
	s, err := u.graphqlClient.ListSample(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (u *sampleUsecase) Create(ctx context.Context, m *models.Sample) (*models.Sample, error) {
	s, err := u.graphqlClient.CreateSample(ctx, m)
	if err != nil {
		return nil, err
	}
	return s, nil
}
