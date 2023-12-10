package http

import (
	"html/template"
	"net/http"

	"github.com/cativovo/wisik/pkg/image"
)

var homePageTmpl *template.Template = template.Must(template.ParseFiles("web/templates/base.html"))

type page struct {
	imageService *image.Service
}

func (s *Server) registerPages(is *image.Service) {
	p := page{
		imageService: is,
	}

	s.router.Get("/", p.homePage)
}

func (p *page) homePage(w http.ResponseWriter, r *http.Request) {
	images, err := p.imageService.ListImages()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	homePageTmpl.Execute(w, map[string]any{
		"Images": images,
	})
}
