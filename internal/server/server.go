package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/met4tron/finances-api/internal/config"
	"github.com/rs/zerolog/log"
)

type Server struct {
	chi *chi.Mux
	cfg *config.Config
}

func NewServer() *Server {
	return &Server{
		chi: chi.NewRouter(),
		cfg: config.GetConfig(),
	}
}

func (s *Server) Init() {
	log.Info().Msg("Init API")

	s.cfg = config.LoadConfig()

	s.InitHandlers(s.chi)

	log.Info().Msg("Handlers Mapped")

	http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.Api.Port), s.chi)
}
