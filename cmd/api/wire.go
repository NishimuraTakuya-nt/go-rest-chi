//go:build wireinject
// +build wireinject

package main

import (
	"github.com/NishimuraTakuya-nt/go-rest-chi/config"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/adapter/primary/http/custommiddleware"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/adapter/primary/http/handler"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/adapter/primary/http/presenter"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/adapter/primary/http/route"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/adapter/primary/http/route/v1"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/adapter/secondary/piyographql"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/common/logger"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/common/telemetry/datadog"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/domain/service"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/domain/usecase"
	"github.com/google/wire"
)

func InitializeRouter(cfg *config.AppConfig, logger logger.Logger, metricsManager *datadog.MetricsManager) (*route.Router, error) {
	wire.Build(
		presenter.Set,
		custommiddleware.Set,
		piyographql.Set,
		service.Set,
		usecase.Set,
		handler.Set,
		v1.Set,
		route.Set,
	)
	return nil, nil
}
