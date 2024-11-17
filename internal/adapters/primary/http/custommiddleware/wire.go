package custommiddleware

import "github.com/google/wire"

var Set = wire.NewSet(
	NewDDTracer,
	NewMetrics,
	NewErrorHandling,
	NewTimeout,
	NewAuthentication,
)
