package http

import (
	"net/http"

	"github.com/cativovo/wisik/pkg/image"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	router *chi.Mux
}

func NewServer(is *image.Service) *Server {
	s := &Server{
		router: chi.NewRouter(),
	}

	// middlewares
	s.router.Use(middleware.Logger)

	// pages/handlers
	s.serveStaticFiles()
	s.registerPages(is)
	s.registerHandlers(is)

	return s
}

func (s *Server) serveStaticFiles() {
	dir := "public"
	path := "/" + dir

	s.router.Get(path+"/*", func(w http.ResponseWriter, r *http.Request) {
		fs := http.StripPrefix(path, http.FileServer(http.Dir(dir)))
		fs.ServeHTTP(w, r)
	})
}

func (s *Server) ListenAndServe(addr string) {
	http.ListenAndServe(addr, s.router)
}
