package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router *chi.Mux
}

func NewServer() *Server {
	s := &Server{
		router: chi.NewRouter(),
	}

	s.registerPages()

	return s
}

func (s *Server) ListenAndServe(addr string) {
	http.ListenAndServe(addr, s.router)
}
