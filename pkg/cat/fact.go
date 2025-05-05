package cat

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog"
)

type FactConfig struct {
	URL string `yaml:"url"`
}

type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

// FetchFact gets a cat fact from the cat fact API
func FetchFact(ctx context.Context, cfg FactConfig) (fact CatFact, err error) {
	log := zerolog.Ctx(ctx)
	log.Info().Str("url", cfg.URL).Msg("Getting a cat fact from the cat fact API")

	resp, err := http.Get(cfg.URL)
	if err != nil {
		return fact, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fact, err
	}

	err = json.Unmarshal(body, &fact)
	if err != nil {
		return fact, err
	}

	return fact, nil
}
