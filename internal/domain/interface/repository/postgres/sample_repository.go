package postgres

import (
	"context"

	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/domain/model"
)

type SampleRepository interface {
	Get(ctx context.Context, ID string) (*model.Sample, error)
}
