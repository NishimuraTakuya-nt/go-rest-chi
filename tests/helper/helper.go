package helper

import (
	"time"

	"github.com/NishimuraTakuya-nt/go-rest-chi/config"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/common/logger"
)

// TestConfig はテスト用の設定を提供します
type TestConfig struct {
	*config.AppConfig
}

// NewTestConfig は基本的なテスト用設定を生成します
func NewTestConfig() *TestConfig {
	return &TestConfig{
		AppConfig: &config.AppConfig{
			Env:            "test",
			LogLevel:       "INFO",
			Version:        "test-version",
			ServerAddress:  ":8081",
			AllowedOrigins: []string{"*"},
			JWTSecretKey:   "test-secret-key",
			RequestTimeout: 30 * time.Second,

			// Piyo GraphQL settings
			PiyoGraphQLEndpoint:        "http://test-endpoint",
			PiyoGraphQLTimeout:         5 * time.Second,
			PiyoGraphQLMaxIdleConn:     10,
			PiyoGraphQLMaxPerHost:      10,
			PiyoGraphQLIdleConnTimeout: 90 * time.Second,

			// DataDog settings
			DDEnabled:          false,
			DDAgentHost:        "localhost",
			DDAgentTracePort:   "8126",
			DDAgentMetricsPort: "8125",
			DDSamplingRate:     1.0,
		},
	}
}

// ConfigOption は設定を変更するための関数型です
type ConfigOption func(*config.AppConfig)

// NewTestConfigWithOptions は指定されたオプションで設定を生成します
func NewTestConfigWithOptions(opts ...ConfigOption) *config.AppConfig {
	cfg := NewTestConfig().AppConfig
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// WithPiyoGraphQLEndpoint はGraphQLエンドポイントを設定します
func WithPiyoGraphQLEndpoint(endpoint string) ConfigOption {
	return func(cfg *config.AppConfig) {
		cfg.PiyoGraphQLEndpoint = endpoint
	}
}

// WithPiyoGraphQLTimeout はGraphQLタイムアウトを設定します
func WithPiyoGraphQLTimeout(timeout time.Duration) ConfigOption {
	return func(cfg *config.AppConfig) {
		cfg.PiyoGraphQLTimeout = timeout
	}
}

// WithJWTSecretKey はJWTシークレットキーを設定します
func WithJWTSecretKey(key string) ConfigOption {
	return func(cfg *config.AppConfig) {
		cfg.JWTSecretKey = key
	}
}

// TestLogger はテスト用のロガーを提供します
type TestLogger struct {
	logger.Logger
}

// NewTestLogger は基本的なテスト用ロガーを生成します
func NewTestLogger() *TestLogger {
	return &TestLogger{
		Logger: logger.NewLogger(NewTestConfig().AppConfig),
	}
}

// TestHelper はテストヘルパーの機能を集約します
type TestHelper struct {
	Config *config.AppConfig
	Logger logger.Logger
}

// NewTestHelper は基本的なテストヘルパーを生成します
func NewTestHelper() *TestHelper {
	cfg := NewTestConfig().AppConfig
	return &TestHelper{
		Config: cfg,
		Logger: logger.NewLogger(cfg),
	}
}

// NewTestHelperWithOptions は指定されたオプションでテストヘルパーを生成します
func NewTestHelperWithOptions(opts ...ConfigOption) *TestHelper {
	cfg := NewTestConfigWithOptions(opts...)
	return &TestHelper{
		Config: cfg,
		Logger: logger.NewLogger(cfg),
	}
}
