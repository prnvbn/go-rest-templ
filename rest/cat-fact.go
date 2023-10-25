package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
)

type catFactAPIResp struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

// getCatFact gets a cat fact from the cat fact API
func getCatFact(catFactURL string) (cr catFactAPIResp, err error) {

	log.Info().Str("url", catFactURL).Msg("Getting a cat fact from the cat fact API")
	resp, err := http.Get(catFactURL)
	if err != nil {
		return cr, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return cr, err
	}

	err = json.Unmarshal(body, &cr)
	if err != nil {
		return cr, err
	}

	return cr, nil
}
