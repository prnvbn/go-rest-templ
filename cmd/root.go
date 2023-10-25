package cmd

import (
	"go-rest/rest"
	"os"

	"github.com/spf13/cobra"
)

const (
	DEFAULT_CAT_FACT_URL = "https://catfact.ninja/fact"
	DEFAULT_PORT         = 8080
	DEFAULT_HOST         = "localhost"
)

// rootCmd represents the base command when called without any subcommands
var (
	cfg     = &rest.Config{}
	rootCmd = &cobra.Command{
		Use:   "go-rest",
		Short: "A simple REST server template",

		Run: func(cmd *cobra.Command, args []string) {
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
	rootCmd.Flags().StringVarP(&cfg.Addr, "addr", "a", DEFAULT_HOST, "Address to run the server on")

	rootCmd.Flags().BoolVarP(&cfg.CatFact.Enabled, "cat-facts", "c", true, "Enable cat facts")
	rootCmd.Flags().StringVarP(&cfg.CatFact.URL, "cat-facts-url", "u", DEFAULT_CAT_FACT_URL, "URL to get cat facts from")
}
