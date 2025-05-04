package server

import (
	"context"
	"go-rest/pkg/cat"
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
	cr, err := cat.FetchFact(s.cfg.CatFact.URL)
	if err != nil {
		return nil, err
	}

	resp := CatFactResponse{}
	resp.Body.CatFact = cr
	return &resp, nil
}
