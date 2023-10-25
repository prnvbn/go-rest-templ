package rest

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Addr    string         `yaml:"addr"`
	Port    int            `yaml:"port"`
	CatFact CatFactsConfig `yaml:"cat-facts"`
}

type CatFactsConfig struct {
	Enabled bool   `yaml:"enabled"`
	URL     string `yaml:"url"`
}

func (cfg *Config) LoadYAMLFile(cfgFilepath string) error {
	absCfgFile, err := filepath.Abs(cfgFilepath)
	if err != nil {
		return errors.Wrapf(err, "failed to get absolute path of config file %s", cfgFilepath)
	}

	// read yaml file
	bs, err := os.ReadFile(absCfgFile)
	if err != nil {
		return errors.Wrapf(err, "error reading config file %s", absCfgFile)
	}

	if err = yaml.Unmarshal(bs, cfg); err != nil {
		return errors.Wrapf(err, "error unmarshalling config file %s", absCfgFile)
	}

	return nil
}
