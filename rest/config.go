package rest

type Config struct {
	Addr    string         `yaml:"addr"`
	Port    int            `yaml:"port"`
	CatFact CatFactsConfig `yaml:"cat-facts"`
}

type CatFactsConfig struct {
	Enabled bool   `yaml:"enabled"`
	URL     string `yaml:"url"`
}
