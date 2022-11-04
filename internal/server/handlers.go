package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) InitHandlers(c *chi.Mux) {

	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	c.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
}
