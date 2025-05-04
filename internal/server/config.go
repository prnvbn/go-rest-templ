package server

import (
	"os"
	"path/filepath"

	"go-rest/pkg/cat"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Port    int            `yaml:"port"`
	CatFact cat.FactConfig `yaml:"cat-facts"`
}

func (cfg *Config) LoadYAMLFile(cfgFilepath string) error {
	absCfgFile, err := filepath.Abs(cfgFilepath)
	if err != nil {
		return errors.Wrapf(err, "failed to get absolute path of config file %s", cfgFilepath)
	}

	bs, err := os.ReadFile(absCfgFile)
	if err != nil {
		return errors.Wrapf(err, "error reading config file %s", absCfgFile)
	}

	if err = yaml.Unmarshal(bs, cfg); err != nil {
		return errors.Wrapf(err, "error unmarshalling config file %s", absCfgFile)
	}

	return nil
}
