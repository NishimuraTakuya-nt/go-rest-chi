package piyographql

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/Khan/genqlient/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/NishimuraTakuya-nt/go-rest-chi/config"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/common/apperror"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/common/logger"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/common/telemetry/datadog"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/domain/interface/external/piyographqls"
)

type baseClient struct {
	logger    logger.Logger
	gqlClient graphql.Client
}

func NewBaseClient(logger logger.Logger, cfg *config.AppConfig) piyographqls.BaseClient {
	gqlClient := &http.Client{
		Timeout: cfg.PiyoGraphQLTimeout,
		Transport: &http.Transport{
			MaxIdleConns:        cfg.PiyoGraphQLMaxIdleConn,
			MaxIdleConnsPerHost: cfg.PiyoGraphQLMaxPerHost,
			IdleConnTimeout:     cfg.PiyoGraphQLIdleConnTimeout,
		},
	}
	return &baseClient{
		logger:    logger,
		gqlClient: graphql.NewClient(cfg.PiyoGraphQLEndpoint, gqlClient),
	}
}

func (c *baseClient) GqlClient() graphql.Client {
	return c.gqlClient
}

func (c *baseClient) Execute(
	ctx context.Context,
	operation string,
	query func() (any, error),
) (any, error) {
	ctx, span := datadog.StartOperation(ctx)
	defer span.Finish()
	span.SetTag("graphql.operation", operation)

	resp, err := query()
	if err != nil {
		wrapErr := c.handleGraphQLError(ctx, operation, err)
		var errorTracer apperror.ErrorTracer
		if errors.As(wrapErr, &errorTracer) {
			errorTracer.AddToSpan(span)
		}
		return nil, wrapErr
	}

	return resp, nil
}

// handleGraphQLError はGraphQLのエラーを適切なAppErrorに変換します
func (c *baseClient) handleGraphQLError(ctx context.Context, operation string, err error) error {
	var gqlErrors gqlerror.List
	if errors.As(err, &gqlErrors) {
		return c.processGraphQLErrors(ctx, operation, gqlErrors)
	}

	// ネットワークエラーの処理
	var netErr *net.OpError
	if errors.As(err, &netErr) {
		return c.handleNetworkError(ctx, operation, netErr)
	}

	// タイムアウトエラーの処理
	if errors.Is(err, context.DeadlineExceeded) {
		return apperror.NewTimeoutError(
			fmt.Sprintf("operation %s timed out", operation),
			err,
		)
	}

	// その他の予期せぬエラー
	return apperror.NewExternalServiceError(
		fmt.Sprintf("unexpected error in operation %s: %v", operation, err),
		err,
	)
}

// processGraphQLErrors はGraphQLエラーを処理します
func (c *baseClient) processGraphQLErrors(ctx context.Context, operation string, errors gqlerror.List) error {
	if len(errors) == 0 {
		return apperror.NewInternalError("empty GraphQL error list", nil)
	}

	c.logger.ErrorContext(ctx, "GraphQL errors occurred",
		"operation", operation,
		"errors", errors)

	// 最初のエラーを使用してエラータイプを判定
	firstErr := errors[0]
	code := ""
	if firstErr.Extensions != nil {
		if c, ok := firstErr.Extensions["code"].(string); ok {
			code = c
		}
	}

	// バリデーションエラーの特別処理
	if isValidationError(code, firstErr.Message) {
		return c.handleValidationErrors(errors)
	}

	// その他のエラータイプの処理
	switch {
	case isAuthError(code, firstErr.Message):
		return apperror.NewUnauthorizedError(
			fmt.Sprintf("authentication failed in %s: %s", operation, firstErr.Message),
			errors[0],
		)
	case isNotFoundError(code, firstErr.Message):
		return apperror.NewNotFoundError(
			fmt.Sprintf("resource not found in %s: %s", operation, firstErr.Message),
			errors[0],
		)
	case isRateLimitError(code, firstErr.Message):
		return apperror.NewRateLimitError(
			fmt.Sprintf("rate limit exceeded in %s: %s", operation, firstErr.Message),
			errors[0],
		)
	default:
		return apperror.NewExternalServiceError(
			fmt.Sprintf("graphql error in %s: %s", operation, firstErr.Message),
			errors[0],
		)
	}
}

// handleValidationErrors はバリデーションエラーを処理します
func (c *baseClient) handleValidationErrors(errors gqlerror.List) error {
	validationErrors := apperror.NewValidationErrors()

	for _, err := range errors {
		field := "unknown"
		if len(err.Path) > 0 {
			//field = err.Path[len(err.Path)-1]
			field = err.Path.String()
		}

		var value interface{}
		if err.Extensions != nil {
			value = err.Extensions["value"]
		}

		validationErrors.AddError(field, value, err.Message)
	}

	return validationErrors
}

// handleNetworkError はネットワークエラーを処理します
func (c *baseClient) handleNetworkError(ctx context.Context, operation string, err *net.OpError) error {
	c.logger.ErrorContext(ctx, "Network error occurred",
		"operation", operation,
		"error_type", err.Op,
		"network", err.Net)

	if err.Timeout() {
		return apperror.NewTimeoutError(
			fmt.Sprintf("network timeout in %s", operation),
			err,
		)
	}

	return apperror.NewExternalServiceError(
		fmt.Sprintf("network error in %s", operation),
		err,
	)
}

// エラー判定ヘルパー関数 // fixme: code を合わせる todo: 必要なcodeを追加
func isValidationError(code string, message string) bool {
	return code == "GRAPHQL_VALIDATION_FAILED" ||
		contains(message, "validation failed", "invalid input")
}

func isAuthError(code string, message string) bool {
	return code == "UNAUTHENTICATED" || code == "FORBIDDEN" ||
		contains(message, "unauthorized", "authentication failed")
}

func isNotFoundError(code string, message string) bool {
	return code == "NOT_FOUND" ||
		contains(message, "not found", "does not exist")
}

func isRateLimitError(code string, message string) bool {
	return code == "RATE_LIMITED" ||
		contains(message, "rate limit", "too many requests")
}

func contains(message string, keywords ...string) bool {
	lowerMsg := strings.ToLower(message)
	for _, keyword := range keywords {
		if strings.Contains(lowerMsg, keyword) {
			return true
		}
	}
	return false
}
