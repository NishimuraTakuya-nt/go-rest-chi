package route

import (
	"net/http"

	"github.com/NishimuraTakuya-nt/go-rest-chi/config"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/adapter/primary/http/custommiddleware"
	v1 "github.com/NishimuraTakuya-nt/go-rest-chi/internal/adapter/primary/http/route/v1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Router struct {
	cfg *config.AppConfig
	// middleware
	DDTracer       *custommiddleware.DDTracer
	DDMetrics      *custommiddleware.DDMetrics
	errorHandler   *custommiddleware.ErrorHandling
	timeout        *custommiddleware.Timeout
	authentication *custommiddleware.Authentication
	// router
	healthcheckRouter *v1.HealthcheckRouter
	authRouter        *v1.AuthRouter
	sampleRouter      *v1.SampleRouter
}

func NewRouter(
	cfg *config.AppConfig,
	DDTracer *custommiddleware.DDTracer,
	DDMetrics *custommiddleware.DDMetrics,
	errorHandler *custommiddleware.ErrorHandling,
	Timeout *custommiddleware.Timeout,
	authentication *custommiddleware.Authentication,
	healthcheckRouter *v1.HealthcheckRouter,
	authRouter *v1.AuthRouter,
	sampleRouter *v1.SampleRouter,
) *Router {
	return &Router{
		cfg:               cfg,
		DDTracer:          DDTracer,
		DDMetrics:         DDMetrics,
		errorHandler:      errorHandler,
		timeout:           Timeout,
		authentication:    authentication,
		healthcheckRouter: healthcheckRouter,
		authRouter:        authRouter,
		sampleRouter:      sampleRouter,
	}
}

func (ro *Router) Setup() http.Handler {
	r := chi.NewRouter()
	ro.setupGlobalMiddleware(r)
	ro.setupSwagger(r)
	ro.setupAPIRoutes(r)
	return r
}

func (ro *Router) setupGlobalMiddleware(r *chi.Mux) { // case: datadog SDK
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(custommiddleware.Context())
	r.Use(ro.DDTracer.Handle())
	r.Use(ro.DDMetrics.Handle())

	// セキュリティ関連
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   ro.cfg.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{},
		AllowCredentials: false,
		MaxAge:           300, // 5 minutes
	}))
	r.Use(middleware.SetHeader("X-Content-Type-Options", "nosniff"))
	r.Use(middleware.SetHeader("X-Frame-Options", "DENY"))
	// APP独自
	r.Use(ro.errorHandler.Handle())
	r.Use(ro.timeout.Handle())
}

func (ro *Router) setupSwagger(r *chi.Mux) {
	if ro.cfg.Env == "prd" {
		return
	}

	// Swagger 2.0
	r.Get("/swagger/2.0/*", httpSwagger.Handler(httpSwagger.URL("/docs/openapi/swagger.json")))
	// OAS 3.0
	r.Get("/openapi/*", httpSwagger.Handler(httpSwagger.URL("/docs/openapi/openapi3.json")))
	r.Handle("/docs/openapi/*", http.StripPrefix("/docs/openapi/", http.FileServer(http.Dir("./docs/openapi"))))
}

func (ro *Router) setupAPIRoutes(r *chi.Mux) {
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {

			r.Group(func(r chi.Router) {
				// 認証不要のパブリックルート
				r.Mount("/healthcheck", ro.healthcheckRouter.Handler)
				r.Mount("/auth", ro.authRouter.Handler)
			})

			r.Group(func(r chi.Router) {
				// 認証必要のプライベートルート
				r.Use(ro.authentication.Handle())
				r.Mount("/samples", ro.sampleRouter.Handler)
			})
		})
	})
}
