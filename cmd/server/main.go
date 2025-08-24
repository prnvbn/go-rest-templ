package main

import (
	"context"
	"fmt"
	"go-rest/internal/server"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/adrg/xdg"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"

	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/log/global"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.27.0"
)

const (
	ConfigPathEnvVar = "SERVER_CONFIG_PATH"
)

var (
	DefaultConfigPath = xdg.ConfigHome + "/server/config.yaml"
)

type Config struct {
	DisableOTELLogs bool          `yaml:"disable-otel-logs"`
	Server          server.Config `yaml:"server"`
}

func initOTelLogs(ctx context.Context) (shutdown func(context.Context) error, err error) {
	exp, err := otlploggrpc.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create OTLP exporter: %w", err)
	}

	lp := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(sdklog.NewBatchProcessor(exp)),
		sdklog.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("auth"),
		)),
	)

	global.SetLoggerProvider(lp)

	return lp.Shutdown, nil
}

func getConfigPath() string {
	if path := os.Getenv(ConfigPathEnvVar); path != "" {
		return path
	}
	return DefaultConfigPath
}

func main() {
	cfgPath := getConfigPath()
	cfg, err := loadConfig(cfgPath)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	if !cfg.DisableOTELLogs {
		shutdown, err := initOTelLogs(context.Background())
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to initialize OTEL logs")
		}
		defer shutdown(context.Background())
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	srv := server.New(&cfg.Server)

	srv.Run(ctx)
}

func loadConfig(cfgPath string) (*Config, error) {
	absCfgFile, err := filepath.Abs(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path of the config file %s: %w", cfgPath, err)
	}

	bs, err := os.ReadFile(absCfgFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err = yaml.Unmarshal(bs, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml file: %w", err)
	}

	return &cfg, nil
}
