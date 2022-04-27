package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func newRouter(h *handler) (r *chi.Mux) {

	r = chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/", h.uploadCSVFile)
	r.Get("/", h.listCSVData)

	return
}
