package config

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Loader struct {
	v *viper.Viper
}

func NewLoader() *Loader {
	return &Loader{
		v: viper.New(),
	}
}

func (l *Loader) Load() (*AppConfig, error) {
	l.setDefaults()
	l.v.AutomaticEnv()

	var cfg AppConfig
	if err := l.v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (l *Loader) setDefaults() {
	l.v.SetDefault("env", "dev")
	l.v.SetDefault("log_level", "INFO")
	l.v.SetDefault("server_address", ":8081")
	l.v.SetDefault("allowed_origins", []string{"*"})
	l.v.SetDefault("jwt_secret_key", "jwt-secret")
	l.v.SetDefault("request_timeout", 180*time.Second)

	l.v.SetDefault("piyo_graphql_endpoint", "http://localhost:8082/query")
	l.v.SetDefault("piyo_graphql_timeout", 10*time.Second)
	l.v.SetDefault("piyo_graphql_max_idle_conn", 100)
	l.v.SetDefault("piyo_graphql_max_per_host", 100)
	l.v.SetDefault("piyo_graphql_idle_conn_timeout", 90*time.Second)

	l.v.SetDefault("dd_enabled", true)
	l.v.SetDefault("dd_agent_host", "localhost")
	l.v.SetDefault("dd_agent_trace_port", "8126") // case: datadog SDK
	l.v.SetDefault("dd_agent_metrics_port", "8125")
	l.v.SetDefault("dd_sampling_rate", 1.0)
}

type AppConfig struct {
	Env            string        `mapstructure:"env" validate:"required"`
	LogLevel       string        `mapstructure:"log_level"`
	Version        string        `mapstructure:"version"` // fixme increment version
	ServerAddress  string        `mapstructure:"server_address" validate:"required"`
	AllowedOrigins []string      `mapstructure:"allowed_origins" validate:"required"`
	JWTSecretKey   string        `mapstructure:"jwt_secret_key" validate:"required"`
	RequestTimeout time.Duration `mapstructure:"request_timeout" validate:"required"`

	// Piyo GraphQL
	PiyoGraphQLEndpoint        string        `mapstructure:"piyo_graphql_endpoint" validate:"required"`
	PiyoGraphQLTimeout         time.Duration `mapstructure:"piyo_graphql_timeout" validate:"required"`
	PiyoGraphQLMaxIdleConn     int           `mapstructure:"piyo_graphql_max_idle_conn" validate:"required"`
	PiyoGraphQLMaxPerHost      int           `mapstructure:"piyo_graphql_max_per_host" validate:"required"`
	PiyoGraphQLIdleConnTimeout time.Duration `mapstructure:"piyo_graphql_idle_conn_timeout" validate:"required"`

	// DataDog Agent
	DDEnabled          bool    `mapstructure:"dd_enabled" validate:"required"`
	DDAgentHost        string  `mapstructure:"dd_agent_host" validate:"required"`
	DDAgentTracePort   string  `mapstructure:"dd_agent_trace_port" validate:"required"`
	DDAgentMetricsPort string  `mapstructure:"dd_agent_metrics_port" validate:"required"`
	DDSamplingRate     float64 `mapstructure:"dd_sampling_rate" validate:"required"`
}

// Validate validates the config values.
func (c *AppConfig) Validate() error {
	return validator.New().Struct(c)
}
