package server

import (
	"go-rest/pkg/cat"
)

type Config struct {
	Port    int            `yaml:"port"`
	CatFact cat.FactConfig `yaml:"cat-facts"`
}
