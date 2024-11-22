package v1

import (
	"net/http"

	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/adapter/primary/http/handler"
	"github.com/go-chi/chi/v5"
)

type HealthcheckRouter struct {
	Handler http.Handler
}

func NewHealthcheckRouter(healthcheckHandler *handler.HealthcheckHandler) *HealthcheckRouter {
	r := chi.NewRouter()
	r.Get("/", healthcheckHandler.Get)

	return &HealthcheckRouter{Handler: r}
}
