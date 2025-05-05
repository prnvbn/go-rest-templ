package server

import (
	"context"
	"go-rest/pkg/cat"

	"github.com/rs/zerolog/log"
)

type NameResponse struct {
	Body struct {
		Message string `json:"message"`
	} `json:"body"`
}

func (s *Server) NameHandler(ctx context.Context, params *struct {
	Name string `path:"name"`
}) (*NameResponse, error) {
	resp := NameResponse{}
	resp.Body.Message = "Hello, " + params.Name + "!"
	return &resp, nil
}

type CatFactResponse struct {
	Body struct {
		cat.CatFact
	} `json:"body"`
}

func (s *Server) CatFactHandler(ctx context.Context, _ *struct{}) (*CatFactResponse, error) {
	log := log.With().Str("handler", "CatFactHandler").Logger()

	log.Info().Msg("Fetching cat fact")

	ctx = log.WithContext(ctx)
	cr, err := cat.FetchFact(ctx, s.cfg.CatFact)
	if err != nil {
		return nil, err
	}

	resp := CatFactResponse{}
	resp.Body.CatFact = cr
	return &resp, nil
}
