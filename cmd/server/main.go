package main

import (
	"fmt"
	"go-rest/internal/server"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

const (
	ConfigPathEnvVar = "SERVER_CONFIG_PATH"
)

var (
	DefaultConfigPath = xdg.ConfigHome + "/server/config.yaml"
)

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

	// Run starts the server
	// TODO: add graceful shutdown
	server.New(cfg).Run()
}

func loadConfig(cfgPath string) (*server.Config, error) {
	absCfgFile, err := filepath.Abs(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path of the config file %s: %w", cfgPath, err)
	}

	bs, err := os.ReadFile(absCfgFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg server.Config
	if err = yaml.Unmarshal(bs, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml file: %w", err)
	}

	return &cfg, nil
}
