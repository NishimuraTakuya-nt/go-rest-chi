package piyographqls

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

type BaseClient interface {
	GqlClient() graphql.Client
	Execute(ctx context.Context, operation string, query func() (any, error)) (any, error)
}
