package cmd

import (
	"go-rest/rest"

	"github.com/rs/zerolog/log"

	"os"

	"github.com/spf13/cobra"
)

const (
	DEFAULT_CAT_FACT_URL     = "https://catfact.ninja/fact"
	DEFAULT_PORT             = 8080
	DEFAULT_ADDR             = "localhost"
	DEFAULT_CAT_FACT_ENABLED = true
)

// rootCmd represents the base command when called without any subcommands
var (
	cfgFile string
	cfg     = &rest.Config{
		Addr: DEFAULT_ADDR,
		Port: DEFAULT_PORT,
		CatFact: rest.CatFactsConfig{
			Enabled: false,
			URL:     DEFAULT_CAT_FACT_URL,
		},
	}
	rootCmd = &cobra.Command{
		Use:   "go-rest",
		Short: "A simple REST server template",

		Run: func(cmd *cobra.Command, args []string) {

			if cfgFile != "" {
				if err := cfg.LoadYAMLFile(cfgFile); err != nil {
					log.Fatal().Err(err).Msg("Error loading config")
				}
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
	rootCmd.Flags().IntVarP(&cfg.Port, "port", "p", DEFAULT_PORT, "Port to run the server on")
	rootCmd.Flags().StringVarP(&cfg.Addr, "addr", "a", DEFAULT_ADDR, "Address to run the server on")

	rootCmd.Flags().BoolVarP(&cfg.CatFact.Enabled, "cat-facts", "f", DEFAULT_CAT_FACT_ENABLED, "Enable cat facts")
	rootCmd.Flags().StringVarP(&cfg.CatFact.URL, "cat-facts-url", "u", DEFAULT_CAT_FACT_URL, "URL to get cat facts from")

	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "Config file path (cannot use with any other flag)")

	// mark config flag as mutually exclusive with other flags
	for _, f := range []string{"port", "addr", "cat-facts", "cat-facts-url"} {
		rootCmd.MarkFlagsMutuallyExclusive("config", f)
	}

	// rootCmd.MarkFlagsMutuallyExclusive("config", "port", "addr", "cat-facts", "cat-facts-url")
}
