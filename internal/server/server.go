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
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog/log"
)

type Server struct {
	*chi.Mux
	cfg *Config
}

func New(cfg *Config) *Server {
	s := &Server{
		Mux: chi.NewRouter(),
		cfg: cfg,
	}
	s.initRoutes()
	return s
}

func (s *Server) initRoutes() {
	s.Use(middleware.Logger)
	s.Use(middleware.Recoverer)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	s.Use(corsMiddleware.Handler)

	s.Get("/", s.homePageHandler)
	s.Get("/query", s.nameAsParamHandler)
	s.Get("/catFact", s.catFactHandler)
	s.Get("/{name}", s.nameHandler)
}

func (s *Server) Run() {
	addr := fmt.Sprintf("%s:%d", s.cfg.Addr, s.cfg.Port)
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
