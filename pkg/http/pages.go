package http

import (
	"html/template"
	"net/http"
)

var homePageTmpl *template.Template = template.Must(template.ParseFiles("web/templates/base.html"))

type page struct{}

func (s *Server) registerPages() {
	p := page{}

	s.router.Get("/", p.homePage)
}

func (p *page) homePage(w http.ResponseWriter, r *http.Request) {
	homePageTmpl.Execute(w, nil)
}
