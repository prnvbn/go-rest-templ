package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type Server struct {
	*chi.Mux
	cfg *Config
}

func New(cfg *Config) *Server {
	router := chi.NewRouter()
	s := &Server{
		Mux: router,
		cfg: cfg,
	}

	s.init()
	return s
}

func (s *Server) Run() {
	addr := fmt.Sprintf("0.0.0.0:%d", s.cfg.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: s,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Info().Str("addr", addr).Msg("Serving API")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("Server error")
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Error().Err(err).Msg("Graceful shutdown did not complete")
		if err := srv.Close(); err != nil {
			log.Error().Err(err).Msg("Could not stop server")
		}
	}
}
