package v1

import (
	"net/http"

	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/adapter/primary/http/handler"
	"github.com/go-chi/chi/v5"
)

type AuthRouter struct {
	Handler http.Handler
}

func NewAuthRouter(authHandler *handler.AuthHandler) *AuthRouter {
	r := chi.NewRouter()
	r.Post("/login", authHandler.Login)

	return &AuthRouter{Handler: r}
}
