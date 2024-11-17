package telemetry

import (
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/infrastructure/telemetry/opentelemetry"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	opentelemetry.InitTelemetry,
)
