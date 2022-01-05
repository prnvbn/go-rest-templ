package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

const catFactURL = "https://catfact.ninja/fact"

type catFactAPIResp struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

// getCatFact gets a cat fact from the cat fact API
func getCatFact() (cr catFactAPIResp, err error) {

	logrus.Info("Getting a cat fact from the cat fact API")
	resp, err := http.Get(catFactURL)
	if err != nil {
		return cr, err
	}
	defer resp.Body.Close()

	logrus.Info("Reading the cat fact response body")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return cr, err
	}

	logrus.Info("Unmarshalling the cat fact response body")
	err = json.Unmarshal(body, &cr)
	if err != nil {
		return cr, err
	}

	return cr, nil
}
