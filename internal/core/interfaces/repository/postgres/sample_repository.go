package postgres

import (
	"context"

	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/core/domain/models"
)

type SampleRepository interface {
	Get(ctx context.Context, ID string) (*models.Sample, error)
}
