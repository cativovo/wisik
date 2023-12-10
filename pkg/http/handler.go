package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cativovo/wisik/pkg/image"
)

type handler struct {
	imageService *image.Service
}

func (s *Server) registerHandlers(is *image.Service) {
	h := handler{
		imageService: is,
	}

	s.router.Post("/image", h.addImage)
}

func (h *handler) addImage(w http.ResponseWriter, r *http.Request) {
	src := r.FormValue("src")
	label := r.FormValue("label")

	log.Println(src, label)

	result, err := h.imageService.AddImage(label, src)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println(result)

	fmt.Fprintln(w, "added")
}
