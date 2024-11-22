package handler

import (
	"encoding/json"
	"net/http"

	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/adapter/primary/http/dto/request"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/adapter/primary/http/dto/response"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/adapter/primary/http/presenter"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/common/apperror"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/common/logger"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/domain/usecase"
)

type AuthHandler struct {
	logger      logger.Logger
	JSONWriter  *presenter.JSONWriter
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(logger logger.Logger, JSONWriter *presenter.JSONWriter, authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		logger:      logger,
		JSONWriter:  JSONWriter,
		authUsecase: authUsecase,
	}
}

// Login godoc
// @Summary User login
// @Description Authenticate a user and return a JWT token
// @Tags authentication
// @Accept json
// @Produce json
// @Param request body request.LoginRequest true "Login credentials"
// @Success 200 {object} response.LoginResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req request.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.ErrorContext(ctx, "Failed to decode login request", "error", err)
		h.JSONWriter.WriteError(w, apperror.NewBadRequestError("Invalid request body", err))
		return
	}

	// TODO: ユーザー認証のロジックを実装する
	// この例では、単純化のためにユーザー名とパスワードのチェックを省略しています
	userID := req.UserID
	roles := []string{"role:teamA:editor", "role:teamB:viewer"} // 実際のアプリケーションでは、データベースからユーザーのロールを取得する必要があります

	token, err := h.authUsecase.Login(r.Context(), userID, roles)
	if err != nil {
		h.logger.ErrorContext(ctx, "Failed to generate token", "error", err)
		h.JSONWriter.WriteError(w, apperror.NewInternalError("Failed to generate token", err))
		return
	}

	res := response.LoginResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		h.logger.ErrorContext(ctx, "Failed to encode login response", "error", err)
		return
	}

	h.logger.InfoContext(ctx, "Login successful")
}
