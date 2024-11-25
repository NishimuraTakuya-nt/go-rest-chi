package piyographql

import (
	"context"
	"fmt"

	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/adapter/secondary/piyographql/generated"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/adapter/secondary/piyographql/mapper"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/common/logger"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/domain/interface/external/piyographqls"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/domain/model"
)

type sampleClient struct {
	logger logger.Logger
	piyographqls.BaseClient
}

func NewSampleClient(logger logger.Logger, baseClient piyographqls.BaseClient) piyographqls.SampleClient {
	return &sampleClient{
		logger:     logger,
		BaseClient: baseClient,
	}
}

const (
	OpSample       = "sample"
	OpListSample   = "listSample"
	OpCreateSample = "createSample"
)

func (c *sampleClient) Sample(ctx context.Context, ID string) (*model.Sample, error) {
	gc := c.GqlClient()
	resp, err := c.Execute(
		ctx,
		OpSample,
		func() (any, error) {
			return generated.SampleQuery(ctx, gc, ID)
		},
	)
	if err != nil {
		return nil, err
	}

	res, ok := resp.(*generated.SampleQueryResponse)
	if !ok {
		return nil, fmt.Errorf("unexpected response type")
	}
	return mapper.ToSampleModel(&res.Sample), nil
}

func (c *sampleClient) ListSample(ctx context.Context, offset, limit *int) ([]*model.Sample, error) {
	gc := c.GqlClient()
	resp, err := c.Execute(
		ctx,
		OpListSample,
		func() (any, error) {
			return generated.ListSampleQuery(ctx, gc, *offset, *limit)
		},
	)
	if err != nil {
		return nil, err
	}

	res, ok := resp.(*generated.ListSampleQueryResponse)
	if !ok {
		return nil, fmt.Errorf("unexpected response type")
	}
	return mapper.ToSampleModelList(res.ListSample), nil
}

func (c *sampleClient) CreateSample(ctx context.Context, m *model.Sample) (*model.Sample, error) {
	gc := c.GqlClient()
	resp, err := c.Execute(
		ctx,
		OpCreateSample,
		func() (any, error) {
			return generated.CreateSampleMutation(ctx, gc, mapper.ToCreateSampleInput(m))
		},
	)
	if err != nil {
		return nil, err
	}

	res, ok := resp.(*generated.CreateSampleMutationResponse)
	if !ok {
		return nil, fmt.Errorf("unexpected response type")
	}
	return mapper.ToSampleModelByCreateSample(&res.CreateSample.Sample), nil
}
