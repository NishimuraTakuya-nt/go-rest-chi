package mockservice

//go:generate mockgen -package=mockservice -destination=./mock_service.go github.com/NishimuraTakuya-nt/go-rest-chi/internal/domain/service TokenService
