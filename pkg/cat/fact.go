package cat

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

// FetchFact gets a cat fact from the cat fact API
func FetchFact(url string) (fact CatFact, err error) {
	log.Info().Str("url", url).Msg("Getting a cat fact from the cat fact API")
	resp, err := http.Get(url)
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
