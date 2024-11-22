package presenter

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/common/apperror"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/common/logger"
)

type JSONWriter struct {
	logger logger.Logger
}

func NewJSONWriter(logger logger.Logger) *JSONWriter {
	return &JSONWriter{
		logger: logger,
	}
}

func (p *JSONWriter) Write(ctx context.Context, w http.ResponseWriter, data any) {
	select {
	case <-ctx.Done():
		return
	default:
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		p.logger.ErrorContext(ctx, "Failed to encode response", "error", err)
		p.WriteError(w, apperror.NewInternalError("Failed to encode response", err))
	}
}

func (p *JSONWriter) WriteError(w http.ResponseWriter, err error) {
	if ew, ok := w.(apperror.ErrorWriter); ok {
		ew.WriteError(err)
	} else {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}