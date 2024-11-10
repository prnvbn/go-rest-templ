package cmd

import (
	"fmt"
	"go-rest/rest"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"

	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	cfgPath string

	rootCmd = &cobra.Command{
		Use:   "go-rest",
		Short: "A simple REST server template",

		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := loadConfig(cfgPath)
			if err != nil {
				log.Fatal().Err(err).Msg("Error loading config")
			}

			server := rest.NewServer(cfg)
			server.Run()
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var ok bool
	cfgPath, ok = os.LookupEnv("REST_CONFIG_PATH")
	if !ok {
		cfgPath = xdg.ConfigHome + "/rest/config.yaml"
	}
}

func loadConfig(cfgPath string) (*rest.Config, error) {
	absCfgFile, err := filepath.Abs(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path of the config file %s: %w", cfgPath, err)
	}

	// read yaml file
	bs, err := os.ReadFile(absCfgFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg rest.Config
	if err = yaml.Unmarshal(bs, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal yalml file: %w", err)
	}

	return &cfg, nil
}
